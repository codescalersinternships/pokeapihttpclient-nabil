package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetPokemonByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"name": "pikachu"}`)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	t.Run("valid req with mock", func(t *testing.T) {
		client := NewClient(time.Duration(1)*time.Second, WithApiUrl(server.URL))
		poke, err := client.GetPokemonByname(context.Background(), "pikachu")
		assertNotError(t, err)
		if !reflect.DeepEqual("pikachu", poke.Name) {
			t.Errorf("got %q, want %q", poke.Name, "pikachu")
		}
	})

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
	t.Run("withapirltest", func(t *testing.T) {
		client := NewClient(time.Duration(1)*time.Second, WithApiUrl("hello"))
		if client.apiUrl != "hello" {
			t.Errorf("clint url wanted: %q, got %q", "hello", client.apiUrl)
		}

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
