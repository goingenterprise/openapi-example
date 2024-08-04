package main

import (
	"context"
	petstoregosdk "github.com/goingenterprise/petstore-go-sdk"
	"log"
)

func main() {
	s := petstoregosdk.New(petstoregosdk.WithServerURL("http://localhost:8080"))
	var limit = petstoregosdk.Int(21453)
	ctx := context.Background()
	res, err := s.Pets.ListPets(ctx, limit)
	if err != nil {
		log.Fatal(err)
	}
	if res.Pets != nil {
		log.Println(res.Pets)
	}
}
