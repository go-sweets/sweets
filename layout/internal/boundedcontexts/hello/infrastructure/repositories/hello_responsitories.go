package repositories

import (
	"github.com/mix-plus/go-mixplus/layout/internal/boundedcontexts/hello/domain/entities"
	"github.com/mix-plus/go-mixplus/layout/internal/svc"
)

type HelloRepository struct {
	Svc *svc.ServiceContext
}

func (repo *HelloRepository) GetUser(userId int64) (*entities.User, error) {
	user := entities.NewUser(userId, "mix-plus")
	return user, nil
}
