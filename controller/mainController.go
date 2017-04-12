package controller

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"io/ioutil"
)

//TODO: define where to put these structs
type suma_req struct {
	A      int32    `json:"a"`
	B      int32    `json:"b"`
}
type suma_response struct {
	Response int32 `json:"response"`
}


func setup(r *http.Request){//TODO: hide the calls to this.
	log.Println(r.URL, "requested. From:",r.RemoteAddr)//TODO: change this to log.
	log.Println("Parameters: ",mux.Vars(r))
}


func Index(w http.ResponseWriter, r *http.Request) {
	setup(r)
	fmt.Fprintln(w, "Welcome controller!")
}

/*func Message(w http.ResponseWriter, r *http.Request){
	setup(r)
	var resp Resp
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &resp)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
	//w.Header().Set("content-type", "application/json")
}*/

func Suma(w http.ResponseWriter, r *http.Request) {
	setup(r)
	var req suma_req
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &req)
	result :=req.A+req.B

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suma_response{Response:result})
}