package service

import (
	"context"
	"database/sql"
	"project-workshop/go-api-ecom/exception"
	"project-workshop/go-api-ecom/helper"
	// "project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
)

type OrderItemsServiceImpl struct {
	OrderItemsRepository repository.OrderItemsRepository
	UserRepository       repository.UserRepository
	CartRepository       repository.CartRepository
	OrderRepository      repository.OrderRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewOrderItemsService(orderItemsRepository repository.OrderItemsRepository, userRepository repository.UserRepository, cartRepository repository.CartRepository, orderRepository repository.OrderRepository, DB *sql.DB, validate *validator.Validate) OrderItemsService {
	return &OrderItemsServiceImpl{
		OrderItemsRepository: orderItemsRepository,
		UserRepository:       userRepository,
		CartRepository:       cartRepository,
		OrderRepository:      orderRepository,
		DB:                   DB,
		Validate:             validate,
	}
}

func (service *OrderItemsServiceImpl) FindById(ctx context.Context, Id int) web.OrderItemsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderItems, err := service.OrderItemsRepository.FindById(ctx, tx, Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderItemsResponse(orderItems)
}

func (service *OrderItemsServiceImpl) FindByCartId(ctx context.Context, orderId int) []web.OrderItemsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderItems, err := service.OrderItemsRepository.FindByOrderId(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderItemsResponses(orderItems)
}

func (service *OrderItemsServiceImpl) FindByUserId(ctx context.Context, userId int) []web.OrderItemsResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderItems, err := service.OrderItemsRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderItemsResponses(orderItems)
}

// func (service *OrderItemsServiceImpl) CreateOrderItems(ctx context.Context, request web.OrderItemsCreateRequest, userId int) web.OrderItemsResponse {
// 	err := service.Validate.Struct(request)
// 	helper.PanicIfError(err)
	
// 	tx, err := service.DB.Begin()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)

// 	user, err := service.UserRepository.FindById(ctx, tx, userId)
// 	if err != nil {
// 		panic(exception.NewNotFoundError(err.Error()))
// 	}

// 	cartItems, err := service.CartRepository.FindById(ctx, tx, request.CartId)
// 	if err != nil {
// 		panic(exception.NewNotFoundError(err.Error()))
// 	}

// 	var totalItems int
// 	var totalPrice int
// 	var orderItems []domain.OrderItems
// 	var cartResponses []web.CartResponse

// 	for _, cartItem := range cartItems {
// 		orderItem := domain.OrderItems{
// 			ProductId: cartItem.ProductId,
// 			Quantity:  cartItem.Quantity,
// 			Price:     cartItem.Price,
// 		}
// 		orderItems = append(orderItems, orderItem)

// 		totalItems += cartItem.Quantity
// 		totalPrice += cartItem.Price * cartItem.Quantity

// 		cartResponse := web.CartResponse{
// 			Id:        cartItem.Id,
// 			ProductId: cartItem.ProductId,
// 			Quantity:  cartItem.Quantity,
// 			Price:     cartItem.Price,
// 		}
// 		cartResponses = append(cartResponses, cartResponse)

// 		// Delete the cart item after adding to order
// 		service.CartRepository.DeleteCart(ctx, tx, cartItem)
// 	}

// 	// You can use the calculated totalItems and totalPrice to create the order
// 	orderStatus := mapToOrderStatus("PENDING")     // Set the default order status
// 	paymentStatus := mapToPaymentStatus("PENDING") // Set the default payment status

// 	order := domain.Order{
// 		UserId:        user.Id,
// 		TotalItems:    totalItems,
// 		TotalPrice:    totalPrice,
// 		OrderStatus:   orderStatus,
// 		PaymentStatus: paymentStatus,
// 	}

// 	// Insert the order into the database and get the Order ID
// 	order, err = service.OrderRepository.Insert(ctx, tx, order)
// 	if err != nil {
// 		panic(exception.NewNotFoundError(err.Error()))
// 	}

// 	// Update OrderItems with the actual Order ID
// 	for i := range orderItems {
// 		orderItems[i].OrderId = order.Id
// 	}

// 	// Insert each OrderItem into the database
// 	for _, item := range orderItems {
// 		_, err = service.OrderItemsRepository.Insert(ctx, tx, item)
// 		if err != nil {
// 			panic(exception.NewNotFoundError(err.Error()))
// 		}
// 	}

// 	return helper.ToOrderItemsResponse(orderItems)
// }
