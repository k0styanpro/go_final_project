package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Run() error {
	port := getPort()
	fmt.Printf("Starting server on port %d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getPort() int {
	if port, err := strconv.Atoi(os.Getenv("TODO_PORT")); err == nil {
		return port
	}
	return 7540
}
