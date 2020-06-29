package main

import (
	"context"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	res, err := Handler(context.Background())

	if err != nil {
		t.Logf("unexpected error (%v)", err)
		t.Fail()
		return
	}

	if res.StatusCode != http.StatusOK {
		t.Logf("unexpected status code: (%v) expected (%v)", res.StatusCode, http.StatusOK)
		t.Fail()
		return
	}
}