package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T){
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not httpHandler, but is %t" , v))


	}
}
func TestSessionLoad(t *testing.T){
	var myH myHandler

	h := sessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("type is not httpHandler, but is %t" , v))


	}
}

