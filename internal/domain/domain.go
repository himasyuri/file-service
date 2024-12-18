package domain

import (
	"context"
	"time"
)

type FileMetadata struct {
	FileName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FileRepository interface {
	SaveFile(fileName string, data []byte) error
	GetFile(fileName string) ([]byte, error)
	ListFiles() ([]FileMetadata, error)
}

type FileService interface {
  UploadFile(ctx context.Context, fileName string, data []byte) error
  ListFiles(ctx context.Context) ([]FileMetadata, error)
  DownloadFile(ctx context.Context, fileName string) ([]byte, error)
}
