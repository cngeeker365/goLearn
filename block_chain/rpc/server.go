package main

import (
	"awesomeProject/block_chain/core"
	"encoding/json"
	"net/http"
)

var blockChain *core.BlockChain

func run(){
	http.HandleFunc("/blockchain/get", blockChainGetHandler)
	http.HandleFunc("/blockchain/put", blockChainPutHandler)
	http.ListenAndServe(":8080", nil)
}

func blockChainGetHandler(w http.ResponseWriter, r *http.Request){
	data, err := json.MarshalIndent(blockChain, "", "\t")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}

func blockChainPutHandler(w http.ResponseWriter, r *http.Request){
	data := r.URL.Query().Get("data")
	blockChain.SendData(data)
	blockChainGetHandler(w, r)
}

func main(){
	blockChain = core.NewBlockChain()
	run()
}