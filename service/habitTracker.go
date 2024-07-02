package service

import (
	"context"
	pb "mymode/generated"
	"mymode/storage/postgres"
)

type Server struct {
	pb.UnimplementedHabbitTrackerServer
	H postgres.HabitTrackerRepo
}

func NewHabitTracker(h postgres.HabitTrackerRepo) *Server {
	return &Server{H: h}
}

func (S *Server) CreateHabit(ctx context.Context, habit *pb.Habbit) (*pb.Status, error) {
	status, err := S.H.CreateHabit(habit)
	if err != nil {
		return nil, err
	}
	return status, err
}

func (S *Server) GetHabit(ctx context.Context, habitId *pb.HabbitId) (*pb.Habbit, error) {
	habit, err := S.H.GetHabit(habitId)
	if err != nil {
		return nil, err
	}
	return habit, nil
}

func (S *Server) UpdateHabit(ctx context.Context, habit *pb.Habbit) (*pb.Status, error) {
	status, err := S.H.UpdateHabit(habit)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (S *Server) DeleteHabit(ctx context.Context, habitId *pb.HabbitId) (*pb.Status, error) {
	status, err := S.H.DeleteHabit(habitId)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (S *Server) GetUserHabits(ctx context.Context, userID *pb.UserId) (*pb.Habbit, error) {
	habit, err := S.H.GetUserHabits(userID)
	if err != nil {
		return nil, err
	}
	return habit, nil
}

func (S *Server) CreateHabitLog(ctx context.Context, habitlog *pb.HabbitLog) (*pb.Status, error) {
	status, err := S.H.CreateHabitLog(habitlog)
	if err != nil {
		return nil, err
	}
	return status, err
}

func (S *Server) GetHabitLogs(ctx context.Context, habitId *pb.HabbitId) (*pb.HabbitLog, error) {
	habit, err := S.H.GetHabitLogs(habitId)
	if err != nil {
		return nil, err
	}
	return habit, nil
}

func (S *Server) GetHabitSuggestions(ctx context.Context, req *pb.Req) (*pb.Habbits, error) {
	habits, err := S.H.GetHabitSuggestions(req)
	if err != nil {
		return nil, err
	}
	return habits, nil
}