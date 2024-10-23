package repository

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type FileStorage struct {
	UploadDir string
}

func NewFileStorage(dir string) *FileStorage {
	return &FileStorage{
		UploadDir: dir,
	}
}

func (r *FileStorage) SaveFile(ctx context.Context, fileName string, file io.Reader) error {
	fullPath := filepath.Join(r.UploadDir, fileName)

	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if _, err = io.Copy(outFile, file); err != nil {
		return err
	}

	return nil
}

func (r *FileStorage) GetFile(ctx context.Context, fileName string) (io.ReadCloser, error) {
	fullPath := filepath.Join(r.UploadDir, fileName)

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (r *FileStorage) DeleteFile(ctx context.Context, fileName string) error {
	fullPath := filepath.Join(r.UploadDir, fileName)

	if err := os.Remove(fullPath); err != nil {
		return err
	}

	return nil
}

func (r *FileStorage) IsExists(fileName string) (bool, error) {
	fullPath := filepath.Join(r.UploadDir, fileName)

	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *FileStorage) AbsPath() string {
	return r.UploadDir
}
