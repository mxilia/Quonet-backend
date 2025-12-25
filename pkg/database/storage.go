package database

import (
	"fmt"
	"io"

	"github.com/mxilia/Quonet-backend/pkg/config"
	storage_go "github.com/supabase-community/storage-go"
)

type StorageService struct {
	client     *storage_go.Client
	bucketName string
}

func ConnectStorage(cfg *config.Config) *StorageService {
	return &StorageService{
		client: storage_go.NewClient(
			fmt.Sprintf("%s/storage/v1", cfg.SUPABASE_URL),
			cfg.SUPABASE_KEY,
			nil,
		),
		bucketName: cfg.BucketName,
	}
}

func (s *StorageService) UploadFile(uploadPath string, fileData io.Reader, contentType string) error {
	_, err := s.client.UploadFile(s.bucketName, uploadPath, fileData, storage_go.FileOptions{
		ContentType: &contentType,
		Upsert:      new(bool),
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *StorageService) DeleteFile(filePath string) error {
	_, err := s.client.RemoveFile(s.bucketName, []string{filePath})
	if err != nil {
		return fmt.Errorf("failed to delete file from supabase: %w", err)
	}

	return nil
}

func (s *StorageService) GetSignedURL(dbPath string) (string, error) {
	resp, err := s.client.CreateSignedUrl(s.bucketName, dbPath, 24*60*60)
	if err != nil {
		return "", err
	}
	return resp.SignedURL, nil
}
