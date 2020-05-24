package main

import (
	"fmt"

	"github.com/mrkvm/rados.go"
)

func main() {
	r, err := rados.New("")
	if err != nil {
		panic(err)
	}

	fmt.Println("* create pool: test-bucket")
	if err = r.CreatePool("test-bucket"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("* list pools")
	pools, _ := r.ListPools()
	fmt.Println(pools)
}

