-- Create and select the DB
CREATE DATABASE IF NOT EXISTS literacy_link_db;

USE literacy_link_db;

-- Create the users and associated organization

CREATE TABLE IF NOT EXISTS Users (
     user_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     username VARCHAR(255) NOT NULL,
     password_hash VARCHAR(255) NOT NULL,
     email VARCHAR(255) NOT NULL,
     phone_number VARCHAR(20),
     first_name VARCHAR(50) NOT NULL,
     last_name VARCHAR(50) NOT NULL,
     mailing_address VARCHAR(255),
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Parents (
     parent_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     user_id BINARY(16) NOT NULL,
     FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS Tutors (
     tutor_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     user_id BINARY(16) NOT NULL,
     FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS Instructors (
     instructor_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     user_id BINARY(16) NOT NULL,
     FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS Admins (
     admin_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     user_id BINARY(16) NOT NULL,
     FOREIGN KEY (user_id) REFERENCES Users(user_id)
);


CREATE TABLE IF NOT EXISTS Child (
     child_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     user_id BINARY(16) NOT NULL,
     parent_id BINARY(16) NOT NULL,
     birth_date DATE,
     grade TINYINT,
     first_name VARCHAR(50) NOT NULL,
     last_name VARCHAR(50) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     FOREIGN KEY (parent_id) REFERENCES Parents(parent_id)
);


-- Create Semester structures

CREATE TABLE IF NOT EXISTS Semesters (
     semester_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     spring_or_fall VARCHAR(20) NOT NULL, -- "spring" for spring, "fall" for fall
     first_school_day DATE NOT NULL,
     last_school_day DATE NOT NULL,
     first_tutor_day DATE NOT NULL,
     last_tutor_day DATE NOT NULL,
     monday_of_break DATE NOT NULL
);


CREATE TABLE IF NOT EXISTS EOS_parent_survey (
     EOS_p_s_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     child_id BINARY(16) NOT NULL,
     parent_id BINARY(16) NOT NULL,
     semester_id BINARY(16) NOT NULL,
     survey_complete_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     -- add in all of the responses that need to be recorded in the survey here
     FOREIGN KEY (child_id) REFERENCES Child(child_id),
     FOREIGN KEY (parent_id) REFERENCES Parents(parent_id),
     FOREIGN KEY (semester_id) REFERENCES Semesters(semester_id)
);


CREATE TABLE IF NOT EXISTS Semester_tutoring_obj (
     semester_tutoring_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     child_id BINARY(16) NOT NULL,
     parent_id BINARY(16) NOT NULL,
     EOS_parent_survey_complete boolean DEFAULT false,
     survey_complete_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     EOS_parent_survey_id BINARY(16),
     EOS_report_posted boolean DEFAULT false,
     EOS_report_file_path VARCHAR(127),
     semester_id BINARY(16) NOT NULL,
     FOREIGN KEY (child_id) REFERENCES Child(child_id),
     FOREIGN KEY (parent_id) REFERENCES Parents(parent_id)
);


CREATE TABLE IF NOT EXISTS App_for_tutoring (
     app_for_tut_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     child_id BINARY(16) NOT NULL,
     parent_id BINARY(16) NOT NULL,
     app_complete boolean DEFAULT false,
     app_complete_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     desired_semester_id BINARY(16) NOT NULL,
     -- add in all of the responses that need to be recorded for the app here
     FOREIGN KEY (child_id) REFERENCES Child(child_id),
     FOREIGN KEY (parent_id) REFERENCES Parents(parent_id),
     FOREIGN KEY (desired_semester_id) REFERENCES Semesters(semester_id)
);


-- Create tutoring structures

CREATE TABLE IF NOT EXISTS Tutor_session (
     tutor_session_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     child_id BINARY(16) NOT NULL,
     parent_id BINARY(16) NOT NULL,
     zoom_join_link VARCHAR(512),
     zoom_recording_link VARCHAR(512),
     meeting_date DATETIME NOT NULL,
     tutor_id BINARY(16) NOT NULL,
     semester_id BINARY(16) NOT NULL,
     FOREIGN KEY (child_id) REFERENCES Child(child_id),
     FOREIGN KEY (parent_id) REFERENCES Parents(parent_id),
     FOREIGN KEY (tutor_id) REFERENCES Tutors(tutor_id),
     FOREIGN KEY (semester_id) REFERENCES Semesters(semester_id)
);


CREATE TABLE IF NOT EXISTS Session_notes (
     session_notes_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     tutor_session_id BINARY(16) NOT NULL,
     created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     FOREIGN KEY (tutor_session_id) REFERENCES Tutor_session(tutor_session_id)
);


