package service

import (
	"context"
	"project-workshop/go-api-ecom/model/web"
)

type OrderItemsService interface {
	FindById(ctx context.Context, Id int) web.OrderItemsResponse
	FindByCartId(ctx context.Context, cartId int) []web.OrderItemsResponse
	FindByUserId(ctx context.Context, userId int) []web.OrderItemsResponse
	// CreateOrderItems(ctx context.Context, request web.OrderItemsCreateRequest, userId int) web.OrderItemsResponse
}
