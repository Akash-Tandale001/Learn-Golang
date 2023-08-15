package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

type apiFunc func(http.ResponseWriter , *http.Request) error
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc{
	 return func(w http.ResponseWriter, r * http.Request){
		if err :=f(w,r); err !=nil{
			
		}
	 }
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer)Run()  {
	router := mux.NewRouter()
	router.HandleFunc("/account",s.handleAccount)
}

func (s *APIServer) handleAccount(w http.ResponseWriter ,r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter ,r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter ,r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter ,r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter ,r *http.Request) error {
	return nil
}