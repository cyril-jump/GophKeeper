package rawtodomain

import (
	"github.com/cyril-jump/gophkeeper/internal/client/pkg/raw"
	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

// RawBlobToDomainBlob func
func RawBlobToDomainBlob(data raw.BlobData) *domain.BlobData {

	return &domain.BlobData{
		ID:       data.ID,
		Data:     []byte(data.Data),
		Metadata: data.Metadata,
	}
}
