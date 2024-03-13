package services

import (
	"github.com/danjelhysenaj-dev/bookstore_items-api/domain/items"
	"github.com/danjelhysenaj-dev/bookstore_utils-go/rest_errors"
	"net/http"
)

// 5 Make the ItemsService available in other packages
var (
	ItemsService itemsServiceInterface = &itemsService{}
)

// 1
type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

// 2
type itemsService struct {
}

// 3
func (s *itemsService) Create(items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}

// 4
func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me!", http.StatusNotImplemented, "not_implemented", nil)
}
