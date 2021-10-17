package controllers

import (
	"fmt"
	"net/http"

	u "api-gateway/utils"

	"github.com/gorilla/mux"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	fmt.Println("In DefaultHandler")
	fmt.Println(params)

	defaultResponse := make(map[string]interface{})

	defaultResponse["apiGateway"] = map[string]interface{}{"version": 1}
	u.Respond(w, defaultResponse)
}
