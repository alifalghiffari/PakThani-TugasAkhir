package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type CartRepositoryImpl struct {
}

func NewCartRepositoryImpl() CartRepository {
	return &CartRepositoryImpl{}
}

func (repository *CartRepositoryImpl) AddToCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "insert into cart(userId, product_id, quantity, price) values (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, cart.UserId, cart.ProductId, cart.Quantity, cart.Price)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	cart.Id = int(id)
	return cart
}

func (repository *CartRepositoryImpl) UpdateCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "update cart set quantity = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.Quantity, cart.Id)
	helper.PanicIfError(err)

	return cart
}

func (repository *CartRepositoryImpl) DeleteCart(ctx context.Context, tx *sql.Tx, cart domain.Cart) domain.Cart {
	SQL := "DELETE FROM cart WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, cart.Id)
	helper.PanicIfError(err)

	return cart
}

// FindById fetches cart items based on a slice of cart IDs
func (repository *CartRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, cartId []int) ([]domain.Cart, error) {
	SQL := `
		SELECT c.id AS cart_id, c.userId, c.product_id, c.quantity, c.price, p.id AS product_id, p.name, p.price, p.image, p.description, p.category_id, cat.id AS category_id, cat.category
		FROM cart c
		LEFT JOIN product p ON c.product_id = p.id
		LEFT JOIN category cat ON p.category_id = cat.id
		WHERE c.id = ?
	`
	rows, err := tx.QueryContext(ctx, SQL, cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		product := domain.Product{}
		category := domain.Category{}
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.ProductId, &cart.Quantity, &cart.Price, &product.Id, &product.Name, &product.Price, &product.Image, &product.Description, &product.CategoryId, &category.Id, &product.Category.Category)
		helper.PanicIfError(err)
		cart.Product = append(cart.Product, product)
		carts = append(carts, cart)
	}

	return carts, nil
}

func (repository *CartRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Cart, error) {
	SQL := `
		SELECT c.id AS cart_id, c.userId, c.product_id, c.quantity, c.price, p.id AS product_id, p.name, p.price, p.image, p.description, p.category_id, cat.id AS category_id, cat.category
		FROM cart c
		LEFT JOIN product p ON c.product_id = p.id
		LEFT JOIN category cat ON p.category_id = cat.id
		WHERE c.userId = ? AND c.deleted_at IS NULL

    `
	rows, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		product := domain.Product{}
		category := domain.Category{}
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.ProductId, &cart.Quantity, &cart.Price, &product.Id, &product.Name, &product.Price, &product.Image, &product.Description, &product.CategoryId, &category.Id, &product.Category.Category)
		helper.PanicIfError(err)
		cart.Product = append(cart.Product, product)
		carts = append(carts, cart)
	}

	return carts, nil
}

func (repository *CartRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Cart {
	SQL := "select id, userId, product_id, quantity, price from cart"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var carts []domain.Cart
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.Id, &cart.UserId, &cart.ProductId, &cart.Quantity, &cart.Price)
		helper.PanicIfError(err)
		carts = append(carts, cart)
	}

	return carts
}
