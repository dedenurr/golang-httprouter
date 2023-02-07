package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS
func TestServeFile(t *testing.T) {
	router := httprouter.New()
	//tujuan directory
	directory, _ := fs.Sub(resources,"resources")
	//alamat url
	router.ServeFiles("/files/*filepath",http.FS(directory))

	//test request langsung ke file hello.txt
	request := httptest.NewRequest("GET","http://localhost:3000/files/hello.txt",nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	assert.Equal(t, "Hellow", string(body))
}

func TestServeFileGoodye(t *testing.T) {
	router := httprouter.New()
	//tujuan directory
	directory, _ := fs.Sub(resources,"resources")
	//alamat url
	router.ServeFiles("/files/*filepath",http.FS(directory))

	//test request langsung ke file hello.txt
	request := httptest.NewRequest("GET","http://localhost:3000/files/goobye.txt",nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder,request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	assert.Equal(t, "Bye-bye", string(body))
}