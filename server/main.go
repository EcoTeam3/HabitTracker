package main

import (
	"mymode/config"
	"mymode/service"
	"mymode/storage"
	"mymode/storage/postgres"
	"net"
	"google.golang.org/grpc"
	pb "mymode/generated/habit"
)



func main(){
	db,err := storage.Connect()
	if err != nil{
		panic(err)
	}
	defer db.Close()

	config := config.Load()
	listener,err := net.Listen("tcp",":"+config.URL_PORT)

	if err != nil{
		panic(err)
	}

	defer listener.Close()

	s := service.NewHabitTracker(*postgres.NewHabitTrackerRepo(db))
	server := grpc.NewServer()
	pb.RegisterHabitTrackerServer(server, s)
	
}