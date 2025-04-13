package antipatterns

import (
	"errors"
)

type Storage map[string]string

func (s Storage) Get(key string) (string, error) {
	if _, ok := s[key]; ok {
		return s[key], nil
	}

	return "", errors.New("not found key " + key)
}
