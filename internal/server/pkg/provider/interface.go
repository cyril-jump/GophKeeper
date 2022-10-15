package provider

import (
	"context"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

type Provider interface {
	GetAllTextData(ctx context.Context, userID int) ([]domain.TextData, error)
	UpdateTextDataByID(ctx context.Context, userID int, data domain.TextData) error
	CreateNewTextData(ctx context.Context, userID int, data domain.TextData) error

	GetAllCredData(ctx context.Context, userID int) ([]domain.CredData, error)
	UpdateCredDataByID(ctx context.Context, userID int, data domain.CredData) error
	CreateNewCredData(ctx context.Context, userID int, data domain.CredData) error

	GetAllCardData(ctx context.Context, userID int) ([]domain.CardData, error)
	UpdateCardDataByID(ctx context.Context, userID int, data domain.CardData) error
	CreateNewCardData(ctx context.Context, userID int, data domain.CardData) error

	GetAllBlobData(ctx context.Context, userID int) ([]domain.BlobData, error)
	UpdateBlobDataByID(ctx context.Context, userID int, data domain.BlobData) error
	CreateNewBlobData(ctx context.Context, userID int, data domain.BlobData) error

	Create(ctx context.Context, user domain.User) error
	GetByCredentials(ctx context.Context, login, password string) (domain.User, error)

	Close() error
}
