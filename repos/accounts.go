package repos

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	models "test-manager/usecase_models/boiler"
)

type AccountsRepository interface {
	UpdateAccounts(ctx context.Context, Account models.Account) error
	GetAccounts(ctx context.Context, id int) (models.Account, error)
	SaveAccounts(ctx context.Context, Account models.Account) (int, error)
}

type accountsRepository struct {
	db *sql.DB
}

func NewAccountsRepositoryRepository(db *sql.DB) AccountsRepository {
	return &accountsRepository{db: db}
}

func (r *accountsRepository) SaveAccounts(ctx context.Context, Account models.Account) (int, error) {
	err := Account.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return Account.ID, nil
}

func (r *accountsRepository) UpdateAccounts(ctx context.Context, Account models.Account) error {
	_, err := Account.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *accountsRepository) GetAccounts(ctx context.Context, id int) (models.Account, error) {
	Account, err := models.Accounts(models.AccountWhere.ID.EQ(id)).One(ctx, r.db)
	if err != nil {
		return models.Account{}, err
	}
	return *Account, nil
}
