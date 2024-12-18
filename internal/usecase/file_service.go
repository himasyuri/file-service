package usecase

import (
	"context"

	"golang.org/x/sync/semaphore"
	"file-storage/internal/domain"
)

type fileService struct {
	repo          domain.FileRepository
	uploadLimiter *semaphore.Weighted
	listLimiter   *semaphore.Weighted
}

func NewFileService(repo domain.FileRepository) domain.FileService {
	return &fileService{
		repo:          repo,
		uploadLimiter: semaphore.NewWeighted(10),
		listLimiter:   semaphore.NewWeighted(100),
	}
}

func (s *fileService) UploadFile(ctx context.Context, fileName string, data []byte) error {
	if !s.uploadLimiter.TryAcquire(1) {
		return context.DeadlineExceeded
	}
	defer s.uploadLimiter.Release(1)
	return s.repo.SaveFile(fileName, data)
}

func (s *fileService) ListFiles(ctx context.Context) ([]domain.FileMetadata, error) {
	if !s.listLimiter.TryAcquire(1) {
		return nil, context.DeadlineExceeded
	}
	defer s.listLimiter.Release(1)
	return s.repo.ListFiles()
}

func (s *fileService) DownloadFile(ctx context.Context, fileName string) ([]byte, error) {
	return s.repo.GetFile(fileName)
}

