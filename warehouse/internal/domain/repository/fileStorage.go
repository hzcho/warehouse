package repository

import (
	"context"
	"io"
)

type FileStorage interface {
	SaveFile(ctx context.Context, fileName string, file io.Reader) error
	GetFile(ctx context.Context, fileName string) (io.ReadCloser, error)
	DeleteFile(ctx context.Context, fileName string) error
	IsExists(fileName string) (bool, error)
	AbsPath() string
}
