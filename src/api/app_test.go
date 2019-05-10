package main

import (
	"net/http"
	"testing"
)

const url = "http://localhost:8080/user/123"

func BenchmarkApp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		http.Get(url)
	}
}