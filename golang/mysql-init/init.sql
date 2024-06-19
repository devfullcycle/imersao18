CREATE DATABASE IF NOT EXISTS test_db;

USE test_db;

CREATE TABLE events (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  location VARCHAR(255) NOT NULL,
  organization VARCHAR(255) NOT NULL,
  rating VARCHAR(10) NOT NULL,
  date DATETIME NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  capacity INT NOT NULL,
  price FLOAT NOT NULL,
  partner_id INT NOT NULL
);

CREATE TABLE spots (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  event_id VARCHAR(36) NOT NULL,
  name VARCHAR(10) NOT NULL,
  status VARCHAR(10) NOT NULL,
  ticket_id VARCHAR(36),
  FOREIGN KEY (event_id) REFERENCES events(id)
);

CREATE TABLE tickets (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  event_id VARCHAR(36) NOT NULL,
  spot_id VARCHAR(36) NOT NULL,
  ticket_type VARCHAR(10) NOT NULL,
  price FLOAT NOT NULL,
  FOREIGN KEY (event_id) REFERENCES events(id),
  FOREIGN KEY (spot_id) REFERENCES spots(id)
);
