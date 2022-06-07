package repos

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	models "test-manager/usecase_models/boiler"
)

type ProjectsRepository interface {
	UpdateProjects(ctx context.Context, Project models.Project) error
	GetProjects(ctx context.Context, accountId int) ([]*models.Project, error)
	GetProject(ctx context.Context, accountId int, id int) (models.Project, error)
	SaveProjects(ctx context.Context, Project models.Project) (int, error)
}

type projectsRepository struct {
	db *sql.DB
}

func NewProjectsRepositoryRepository(db *sql.DB) ProjectsRepository {
	return &projectsRepository{db: db}
}

func (r *projectsRepository) SaveProjects(ctx context.Context, Project models.Project) (int, error) {
	err := Project.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return Project.ID, nil
}

func (r *projectsRepository) UpdateProjects(ctx context.Context, Project models.Project) error {
	_, err := Project.Update(ctx, r.db, boil.Blacklist("account_id"))
	if err != nil {
		return err
	}
	return nil
}

func (r *projectsRepository) GetProjects(ctx context.Context, accountId int) ([]*models.Project, error) {
	Project, err := models.Projects(models.ProjectWhere.AccountID.EQ(accountId)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	return Project, nil
}

func (r *projectsRepository) GetProject(ctx context.Context, accountId int, id int) (models.Project, error) {
	Project, err := models.Projects(models.ProjectWhere.ID.EQ(id)).One(ctx, r.db)
	if err != nil {
		return models.Project{}, err
	}
	return *Project, nil
}
