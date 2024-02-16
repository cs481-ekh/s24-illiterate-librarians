CREATE TABLE IF NOT EXISTS users (
     user_id UUID PRIMARY KEY,
     username VARCHAR(255) NOT NULL,
     password_hash VARCHAR(255) NOT NULL,
     email VARCHAR(255) NOT NULL,
     phone_number VARCHAR(20),
     first_name VARCHAR(50) NOT NULL,
     last_name VARCHAR(50) NOT NULL,
     mailing_address VARCHAR(255),
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS child (
     child_id UUID PRIMARY KEY,
     user_id UUID,
     birth_date DATE,
     grade TINYINT,
     first_name VARCHAR(50) NOT NULL,
     last_name VARCHAR(50) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     FOREIGN KEY (user_id) REFERENCES users(user_id)
);