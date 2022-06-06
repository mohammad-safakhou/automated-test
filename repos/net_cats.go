package repos

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
	"test-manager/usecase_models"
	models "test-manager/usecase_models/boiler"
)

type NetCatRepository interface {
	UpdateNetCat(ctx context.Context, NetCat models.NetCat) error
	GetNetCat(ctx context.Context, projectId int) (netCatUseCase []*usecase_models.NetCats, err error)
	SaveNetCat(ctx context.Context, NetCat models.NetCat) (int, error)
}

type netCatRepository struct {
	db *sql.DB
}

func NewNetCatRepository(db *sql.DB) NetCatRepository {
	return &netCatRepository{db: db}
}

func (r *netCatRepository) SaveNetCat(ctx context.Context, netCat models.NetCat) (int, error) {
	err := netCat.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return netCat.ID, nil
}

func (r *netCatRepository) UpdateNetCat(ctx context.Context, netCat models.NetCat) error {
	_, err := netCat.Update(ctx, r.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func (r *netCatRepository) GetNetCat(ctx context.Context, projectId int) (netCatUseCase []*usecase_models.NetCats, err error) {
	netCats, err := models.NetCats(models.NetCatWhere.ProjectID.EQ(projectId)).All(ctx, r.db)
	if err != nil {
		return []*usecase_models.NetCats{}, err
	}

	for _, value := range netCats {
		var netCat usecase_models.NetCats
		err := json.Unmarshal([]byte(value.Data.String), &netCat)
		if err != nil {
			log.Println(err.Error())
		}
		netCatUseCase = append(netCatUseCase, &netCat)
	}
	return netCatUseCase, nil
}
