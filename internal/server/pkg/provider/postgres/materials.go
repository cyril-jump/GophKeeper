package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

func (p *Provider) GetAllTextData(ctx context.Context, userID string) ([]domain.TextData, error) {
	getTextDataStmt, err := p.db.PrepareContext(ctx, "SELECT id,data,metadata FROM text_data WHERE user_id=$1;")
	if err != nil {
		return nil, &domain.StatementPSQLError{Err: err}
	}
	defer getTextDataStmt.Close()

	rows, err := getTextDataStmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, &domain.ExecutionPSQLError{Err: err}
	}
	defer rows.Close()

	allTextData := make([]domain.TextData, 0)
	for rows.Next() {
		var textData domain.TextData
		err = rows.Scan(&textData.ID, &textData.Text, &textData.Metadata)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, &domain.NotFoundError{Err: domain.ErrDataNotFound}
			default:
				return nil, &domain.ExecutionPSQLError{Err: err}
			}
		}

		allTextData = append(allTextData, textData)
	}

	err = rows.Err()
	if err != nil {
		return nil, &domain.ExecutionPSQLError{Err: err}
	}

	return allTextData, nil
}

func (p *Provider) UpdateTextDataByID(ctx context.Context, userID string, data domain.TextData) error {
	updateTextDataStmt, err := p.db.PrepareContext(ctx, "UPDATE text_data SET data = $1, metadata = $2 WHERE user_id = $3 and id = $4;")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer updateTextDataStmt.Close()

	_, err = updateTextDataStmt.ExecContext(ctx, data.Text, data.Metadata, userID, data.ID)
	if err != nil {
		return &domain.ExecutionPSQLError{Err: err}
	}
	return nil
}

func (p *Provider) CreateNewTextData(ctx context.Context, userID string, data domain.TextData) error {
	crUserStmt, err := p.db.PrepareContext(ctx, "INSERT INTO text_data (user_id, data, metadata) VALUES ($1, $2, $3);")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer crUserStmt.Close()

	if _, err := crUserStmt.ExecContext(ctx, userID, data.Text, data.Metadata); err != nil {
		return &domain.ExecutionPSQLError{Err: err}
	}

	return nil
}

func (p *Provider) GetAllCredData(ctx context.Context, userID string) ([]domain.CredData, error) {

	getTextDataStmt, err := p.db.PrepareContext(ctx, "SELECT id,login, password, metadata FROM auth_data WHERE user_id=$1;")
	if err != nil {
		return nil, &domain.StatementPSQLError{Err: err}
	}
	defer getTextDataStmt.Close()

	rows, err := getTextDataStmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, &domain.ExecutionPSQLError{Err: err}
	}
	defer rows.Close()

	allTextData := make([]domain.CredData, 0)
	for rows.Next() {
		var textData domain.CredData
		err = rows.Scan(&textData.ID, &textData.Login, &textData.Password, &textData.Metadata)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, &domain.NotFoundError{Err: domain.ErrDataNotFound}
			default:
				return nil, &domain.ExecutionPSQLError{Err: err}
			}
		}

		allTextData = append(allTextData, textData)
	}

	err = rows.Err()
	if err != nil {
		return nil, &domain.ExecutionPSQLError{Err: err}
	}

	return allTextData, nil
}

func (p *Provider) UpdateCredDataByID(ctx context.Context, userID string, data domain.CredData) error {
	updateTextDataStmt, err := p.db.PrepareContext(ctx, "UPDATE auth_data SET login = $1, password = $2, metadata = $3 WHERE user_id = $4 and id = $5;")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer updateTextDataStmt.Close()

	_, err = updateTextDataStmt.ExecContext(ctx, data.Login, data.Password, data.Metadata, userID, data.ID)
	if err != nil {
		return &domain.ExecutionPSQLError{Err: err}
	}
	return nil
}

func (r *Provider) CreateNewCredData(ctx context.Context, userID string, data domain.CredData) error {
	crUserStmt, err := r.db.PrepareContext(ctx, "INSERT INTO auth_data (user_id, login, password, metadata) VALUES ($1, $2, $3, $4);")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer crUserStmt.Close()

	if _, err := crUserStmt.ExecContext(ctx, userID, data.Login, data.Password, data.Metadata); err != nil {
		return &domain.ExecutionPSQLError{Err: err}
	}

	return nil
}

