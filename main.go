// @Author: float311@163.com
// @Description:
// @Version: 1.0.0
// Date: 2021/10/29 13:14
// Update: 2021/10/29 13:14

package main

import (
	"github.com/lzwgiter/Go-Blog/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
