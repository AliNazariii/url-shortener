package repository

import "errors"

type URLRepo interface {
	Add(URL string, shortURL string) error
	FindByShortURL(shortURL string) (string, error)
}

type URLRepoImpl struct {
	URLs map[string]string
}

func NewURLRepository() URLRepo {
	return &URLRepoImpl{
		URLs: make(map[string]string),
	}
}

func (U *URLRepoImpl) Add(URL string, shortURL string) error {
	U.URLs[shortURL] = URL
	return nil
}

func (U *URLRepoImpl) FindByShortURL(shortURL string) (string, error) {
	URL, exists := U.URLs[shortURL]
	if !exists {
		return "", errors.New("url not found")
	}
	return URL, nil
}
