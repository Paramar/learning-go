package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	http.HandleFunc("/", headerHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":808", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	//当访问 localhost/healthz 时，应返回 200
	io.WriteString(w, "200\n")
}

func headerHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "------接收客户端 request，并将 request 中带的 header 写入 response header:--------\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}

	log.Println("--------读取当前系统的环境变量中的 VERSION 配置，并写入 response header:-------------------")
	os.Setenv("VERSION","1.0")
	envVersoin := os.Getenv("VERSION")
	w.Header().Set("VERSION",envVersoin)

	log.Println("--------Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出:-------------")
	statusCode :=200
	w.WriteHeader(statusCode)
	remoteAddr := r.RemoteAddr
	log.Println("客户端 IP: ",remoteAddr,"\n HTTP 返回码: ",statusCode)

}
