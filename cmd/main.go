package main

import (
	"backend/doantotnghiep/infrastructure"
	"backend/doantotnghiep/router"
	"fmt"
	"log"
	"net/http"
)

func recoverPanic() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		fmt.Println(err)
	}
}

type fn func()

func runJobWithRecoverPanic(f fn) {
	go func() {
		defer recoverPanic()
		f()
	}()
}

func main() {
	// go run main.go
	log.Println("Database name: ", infrastructure.GetDBName())
	log.Printf("Server running at port: %+v\n", infrastructure.GetAppPort())
	log.Fatal(http.ListenAndServe(":"+infrastructure.GetAppPort(), router.Router()))
}
