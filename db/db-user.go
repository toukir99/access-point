package db

import (
	"access-point/web/model"
	"database/sql"
	"log/slog"

	sq "github.com/Masterminds/squirrel"
)

type UserRepo struct {
	table string
}

var userRepo *UserRepo 

func InitUserRepo() {
	userRepo = &UserRepo{
		table: "users",
	}
}

func GetUserRepo() *UserRepo{
	return userRepo
}

func (r *UserRepo) CreateUser(user *model.User) (int, error) {
	query, args, err := GetQueryBuilder().
		Insert(r.table).
		Columns("username", "email", "password").
		Values(user.Username, user.Email, user.Password).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		slog.Error("Failed to create user insert query", "err", err)
		return 0, err
	}

	var newID int
	err = GetWriteDB().QueryRow(query, args...).Scan(&newID)
	if err != nil {
		slog.Error("Error executing create user query", "err", err)
		return 0, err
	}

	return newID, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	query, args, err := GetQueryBuilder().
		Select("id", "username", "email", "password").
		From(r.table).
		Where(sq.Eq{"email": email}).
		ToSql()
	if err != nil {
		slog.Error("Failed to create get user by email query", "err", err)
		return nil, err
	}

	var user model.User
	err = GetReadDB().QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		slog.Error("Error getting user by email", "err", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetAllUsers() ([]*model.UserResponse, error) {
	query, args, err := GetQueryBuilder().
		Select("id", "username", "email").
		From(r.table).
		OrderBy("id ASC").
		ToSql()
	if err != nil {
		slog.Error("Failed to create get all users query", "err", err)
		return nil, err
	}

	var users []*model.UserResponse
	err = GetReadDB().Select(&users, query, args...)
	if err != nil {
		slog.Error("Error fetching users", "err", err)
		return nil, err
	}

	return users, nil
}


func (r *UserRepo) GetUserByID(id int) (*model.UserResponse, error) {
	query, args, err := GetQueryBuilder().
		Select("id", "username", "email").
		From(r.table).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		slog.Error("Failed to create get user by ID query", "err", err)
		return nil, err
	}

	var user model.UserResponse
	err = GetReadDB().QueryRow(query, args...).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		slog.Error("Error getting user by ID", "err", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) UpdateUser(id int, user *model.User) error{
	query, args, err := GetQueryBuilder().
		Update(r.table).
		Set("username", user.Username).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		slog.Error("Failed to create user update!")
		return err
	}

	_, err = GetWriteDB().Exec(query, args...)
	if err != nil {
		slog.Error("Error updating user", "error", err.Error())
		return err
	}

	return nil
}

func (r *UserRepo) DeleteUser(id int) error {
	query, args, err := GetQueryBuilder().
		Delete(r.table).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		slog.Error("Failed to create user deletion!")
		return err
	}

	_, err = GetWriteDB().Exec(query, args...)
	if err != nil {
		slog.Error("Error deleting user", "error", err.Error())
		return err
	}

	return nil
}