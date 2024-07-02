package postgres

import (
	"database/sql"
	pb "mymode/generated"
)

type HabitTrackerRepo struct{
	Db *sql.DB
}

func NewHabitTrackerRepo(db *sql.DB)*HabitTrackerRepo{
	return &HabitTrackerRepo{Db: db}
}

func(H *HabitTrackerRepo) GetUserHabits(userId *pb.UserId)(*pb.UserHabits, error){
	habits := []*pb.Habit{}
	rows, err := H.Db.Query(`SELECT 
								user_id, habit_id, name, discription, frequency createdAt
							FROM 
								Habits
							WHERE 
								user_id = $1`, userId.UserId)
	if err != nil{
		return &pb.UserHabits{Habbits: habits}, err
	}
	for rows.Next(){
		habit := &pb.Habit{}
		err := rows.Scan(&habit.HabbitId, &habit.UserId, &habit.Name, &habit.Discription, &habit.Frequency, habit.CreatedAt)
		if err != nil{
			return &pb.UserHabits{Habbits: habits}, err
		}
		habits = append(habits, habit)
	}
	return &pb.UserHabits{Habbits: habits}, nil
}

func(H *HabitTrackerRepo) CreateHabitLog(habbitLog *pb.HabitLog)(*pb.Status, error){
	_, err := H.Db.Exec(`INSERT INTO habit_logs(notes) VALUES($1)`, habbitLog.Notes)
	if err != nil{
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func(H *HabitTrackerRepo) GetHabitLogs(habitId *pb.HabitId)(*pb.HabitLog, error){
	habitLog := &pb.HabitLog{}
	err := H.Db.QueryRow(`SELECT habbit_id, logged_at, notes FROM habit_logs`).Scan(
							&habitLog.HabbitId, &habitLog.LoggedAt, &habitLog.Notes)
	return habitLog, err
}