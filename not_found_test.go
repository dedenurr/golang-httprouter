package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	// digunakan untuk menghandle jika halaman website tidak ada
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "gak ketemu")
	})
	request := httptest.NewRequest("GET","http://localhost:3000/404",nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	assert.Equal(t, "gak ketemu", string(body))
}