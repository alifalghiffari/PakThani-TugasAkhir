package repository

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepositoryImpl() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Order {
	query := "SELECT id, user_id, total_items, total_price, order_status, payment_status FROM orders"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		err := rows.Scan(
			&order.Id, &order.UserId, &order.OrderStatus, &order.PaymentStatus,
		)
		helper.PanicIfError(err)
		orders = append(orders, order)
	}
	return orders
}

func (repository *OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderId int) (domain.Order, error) {
	query := "SELECT id, user_id, total_items, total_price, order_status, payment_status FROM orders WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, orderId)
	helper.PanicIfError(err)
	defer rows.Close()

	var order domain.Order
	if rows.Next() {
		err := rows.Scan(
			&order.Id, &order.UserId, &order.OrderStatus, &order.PaymentStatus,
		)
		helper.PanicIfError(err)
		return order, nil
	} else {
		return order, nil
	}
}

func (repository *OrderRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) ([]domain.Order, error) {
	query := "SELECT id, user_id, total_items, total_price, order_status, payment_status FROM orders WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		err := rows.Scan(
			&order.Id, &order.UserId, &order.OrderStatus, &order.PaymentStatus,
		)
		helper.PanicIfError(err)
		orders = append(orders, order)
	}
	return orders, nil
}

func (repository *OrderRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	query := "INSERT INTO orders(user_id, total_items, total_price, order_status, payment_status) VALUES(?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, order.UserId, order.TotalItems, order.TotalPrice, order.OrderStatus, order.PaymentStatus)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.Id = int(id)
	return order
}

func (repository *OrderRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	query := "UPDATE orders SET order_status = ?, payment_status = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, order.OrderStatus, order.PaymentStatus, order.Id)
	helper.PanicIfError(err)

	return order
}
