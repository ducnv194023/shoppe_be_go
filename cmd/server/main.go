package main

import (
	"shoppe_be/internal/routers"
)
func main() {
	r:= routers.NewRouter()
	r.Run(":8080")
}

