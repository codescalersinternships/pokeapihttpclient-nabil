package client

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestGetPokemonByName(t *testing.T) {
	t.Run("valid req", func(t *testing.T) {
		client := NewClient(time.Duration(1) * time.Second)
		poke, err := client.GetPokemonByname(context.Background(), "pikachu")
		assertNotError(t, err)
		if !reflect.DeepEqual("pikachu", poke.Name) {
			t.Errorf("got %q, want %q", poke.Name, "pikachu")
		}
	})
	t.Run("unvalid req name", func(t *testing.T) {
		client := NewClient(time.Duration(1) * time.Second)
		_, err := client.GetPokemonByname(context.Background(), "un-exist")
		assertError(t, err)

	})
	t.Run("listing all pokemons", func(t *testing.T) {
		client := NewClient(time.Duration(1) * time.Second)
		_, err := client.GetPokemonList(context.Background(), 10, 1)
		assertNotError(t, err)

	})

}

func assertNotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("error isn't nill error: %q", err)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("error is nill error: %q", err)
	}
}
