package storage

import (
	"io"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	basePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		basePath: basePath,
	}
}

func (s *LocalStorage) Save(filename string, file io.Reader) error {
	fullPath := filepath.Join(s.basePath, filename)
	var err error

	// 생성된 디렉터리에 대해 777 권한 부여
	if err = os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return err
	}

	if outFile, err := os.Create(fullPath); err != nil {
		return err
	} else {
		defer outFile.Close()
		if _, err = io.Copy(outFile, file); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func (s *LocalStorage) Delete(path string) error {
	fullPath := filepath.Join(s.basePath, path)
	return os.Remove(fullPath)
}

func (s *LocalStorage) Get(path string) (io.ReadCloser, error) {
	fullPath := filepath.Join(s.basePath, path)
	return os.Open(fullPath)
}
