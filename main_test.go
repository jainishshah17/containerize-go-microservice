package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T)  {
	req := httptest.NewRequest("GET", "http://example.com/ping", nil)
	w := httptest.NewRecorder()
	hello(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Error("Got wrong response!")
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}