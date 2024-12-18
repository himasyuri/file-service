package repository

import (
	"file-storage/internal/domain"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"golang.org/x/sys/unix"
)

type fileRepo struct {
	storagePath string
	mu          sync.Mutex
	metadata    map[string]domain.FileMetadata
}

func NewFileRepo(path string) domain.FileRepository {
	return &fileRepo{
		storagePath: path,
		metadata:    make(map[string]domain.FileMetadata),
	}
}

func (r *fileRepo) SaveFile(fileName string, data []byte) error {
  r.mu.Lock()
  defer r.mu.Unlock()

	filePath := filepath.Join(r.storagePath, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	now := time.Now()

	err = os.Chtimes(filePath, now, now) 
	if err != nil {
		return err
	}

	r.metadata[fileName] = domain.FileMetadata{
		FileName:  fileName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return nil
}

func (r *fileRepo) GetFile(fileName string) ([]byte, error) {
  r.mu.Lock()
  defer r.mu.Unlock()

	filePath := filepath.Join(r.storagePath, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *fileRepo) ListFiles() ([]domain.FileMetadata, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	files := make([]domain.FileMetadata, 0)

	dir, err := os.Open(r.storagePath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, fi := range fileInfos {
		if fi.IsDir() {
			continue
		}

		var creationTime time.Time
		if stat, ok := fi.Sys().(*unix.Stat_t); ok {
			if stat.Ctim.Sec != 0 {
				creationTime = time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec)
			} else {
				creationTime = fi.ModTime()
			}
		} else {
			creationTime = fi.ModTime()
		}

		fileMeta := domain.FileMetadata{
			FileName:  fi.Name(),
			CreatedAt: creationTime,
			UpdatedAt: fi.ModTime(),
		}

		files = append(files, fileMeta)
	}

	return files, nil
}
