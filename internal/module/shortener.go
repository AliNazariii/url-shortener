package module

import (
	"usd/internal/repository"
	"usd/pkg/hashing"
)

type Shortener interface {
	Shorten(url string) (string, error)
	Resolve(shortenURL string) (string, error)
}

type ShortenerImpl struct {
	URLRepo repository.URLRepo
}

func NewShortener(URLRepo repository.URLRepo) Shortener {
	return &ShortenerImpl{
		URLRepo: URLRepo,
	}
}

func (s *ShortenerImpl) Shorten(url string) (string, error) {
	hash := hashing.GetBase32MD5Hash(url)
	err := s.URLRepo.Add(url, hash)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (s *ShortenerImpl) Resolve(shortenURL string) (string, error) {
	return s.URLRepo.FindByShortURL(shortenURL)
}
