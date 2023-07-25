package repositories

import "github.com/mix-plus/go-mixplus/layout/internal/boundedcontexts/hello/domain/entities"

type IHelloRepository interface {
	GetUser(userId int64) (*entities.User, error)
}
