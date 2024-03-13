package controllers

import "net/http"

const (
	pong = "pong"
)

// 4
var (
	PingController pingControllerInterface = &pingController{}
)

// 3
type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

// 1
type pingController struct{}

// 2
func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