func (p *Provider) GetAllCardData(ctx context.Context, userID string) ([]domain.CardData, error) {

	getTextDataStmt, err := p.db.PrepareContext(ctx, "SELECT id,card_number, month, year, cvc, name, surname,metadata FROM card_data WHERE user_id=$1;")
	if err != nil {
		return nil, &domain.StatementPSQLError{Err: err}
	}
	defer getTextDataStmt.Close()

	rows, err := getTextDataStmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, &domain.ExecutionPSQLError{Err: err}
	}
	defer rows.Close()

	allTextData := make([]domain.CardData, 0)
	for rows.Next() {
		var textData domain.CardData
		err = rows.Scan(&textData.ID, &textData.CardNumber, &textData.Month, &textData.Year, &textData.CVC, &textData.Name, &textData.Surname, &textData.Metadata)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, &domain.NotFoundError{Err: domain.ErrDataNotFound}
			default:
				return nil, &domain.ExecutionPSQLError{Err: err}
			}
		}

		allTextData = append(allTextData, textData)
	}

	err = rows.Err()
	if err != nil {
		return nil, &domain.ExecutionPSQLError{Err: err}
	}

	return allTextData, nil
}

func (p *Provider) UpdateCardDataByID(ctx context.Context, userID string, data domain.CardData) error {
	updateTextDataStmt, err := p.db.PrepareContext(ctx, "UPDATE card_data SET card_number = $1, month = $2,year = $3, cvc = $4,name = $5, surname = $6, metadata = $7 WHERE user_id = $8 and id = $9;")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer updateTextDataStmt.Close()

	_, err = updateTextDataStmt.ExecContext(ctx, data.CardNumber, data.Month, data.Year, data.CVC, data.Name, data.Surname, data.Metadata, userID, data.ID)
	if err != nil {
		return &domain.ExecutionPSQLError{Err: err}
	}
	return nil
}

func (p *Provider) CreateNewCardData(ctx context.Context, userID string, data domain.CardData) error {
	crUserStmt, err := p.db.PrepareContext(ctx, "INSERT INTO card_data (user_id, card_number, month, year, cvc, name, surname, metadata) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer crUserStmt.Close()

	if _, err = crUserStmt.ExecContext(ctx, userID, data.CardNumber, data.Month, data.Year, data.CVC, data.Name, data.Surname, data.Metadata); err != nil {
		return &domain.ExecutionPSQLError{Err: err}
	}

	return nil
}

func (p *Provider) GetAllBlobData(ctx context.Context, userID string) ([]domain.BlobData, error) {
	getTextDataStmt, err := p.db.PrepareContext(ctx, "SELECT id,data,metadata FROM blob_data WHERE user_id=$1;")
	if err != nil {
		return nil, &domain.StatementPSQLError{Err: err}
	}
	defer getTextDataStmt.Close()

	rows, err := getTextDataStmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, &domain.ExecutionPSQLError{Err: err}
	}
	defer rows.Close()

	allTextData := make([]domain.BlobData, 0)
	for rows.Next() {
		var textData domain.BlobData
		err = rows.Scan(&textData.ID, &textData.Data, &textData.Metadata)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, &domain.NotFoundError{Err: domain.ErrDataNotFound}
			default:
				return nil, &domain.ExecutionPSQLError{Err: err}
			}
		}

		allTextData = append(allTextData, textData)
	}

	err = rows.Err()
	if err != nil {
		return nil, &domain.ExecutionPSQLError{Err: err}
	}

	return allTextData, nil
}

func (p *Provider) UpdateBlobDataByID(ctx context.Context, userID string, data domain.BlobData) error {
	updateTextDataStmt, err := p.db.PrepareContext(ctx, "UPDATE blob_data SET data = $1, metadata = $2 WHERE user_id = $3 and id = $4;")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer updateTextDataStmt.Close()

	_, err = updateTextDataStmt.ExecContext(ctx, data.Data, data.Metadata, userID, data.ID)
	if err != nil {
		return &domain.ExecutionPSQLError{Err: err}
	}
	return nil
}

func (p *Provider) CreateNewBlobData(ctx context.Context, userID string, data domain.BlobData) error {
	crUserStmt, err := p.db.PrepareContext(ctx, "INSERT INTO blob_data (user_id, data, metadata) VALUES ($1, $2, $3);")
	if err != nil {
		return &domain.StatementPSQLError{Err: err}
	}
	defer crUserStmt.Close()

	if _, err := crUserStmt.ExecContext(ctx, userID, data.Data, data.Metadata); err != nil {
		return &domain.ExecutionPSQLError{Err: err}
	}

	return nil
}
