CREATE DATABASE Amitis;
USE Amitis;


CREATE TABLE images (
	id INT auto_increment PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
	uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	data LONGBLOB
    );

#VALUES ('cat.jpg', 'image/jpeg', 204800, LOAD_FILE('/path/to/cat.jpg'));

INSERT INTO images (filename, mime_type, data)
VALUES ('1.jpg', 'image/jpg', LOAD_FILE('.\v0.1-amitis\1.jpg'));




    
