package main

import (
	"context"
	"fmt"
	"time"

	client "github.com/codescalersinternships/pokeapihttpclient-nabil/pkg"
)

func main() {
	client := client.NewClient(time.Duration(5) * time.Second)
	poke, _ := client.GetPokemonList(context.Background(), 10, 1)
	fmt.Println(poke)

}
