package service

import (
	"Core/storage"
	"errors"
	"mime/multipart"
	"path/filepath"
	"slices"
)

type FileService struct {
	storage storage.Storage
}

func NewFileService(storage storage.Storage) *FileService {
	return &FileService{
		storage: storage,
	}
}

func (fs *FileService) Save(filename string, file multipart.File) error {
	if err := validateFile(filename, file); err != nil {
		return err
	} else {
		return fs.storage.Save(filename, file)
	}
}

func validateFile(filename string, file multipart.File) error {
	// 파일 이름 검증
	// TODO : 허용 확장자 논의
	allowedExtensions := []string{".jpg", ".png"}
	ext := filepath.Ext(filename)

	if !slices.Contains(allowedExtensions, ext) {
		return errors.New("invalid file type")
	} else {
		return nil
	}
}
