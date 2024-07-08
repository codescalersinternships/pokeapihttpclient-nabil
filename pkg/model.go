package client

// Pokemon class is used as model to parse response data for one pokemon
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
}

// Pokemon class is used as model to parse response data for array of pokemons
type pokeList struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
