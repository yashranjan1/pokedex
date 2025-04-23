package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next string
	prev string
}

type APIResponse struct {
	Next     *string `json:"next"`     // Pointer to string to allow for `null` values
	Previous *string `json:"previous"` // Same here
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
