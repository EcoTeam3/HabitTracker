package postgres

import (
	"database/sql"
	"fmt"
	pb "mymode/generated"
	"strconv"
	"strings"
	"time"
)

type HabitTrackerRepo struct {
	DB *sql.DB
}

func NewHabitTrackerRepo(db *sql.DB) *HabitTrackerRepo {
	return &HabitTrackerRepo{DB: db}
}

func (H *HabitTrackerRepo) CreateHabit(habit *pb.Habit) (*pb.Status, error) {
	_, err := H.DB.Exec("INSERT INTO Habits(user_id,name,discription,frequency) VALUES($1,$2,$3,$4)", habit.UserId, habit.Name, habit.Discription, habit.Frequency)

	if err != nil {
		return &pb.Status{
			Status: false,
		}, err
	}

	return &pb.Status{
		Status: true,
	}, nil
}

func (H *HabitTrackerRepo) GetHabit(habbitId *pb.HabitId) (*pb.Habit, error) {
	habbit := pb.Habit{}
	err := H.DB.QueryRow("SELECT * FROM Habits WHERE habit_id = $1", habbitId.HabitId).Scan(
		&habbit.HabitId, &habbit.UserId, &habbit.Name, &habbit.Discription, &habbit.Frequency)
		if err != nil {
			return nil, err
	}
	return &habbit, nil
}

func (H *HabitTrackerRepo) UpdateHabit(habit *pb.Habit) (*pb.Status, error) {
	time := time.Now()
	query := "UPDATE Habits SET update_at = $1 "
	arr := []interface{}{time}
	var param []string

	if len(habit.UserId) > 0 {
		arr = append(arr, habit.UserId)
		param = append(param, "user_id")
		query += ", user_id = :user_id "
	}
	
	if len(habit.Name) > 0 {
		arr = append(arr, habit.Name)
		param = append(param, "name")
		query += ", name = :name"
	}
	
	if len(habit.Discription) > 0 {
		arr = append(arr, habit.Discription)
		param = append(param, "discription")
		query += ", discription = :discription"
	}
	
	n := 2
	for _, j := range param {
		query = strings.Replace(query, ":"+j, "$"+strconv.Itoa(n), 1)
		n++
	}
	query += fmt.Sprintf(" WHERE deleted_at is null and habit_id = $%d", n)
	arr = append(arr, habit.HabitId)
	_, err := H.DB.Exec(query, arr...)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}
func (H *HabitTrackerRepo)DeleteHabit(habitId *pb.HabitId)(*pb.Status,error){
	_,err := H.DB.Exec("DELETE FROM Habits WHERE habit_id = $1",habitId.HabitId)
	if err != nil{
		return &pb.Status{
			Status: false,
		},err
	}
	return &pb.Status{
		Status: true,
	},err
}


func(H *HabitTrackerRepo) GetUserHabits(userId *pb.UserId)(*pb.UserHabits, error){
	habits := []*pb.Habit{}
	rows, err := H.DB.Query(`SELECT 
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
		err := rows.Scan(&habit.HabitId, &habit.UserId, &habit.Name, &habit.Discription, &habit.Frequency, habit.CreatedAt)
		if err != nil{
			return &pb.UserHabits{Habbits: habits}, err
		}
		habits = append(habits, habit)
	}
	return &pb.UserHabits{Habbits: habits}, nil
}

func(H *HabitTrackerRepo) CreateHabitLog(habbitLog *pb.HabitLog)(*pb.Status, error){
	_, err := H.DB.Exec(`INSERT INTO habit_logs(notes) VALUES($1)`, habbitLog.Notes)
	if err != nil{
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}



func(H *HabitTrackerRepo)GetHabitLogs(habitId *pb.HabitId)(*pb.HabitLog, error){
	habitLog := &pb.HabitLog{}
	err := H.DB.QueryRow(`SELECT habbit_id, logged_at, notes FROM habit_logs`).Scan(
							&habitLog.HabitId, &habitLog.LoggedAt, &habitLog.Notes)
	return habitLog, err
}

func (H *HabitTrackerRepo)GetHabitSuggestions(request *pb.Req)(*pb.Habits,error){
	rows,err := H.DB.Query("SELECT * FROM Habit_logs")
	if err != nil{
		return nil,err
	}

	habits := pb.Habits{}
	for rows.Next(){
		habitLog := pb.HabitLog{}
		err = rows.Scan(&habitLog.Id,&habitLog.HabitId,&habitLog.LoggedAt,&habitLog.Notes)
		if err != nil{
			return nil,err
		}
		habits.Habits = append(habits.Habits, &habitLog)
	}
	return &habits,nil
}
