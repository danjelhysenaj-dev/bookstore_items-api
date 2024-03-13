package app

import (
	"github.com/danjelhysenaj-dev/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.Create).Methods(http.MethodPost)
}
