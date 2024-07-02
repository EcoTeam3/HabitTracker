-- insert_habits_down.sql

-- Delete initial data from habit_logs table
DELETE FROM habit_logs WHERE notes IN 
('Great run today!', 'Finished the first chapter.', 'Productive month!');

-- Delete initial data from habits table
DELETE FROM habits WHERE name IN 
('Morning Run', 'Read a Book', 'Monthly Reflection');
