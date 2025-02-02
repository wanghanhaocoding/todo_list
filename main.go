package main

import (
	"fmt"
	"todo_list/conf"
	"todo_list/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	fmt.Println(conf.HttpPort)
	err := r.Run(conf.HttpPort)
	if err != nil {
		panic(err)
	}
}
