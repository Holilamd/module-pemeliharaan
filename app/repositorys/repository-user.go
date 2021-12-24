package repositorys

import (
	"context"
	"fmt"
	"log"

	"github.com/Holilamd/module-pemeliharaan/app/models"
	"github.com/Holilamd/module-pemeliharaan/helper"
	"github.com/jackc/pgx/v4"
)

type Repositoryuser interface {
	All() ([]models.User, error)
	Create(input models.Userinput) (models.Userinput, error)
	// Update()
	// FindById()
	FindByEmail(Email string) (models.Userinput, error)
}

type repositoryuser struct {
	db *pgx.Conn
}

func NewRepositoryuser(db *pgx.Conn) *repositoryuser {
	return &repositoryuser{db}
}
func (r *repositoryuser) All() ([]models.User, error) {

	sql := `SELECT id,nama,username,email,locked,csalahpwd,status FROM users`
	rows, err := r.db.Query(context.Background(), sql)

	if err != nil {
		log.Println("Quer All Data Failed %v", err)
	}

	defer rows.Close()
	data := make([]models.User, 0)
	for rows.Next() {
		d := models.User{}
		err := rows.Scan(&d.Id, &d.Nama, &d.Username, &d.Email, &d.Locked, &d.Csalahpwd, &d.Status)
		if err != nil {
			return data, err
		}
		data = append(data, d)
	}
	return data, nil
}
func (r *repositoryuser) Create(input models.Userinput) (models.Userinput, error) {
	tx, err := r.db.Begin(context.Background())
	if err != nil {
		return input, err
	}
	defer tx.Rollback(context.Background())
	sql := `INSERT INTO users (id, nama, username, email, password,status, created_by, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, LOCALTIMESTAMP)`
	head := `US-0`
	id := fmt.Sprintf("%s%d", head, helper.Ctrautonumber(r.db, "users", head))
	input.Id = id
	_, err = tx.Exec(context.Background(), sql, input.Id, input.Nama, input.Username, input.Email, input.Password, true, input.CreatedBy)
	if err != nil {
		return input, err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return input, err
	}
	return input, nil
}

func (r *repositoryuser) FindByEmail(Email string) (models.Userinput, error) {
	var us models.Userinput
	sql := `SELECT id,nama,username,email,locked,csalahpwd,status FROM users WHERE email=$1`
	rows := r.db.QueryRow(context.Background(), sql, Email)
	err := rows.Scan(&us.Id, &us.Nama, &us.Username, &us.Email, &us.Locked, &us.Csalahpwd, &us.Status)
	return us, err
}
