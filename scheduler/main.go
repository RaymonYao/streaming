package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"streaming/config"
	"streaming/logger"
	"streaming/scheduler/handler"
	"streaming/scheduler/model"
	"streaming/scheduler/taskrunner"
	"strings"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", handler.VidDelRecHandler)
	return router
}

func main() {
	// 初始化配置
	config.InitConfig("./conf.json")
	fmt.Printf("%+v\n", config.DefaultConfig)

	// 日志配置
	fmt.Println("logger init...")
	path := "logs"
	mode := os.ModePerm
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
	file, _ := os.Create(strings.Join([]string{path, "scheduler_log.txt"}, "/"))
	defer file.Close()
	loger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.SetDefault(loger)

	// 初始化DB
	fmt.Println("mysql init...")
	model.InitMysql()

	go taskrunner.Start()
	r := RegisterHandlers()
	addr := config.DefaultConfig.Address + ":" + config.DefaultConfig.SchedulerPort
	fmt.Println("streamServer start...\t", addr)
	http.ListenAndServe(addr, r)
}
