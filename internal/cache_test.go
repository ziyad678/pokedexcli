package pokecache

import (
	"testing"
)

func TestCacheAdd(t *testing.T) {
	cache := NewCache(5)
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
			expected: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		},
		{
			input:    "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
			expected: "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
		},
		// add more cases here
	}
	for _, c := range cases {
		cache.Add(c.input,[]byte{})
		_, found := cache.Get(c.input)
		if !found{
			t.Errorf("Test Failed")
			return
		}
		
	}
}
