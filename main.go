package main

import (
	"fmt"
	"github.com/ichyycc/gin-demo/pkg/setting"
	"github.com/ichyycc/gin-demo/routers"
	"net/http"
)

func main() {

	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}

}
