package postgres

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	pb "mymode/generated/habit"
)

func TestCreateHabit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTrackerRepo(db)
	req := &pb.Habit{
		UserId:      "user1",
		Name:        "Exercise",
		Discription: "",
		Frequency:   "",
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO habits (id, user_id, name, discription, frequency, created_at) VALUES ($1, $2, $3, $4, $5, $6)`)).
		WithArgs(sqlmock.AnyArg(), req.UserId, req.Name, req.Discription, req.Frequency, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	habit, err := habitTracker.CreateHabit(req)
	assert.NoError(t, err)
	assert.NotNil(t, habit)
	assert.True(t, habit.Status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetHabit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTrackerRepo(db)
	req := &pb.HabitId{
		HabitId: "habit1",
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "discription", "frequency", "created_at"}).
		AddRow(uuid.New().String(), "user1", "Exercise", "Daily exercise routine", "DAILY", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, user_id, name, discription, frequency, created_at FROM habits WHERE id = $1`)).
		WithArgs(req.HabitId).
		WillReturnRows(rows)

	resp, err := habitTracker.GetHabit(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Exercise", resp.Name)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateHabit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTrackerRepo(db)
	req := &pb.Habit{
		HabitId:     "habit1",
		Name:        "Updated Exercise",
		Discription: "Updated exercise routine",
		Frequency:   "WEEKLY",
	}

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE habits SET name = $1, discription = $2, frequency = $3 WHERE id = $4`)).
		WithArgs(req.Name, req.Discription, req.Frequency, req.HabitId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := habitTracker.UpdateHabit(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteHabit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTrackerRepo(db)
	req := &pb.HabitId{
		HabitId: "habit1",
	}

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM habits WHERE id = $1`)).
		WithArgs(req.HabitId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := habitTracker.DeleteHabit(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserHabits(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTrackerRepo(db)
	req := &pb.UserId{
		UserId: "user1",
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "discription", "frequency", "created_at"}).
		AddRow(uuid.New().String(), req.UserId, "Exercise", "Daily exercise routine", "DAILY", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, user_id, name, discription, frequency, created_at FROM habits WHERE user_id = $1`)).
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := habitTracker.GetUserHabits(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Habbits, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}


func TestCreateHabitLog(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTrackerRepo(db)
	req := &pb.HabitLog{
		HabitId: "habit1",
		Notes:   "Completed workout",
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO habit_logs (id, habit_id, notes) VALUES ($1, $2, $3)`)).
		WithArgs(sqlmock.AnyArg(), req.HabitId, req.Notes).
		WillReturnResult(sqlmock.NewResult(1, 1))

	habitLog, err := habitTracker.CreateHabitLog(req)
	assert.NoError(t, err)
	assert.NotNil(t, habitLog)
	assert.True(t, habitLog.Status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetHabitLogs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTrackerRepo(db)
	req := &pb.HabitId{
		HabitId: "habit1",
	}
	rows := sqlmock.NewRows([]string{"id", "habit_id", "logged_at", "notes"}).
		AddRow(uuid.New().String(), req.HabitId, time.Now().Format(time.RFC3339), "Completed workout")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, habit_id, logged_at, notes FROM habit_logs WHERE habit_id = $1 ORDER BY logged_at DESC`)).
		WithArgs(req.HabitId).
		WillReturnRows(rows)

	resp, err := habitTracker.GetHabitLogs(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Notes, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetHabitSuggestions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTrackerRepo(db)
	req := &pb.Req{}
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "discription", "frequency", "created_at"}).
		AddRow(uuid.New().String(), "user1", "Exercise", "Daily exercise routine", "DAILY", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, user_id, name, discription, frequency, created_at FROM habits ORDER BY created_at DESC LIMIT 10`)).
		WillReturnRows(rows)

	resp, err := habitTracker.GetHabitSuggestions(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Habits, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
