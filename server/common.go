package server

import (
	"NFTLookingGlass/util"
	"log"
	"net/http"
)

func InitServer() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/nft", nftHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	util.ServeJson(w, "health: ok")
}

func nftHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getNftHandler(w, r)
	default:
		util.ServeBadRequest(w)
	}
}
