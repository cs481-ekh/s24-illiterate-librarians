-- Create and select the DB
CREATE DATABASE IF NOT EXISTS literacy_link_db;

USE literacy_link_db;

-- Create the users and associated organization

CREATE TABLE IF NOT EXISTS Users (
     user_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     username VARCHAR(255) NOT NULL UNIQUE,
     password_hash VARCHAR(255) NOT NULL,
     email VARCHAR(255) NOT NULL UNIQUE,
     phone_number VARCHAR(20) NULL UNIQUE,
     first_name VARCHAR(50) NOT NULL,
     last_name VARCHAR(50) NOT NULL,
     mailing_address VARCHAR(255),
     pref_method_comm VARCHAR(10), -- Has to be "C" for call, "T" for text, or "E" for email.
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
     parent_id BINARY(16) NOT NULL,
     birth_date DATE,
     grade TINYINT NOT NULL,
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
     tuesday_of_break DATE NOT NULL,
     thursday_of_break DATE NOT NULL
);


CREATE TABLE IF NOT EXISTS EOS_parent_survey (
     EOS_p_s_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     child_id BINARY(16) NOT NULL,
     parent_id BINARY(16) NOT NULL,
     semester_id BINARY(16) NOT NULL,
     survey_complete_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     -- add in all of the responses that need to be recorded in the survey here
     child_barrier_reading VARCHAR(511) NOT NULL,
     child_barrier_writing VARCHAR(511) NOT NULL,
     want_parent_training TINYINT NOT NULL, -- 1 (very likely), 2 (Likely), 3 (Unsure), 4 (unlikely), 5 (very unlikely)

     -- The next group of booleans should together make a single select all type of question, but will be stored using multiple variables. 
     online_modules boolean DEFAULT false,
     zoom_meetings boolean DEFAULT false,
     in_person boolean DEFAULT false,
     blended boolean DEFAULT false,
     individual_coaching boolean DEFAULT false,

     -- Below group of questions are satisfied/disatisfied ratings: 1 (very dissatisfied) up to 10 (very satisfied)
     family_tutor_relationship TINYINT NOT NULL, 
     family_tutor_communication TINYINT NOT NULL,
     child_instruction_recieved TINYINT NOT NULL,
     child_enjoyment TINYINT NOT NULL,
     child_confidence_r TINYINT NOT NULL, -- confidence in reading
     child_confidence_w TINYINT NOT NULL, -- writing
     child_confidence_s TINYINT NOT NULL, -- spelling

     prefer_zoom boolean NOT NULL, -- true if prefer online(zoom), false if prefer in-person (college of education building)
     child_enjoy_most VARCHAR(511) NOT NULL,
     improvments_recommendation VARCHAR(511) NOT NULL,
     misc_feedback VARCHAR(511) NOT NULL,


     FOREIGN KEY (child_id) REFERENCES Child(child_id),
     FOREIGN KEY (parent_id) REFERENCES Parents(parent_id),
     FOREIGN KEY (semester_id) REFERENCES Semesters(semester_id)

);


CREATE TABLE IF NOT EXISTS Semester_tutoring_obj (
     semester_tutoring_id BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY,
     child_id BINARY(16) NOT NULL,
     parent_id BINARY(16) NOT NULL,
     application_approved boolean DEFAULT false,
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
     child_data_consent boolean DEFAULT false,
     photo_release_consent boolean DEFAULT false,
     need_financial_assistance boolean DEFAULT false,
     -- NOTE: guardian 2 is not required to be filled out
     guardian2_first_n VARCHAR(50),
     guardian2_last_n VARCHAR(50),
     guardian2_phone VARCHAR(20),
     guardian2_email VARCHAR(255),

     emergency_con_name VARCHAR(50) NOT NULL,
     emergency_con_relation VARCHAR(255) NOT NULL,
     emergency_con_phone VARCHAR(20) NOT NULL,

     previous_child_participation boolean DEFAULT false,
     what_semester VARCHAR(50),
     child_current_school VARCHAR(50) NOT NULL,
     list_languages_spoken VARCHAR(255) NOT NULL,
     received_special_ed VARCHAR(511) NOT NULL,
     list_challenges VARCHAR(511) NOT NULL,
     how_long_concerned VARCHAR(255) NOT NULL,
     describe_hopes VARCHAR(511) NOT NULL,
     child_allergy_meds VARCHAR(255) NOT NULL,
     misc_info VARCHAR(511),
     hear_about_litLab VARCHAR(255),

     --
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
     parent_avail boolean DEFAULT true,
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


