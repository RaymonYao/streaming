package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"streaming/config"
	"streaming/logger"
	"streaming/web/handler"
	"strings"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", handler.HomeHandler)
	router.POST("/", handler.HomeHandler)

	router.GET("/userhome", handler.UserHomeHandler)
	router.POST("/userhome", handler.UserHomeHandler)

	router.POST("/api", handler.ApiHandler)

	router.GET("/videos/:vid-id", handler.ProxyVideoHandler)
	router.POST("/upload/:vid-id", handler.ProxyUploadHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}

func main() {
	// 初始化配置
	config.InitConfig("../config/conf.json")
	fmt.Printf("%+v\n", *config.DefaultConfig)

	// 日志配置
	fmt.Println("logger init...")
	path := "logs"
	mode := os.ModePerm
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			panic(err)
		}
	}
	file, _ := os.Create(strings.Join([]string{path, "web_log.txt"}, "/"))
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	_logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetDefault(_logger)

	r := RegisterHandler()

	addr := config.DefaultConfig.Address + ":" + config.DefaultConfig.WebPort
	fmt.Println("handler init...Port:\t", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		return
	}
}
