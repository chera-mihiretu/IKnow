package repository

import (
	"mime/multipart"

	"github.com/chera-mihiretu/IKnow/infrastructure/storage"
)

type StorageRepository interface {
	UploadFile(file []*multipart.FileHeader) ([]string, error)
	DeleteFile(file []string) error
}

type storageRepository struct {
	supabase storage.SupabaseStorage
}

func NewStorageRepository(supabase storage.SupabaseStorage) StorageRepository {
	return &storageRepository{
		supabase: supabase,
	}
}

func (s *storageRepository) UploadFile(file []*multipart.FileHeader) ([]string, error) {
	return s.supabase.UploadFile(file)
}

func (s *storageRepository) DeleteFile(file []string) error {
	return s.supabase.DeleteFile(file)
}
