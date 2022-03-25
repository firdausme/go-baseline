package repository

import (
	"database/sql"
	"go-baseline/entity"
	"go-baseline/exception"
)

const (
	FIND_ALL         = "SELECT id, username, password FROM user"
	FIND_BY_USERNAME = "SELECT id, username, password FROM user WHERE username = ?"
	INSERT           = "INSERT INTO user (id, username, password) VALUES (?, ?, ?)"
	UPDATE           = "UPDATE user SET password = ? WHERE id = ?"
	DELETE           = "DELETE FROM user where id = ?"
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{
		DB: db,
	}
}

type userRepositoryImpl struct {
	DB *sql.DB
}

func (repo userRepositoryImpl) FindAll() (users []entity.User) {
	var user entity.User
	rows, err := repo.DB.Query(FIND_ALL)
	exception.PanicIfNeeded(err)

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		exception.PanicIfNeeded(err)
		users = append(users, user)
	}

	return users
}

func (repo userRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	var user entity.User

	err := repo.DB.QueryRow(FIND_BY_USERNAME, username).Scan(
		&user.Id, &user.Username, &user.Password)

	return user, err
}

func (repo userRepositoryImpl) Create(user entity.User) {
	stmt, err := repo.DB.Prepare(INSERT)
	exception.PanicIfNeeded(err)

	_, err = stmt.Exec(user.Id, user.Username, user.Password)
	exception.PanicIfNeeded(err)
}

func (repo userRepositoryImpl) Update(user entity.User) {
	stmt, err := repo.DB.Prepare(UPDATE)
	exception.PanicIfNeeded(err)

	_, err = stmt.Exec(user.Password, user.Id)
	exception.PanicIfNeeded(err)
}

func (repo userRepositoryImpl) Delete(id string) {
	stmt, err := repo.DB.Prepare(DELETE)
	exception.PanicIfNeeded(err)

	_, err = stmt.Exec(id)
	exception.PanicIfNeeded(err)
}
