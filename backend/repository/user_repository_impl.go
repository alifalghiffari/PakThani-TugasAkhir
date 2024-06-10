package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into user(username, password, email, role, no_telepon) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Email, user.Role, user.NoTelepon)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user.Id = int(id)
	return user
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "select id, username, password, email, no_telepon, role from user where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role)
		if err != nil {
			panic(err)
		}
		return user
	} else {
		panic("email or password is incorrect")
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select id, username, password, email, no_telepon, role from user where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "select id, username, password, email, no_telepon, role from user"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindByRole(ctx context.Context, tx *sql.Tx, role bool) (domain.User, error) {
	SQL := "select id, username, password, email, no_telepon, role from user where role = ?"
	rows, err := tx.QueryContext(ctx, SQL, role)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "select id, username, password, email, no_telepon, role from user where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "select id, username, password, email, no_telepon, role from user where email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.NoTelepon, &user.Role)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {
		return user, err
	}
}

func (repository *UserRepositoryImpl) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "update user set password = ? where email = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Email)
	if err != nil {
		panic(err)
	}

	return user
}
