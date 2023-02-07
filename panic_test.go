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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	// digunakan untuk menghandle sebuah proses yang error pada website
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer,"Panic : ",error)
	}
	router.GET("/",func(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
		panic("ups")
	})

	request := httptest.NewRequest("GET","http://localhost:3000/",nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : ups", string(body))
}