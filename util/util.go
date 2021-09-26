package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func EnableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")
}

func GetHeader(r *http.Request, key string) string {
	return r.Header.Get(key)
}

func GetURLParam(r *http.Request, name string) string {
	query := r.URL.Query()
	params := query[name]
	if len(params) == 0 {
		return ""
	}
	return params[0]
}

func GetURLBoolParam(r *http.Request, name string) bool {
	paramStr := GetURLParam(r, name)
	return paramStr == "1"
}

func GetURLIntParam(r *http.Request, name string) int {
	paramStr := GetURLParam(r, name)
	param, err := strconv.Atoi(paramStr)
	if err != nil {
		fmt.Println("error")
	}
	return param
}

func GetBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error")
	}
	return body
}

func ServeJson(w http.ResponseWriter, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("error")
	}

	_, err = w.Write(data)
}

func ServeCreated(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
}

func ServeNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func ServeBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func ServeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
}

func ServeForbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
}

func ServeNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func ServeConflict(w http.ResponseWriter) {
	w.WriteHeader(http.StatusConflict)
}

func ServeInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}
