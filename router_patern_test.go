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

// 1. Named parameter adalah pola pembuatan parameter dengan menggunakan nama
/* contoh /user/:user artinya /user/eko */
func TestRouterPaternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId",func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET","http://localhost:3000/products/1/items/1",nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

// catch all parameter, yaitu menangkap semua parameter
/* contoh /src/*filepath artinya bisa /src/somefile atau bisa gini /src/subdir/somefile asal tidak boleh hanya /src/  */

func TestRouterPaternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image",func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET","http://localhost:3000/images/small/profile.png",nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}