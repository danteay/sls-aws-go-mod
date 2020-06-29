package utils

import (
	"net/http"
	"testing"
)

func TestNewResponse(t *testing.T) {
	res, err := NewResponse(http.StatusOK, nil)
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
