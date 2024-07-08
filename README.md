# Test HTTP Client for pokemon api Implemented in Go
Create an HTTP client in Go that consumes the Pokemon server APIs. This client will focus on building an HTTP client development and testing.

# Installation

To install the project use:

```golang
go get github.com/codescalersinternships/pokeapihttpclient-nabil
```

## Usage

1. Import the Package
```golang
import Client "github.com/codescalersinternships/pokeapihttpclient-nabil/pkg"
```

Using url string

2. Create a New Client Instance
```golang
client := Client.NewClient(time.Duration(5) * time.Second)
```

3. Get date and time
```golang
poke, err := client.GetPokemonList(context.Background(), 10, 1)
```


## How to Test

```golang
make test
```