/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/15 23:04
*/
package main

import (
	"os"
	"strings"
	"io"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/return_header", returnHeader)
	http.HandleFunc("/return_local_env", returnLocalEnv)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80",nil)
	if err != nil{
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request){
	w.Header().Add("status_code", "200")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w,"healthz check")
	printLog(r)
}

func returnHeader(w http.ResponseWriter, r *http.Request){
	for k,v := range r.Header{
		//fmt.Printf("k: %T ,v: %T\n", k,v)
		for _,value := range v{
			w.Header().Add(k,value)
		}
	}
	io.WriteString(w, "return custom header")
	printLog(r)
}

func returnLocalEnv(w http.ResponseWriter, r *http.Request){
	os.Setenv("VERSION","1.0.0-test")
	value := os.Getenv("VERSION")
	w.Header().Add("VERSION", value)

	io.WriteString(w,"return local version env.")
	printLog(r)
}


func printLog(r * http.Request){
	ip := strings.Split(r.RemoteAddr, ":")[0]
	status := http.StatusOK
	log.Printf("remote ip: %s, status: %d\n",ip,status)
}


