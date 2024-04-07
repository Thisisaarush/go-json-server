package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	resondWithError(w, r, 400, "Something went wrong!")
}