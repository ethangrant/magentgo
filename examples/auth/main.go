package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethangrant/magentgo"
)

func main() {
	// init client
	magentogoClient, err := magentgo.New(
		magentgo.WithBaseURl("http://heals.local/"),
		magentgo.WithStoreCode("all"),
		magentgo.WithVersion(1),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	res, err := magentogoClient.AuthService.AdminToken("admin", "admin123", ctx)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(res)
	}

	fmt.Println(res)

	res, err = magentogoClient.AuthService.CustomerToken("admin", "admin123", ctx)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(err)
	}

	if res.Token == "" {
		fmt.Println(res.Message)
	}
}
