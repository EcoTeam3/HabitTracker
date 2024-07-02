-- insert_habits_up.sql

-- Insert initial data into habits table
INSERT INTO habits (user_id, name, description, frequency)
VALUES
(gen_random_uuid(), 'Morning Run', 'Running every morning for 30 minutes.', 'daily'),
(gen_random_uuid(), 'Read a Book', 'Read at least one chapter of a book.', 'weekly'),
(gen_random_uuid(), 'Monthly Reflection', 'Reflect on the past month and plan for the next.', 'monthly');

-- Insert initial data into habit_logs table
INSERT INTO habit_logs (habit_id, logged_at, notes)
VALUES
((SELECT habit_id FROM habits WHERE name='Morning Run'), '2024-07-01 06:00:00+00', 'Great run today!'),
((SELECT habit_id FROM habits WHERE name='Read a Book'), '2024-07-01 19:00:00+00', 'Finished the first chapter.'),
((SELECT habit_id FROM habits WHERE name='Monthly Reflection'), '2024-07-01 21:00:00+00', 'Productive month!');
