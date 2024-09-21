-- Create the database
CREATE DATABASE IF NOT EXISTS stepdb;

-- Use the newly created database
USE stepdb;

-- Create a table called 'tasks'
CREATE TABLE IF NOT EXISTS tasks (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
								  status VARCHAR(50) NOT NULL,
								  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert some initial data into the 'tasks' table
INSERT INTO tasks (name, status) VALUES
('Task 1', 'scheduled'),
('Task 2', 'completed'),
('Task 3', 'pending');
