package main

import (
	"context"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	res, err := Handler(context.Background())
	if err != nil {
		t.Fail()
		return
	}

	if res.StatusCode != http.StatusOK {
		t.Log("unexpected status code")
		t.Fail()
		return
	}
}