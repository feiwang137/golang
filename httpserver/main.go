/*
@Description: xxxx
@Author: feiwang
@Date: 2023/5/15 23:04
*/
package main

import (
	"io"
	"log"
	"net/http"
)

func main()  {
	// 注册handle处理函数
	http.HandleFunc("/healthz",healthz)
	err := http.ListenAndServe(":80",nil)
	if err != nil{
		log.Fatal(err)

	}
}

func healthz(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"ok")
}

