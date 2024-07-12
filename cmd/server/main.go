package main

import "github.com/banhmysuawx/Go-ecommerce-backend-api/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run(":8080")
}

