package main

import (
	"mymode/config"
	pb "mymode/generated"
	"mymode/service"
	"mymode/storage"
	"mymode/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	config := config.Load()
	listener, err := net.Listen("tcp", ":"+config.URL_PORT)

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	s := service.NewHabitTracker(*postgres.NewHabitTrackerRepo(db))
	server := grpc.NewServer()
	pb.RegisterHabitTrackerServer(server, s)
	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
