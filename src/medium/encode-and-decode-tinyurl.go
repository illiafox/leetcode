package medium

import (
	"math/rand"
)

type Codec struct {
	m map[string]string
}

func Constructor() Codec {
	return Codec{make(map[string]string)}
}

var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	code := make([]byte, 8)
	for i := range code {
		code[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	id := string(code)

	c.m[id] = longUrl
	return id
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.m[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
