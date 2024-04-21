package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type AddressRepositoryImpl struct {
}

func NewAddressRepositoryImpl() AddressRepository {
	return &AddressRepositoryImpl{}
}

func (repository *AddressRepositoryImpl) AddAddress(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address {
	SQL := "INSERT INTO addresses(user_id, kabupaten, kecamatan, kelurahan, alamat, note) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, address.UserId, address.Kabupaten, address.Kecamatan, address.Kelurahan, address.Alamat, address.Note)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	address.Id = int(id)
	return address
}

func (repository *AddressRepositoryImpl) UpdateAddress(ctx context.Context, tx *sql.Tx, address domain.Address) domain.Address {
	SQL := "UPDATE addresses SET kabupaten = ?, kecamatan = ?, kelurahan = ?, alamat = ?, note = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, address.Kabupaten, address.Kecamatan, address.Kelurahan, address.Alamat, address.Note, address.Id)
	helper.PanicIfError(err)

	return address
}

func (repository *AddressRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Address, error) {
	SQL := `
		SELECT a.id, a.user_id, a.kabupaten, a.kecamatan, a.kelurahan, a.alamat, a.note, u.id
		FROM addresses a
		LEFT JOIN user u ON a.user_id = u.id
		WHERE a.user_id = ?
	`
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []domain.Address
	for rows.Next() {
		address := domain.Address{}
		user := domain.User{}
		if err := rows.Scan(&address.Id, &address.UserId, &address.Kabupaten, &address.Kecamatan, &address.Kelurahan, &address.Alamat, &address.Note, &user.Id); err != nil {
			return nil, err
		}
		address.User = append(address.User, user)
		addresses = append(addresses, address)
	}

	return addresses, nil
}
