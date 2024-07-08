package main

import (
	"context"
	"fmt"

	"github.com/ethangrant/magentgo"
)

func main() {
	magentgoClient, err := magentgo.New(
		magentgo.WithBaseURl("http://heals.local/"),
		magentgo.WithStoreCode("all"),
	)
	if err != nil {
		fmt.Println(err.Error())
		return;
	}

	ctx := context.Background();

	_, err = magentgoClient.AuthService.AdminToken("admin", "admin123", ctx)
	if err != nil {
		fmt.Println(err.Error())
		return;
	}

	res, err := magentgoClient.ProductService.GetBySku("MG0041587", ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)

	res, err = magentgoClient.ProductService.GetById(244530, ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)

	// res, err = magentgoClient.ProductService.
}