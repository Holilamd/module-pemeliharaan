package repositorys

import "github.com/jackc/pgx/v4"

type Repositoryuser interface {
	Create()
	Update()
	FindById()
	FindByEmail()
}

type repositoryuser struct {
	db *pgx.Conn
}

func NewRepositoryuser(db *pgx.Conn) *repositoryuser {
	return &repositoryuser{db}
}
