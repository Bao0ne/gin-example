package main

import (
	"fmt"
	"github.com/Bao0ne/gin-example/pkg/setting"
	"github.com/Bao0ne/gin-example/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:		fmt.Sprintf("%v:%d", "127.0.0.1", setting.HTTPPort),
		Handler: 	router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
