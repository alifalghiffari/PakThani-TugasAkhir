package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderItemsController interface {
	FindOrderItemsById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindOrderItemsByCartId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindOrderItemsByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	// CreateOrderItems(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
