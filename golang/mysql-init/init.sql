DROP DATABASE IF EXISTS test_db;

CREATE DATABASE test_db;

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
  ticket_kind VARCHAR(10) NOT NULL,
  price FLOAT NOT NULL,
  FOREIGN KEY (event_id) REFERENCES events(id),
  FOREIGN KEY (spot_id) REFERENCES spots(id)
);

INSERT INTO events (id, name, location, organization, rating, date, image_url, capacity, price, partner_id) VALUES
  ('10853e59-dc5b-4d7b-a028-01513ef50d76', 'Event 001 - Partner1', 'São Paulo, SP', 'Partner 1', 'L14', '2021-10-10 10:00:00', 'https://images.unsplash.com/photo-1470229722913-7c0e2dbbafd3', 10, 100, 1),
  ('e0352b32-7698-4805-b029-28302b3a911f', 'Event 002 - Partner1', 'Rio de Janeiro, RJ', 'Partner 1', 'L14', '2021-10-10 12:00:00', 'https://images.unsplash.com/photo-1459749411175-04bf5292ceea', 10, 200, 1),
  ('5b79831a-a9d3-4538-8fb5-569494bd17a5', 'Event 003 - Partner2', 'Belo Horizonte, MG', 'Partner 2', 'L12', '2024-10-10 10:00:00', 'https://images.unsplash.com/photo-1540039155733-5bb30b53aa14', 10, 400, 2),
  ('8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'Event 004 - Partner2', 'Uberlândia, MG', 'Partner 2', 'L16', '2024-10-10 12:00:00', 'https://images.unsplash.com/photo-1493225457124-a3eb161ffa5f', 10, 500, 2)
;

INSERT INTO spots (id, event_id, name, status, ticket_id) VALUES
  ('f1b1b1b1-1b1b-1b1b-1b1b-1b1b1b1b1b1b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'A1', 'available', ""),
  ('f2b2b2b2-2b2b-2b2b-2b2b-2b2b2b2b2b2b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'A2', 'available', ""),
  ('f3b3b3b3-3b3b-3b3b-3b3b-3b3b3b3b3b3b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'A3', 'available', ""),
  ('f4b4b4b4-4b4b-4b4b-4b4b-4b4b4b4b4b4b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'A4', 'available', ""),
  ('f5b5b5b5-5b5b-5b5b-5b5b-5b5b5b5b5b5b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'A5', 'available', ""),
  ('7c022408-c6ec-4575-b362-923822ee83b4', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'B1', 'available', ""),
  ('f6b6b6b6-6b6b-6b6b-6b6b-6b6b6b6b6b6b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'B2', 'available', ""),
  ('f7b7b7b7-7b7b-7b7b-7b7b-7b7b7b7b7b7b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'B3', 'available', ""),
  ('f8b8b8b8-8b8b-8b8b-8b8b-8b8b8b8b8b8b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'B4', 'available', ""),
  ('f9b9b9b9-9b9b-9b9b-9b9b-9b9b9b9b9b9b', '10853e59-dc5b-4d7b-a028-01513ef50d76', 'B5', 'available', ""),
  ('deb28dbe-cbe1-4bf3-a7e3-bce4aa52b54f', 'e0352b32-7698-4805-b029-28302b3a911f', 'A1', 'available', ""),
  ('e0e0e0e0-0e0e-0e0e-0e0e-0e0e0e0e0e0e', 'e0352b32-7698-4805-b029-28302b3a911f', 'A2', 'available', ""),
  ('e1e1e1e1-1e1e-1e1e-1e1e-1e1e1e1e1e1e', 'e0352b32-7698-4805-b029-28302b3a911f', 'A3', 'available', ""),
  ('e2e2e2e2-2e2e-2e2e-2e2e-2e2e2e2e2e2e', 'e0352b32-7698-4805-b029-28302b3a911f', 'A4', 'available', ""),
  ('e3e3e3e3-3e3e-3e3e-3e3e-3e3e3e3e3e3e', 'e0352b32-7698-4805-b029-28302b3a911f', 'A5', 'available', ""),
  ('6c7bdf8d-9146-43df-8b0b-3ae3d4c18cba', 'e0352b32-7698-4805-b029-28302b3a911f', 'B1', 'available', ""),
  ('e4e4e4e4-4e4e-4e4e-4e4e-4e4e4e4e4e4e', 'e0352b32-7698-4805-b029-28302b3a911f', 'B2', 'available', ""),
  ('e5e5e5e5-5e5e-5e5e-5e5e-5e5e5e5e5e5e', 'e0352b32-7698-4805-b029-28302b3a911f', 'B3', 'available', ""),
  ('e6e6e6e6-6e6e-6e6e-6e6e-6e6e6e6e6e6e', 'e0352b32-7698-4805-b029-28302b3a911f', 'B4', 'available', ""), 
  ('e7e7e7e7-7e7e-7e7e-7e7e-7e7e7e7e7e7e', 'e0352b32-7698-4805-b029-28302b3a911f', 'B5', 'available', ""),
  ('e8e8e8e8-8e8e-8e8e-8e8e-8e8e8e8e8e8e', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'A1', 'available', ""),
  ('e9e9e9e9-9e9e-9e9e-9e9e-9e9e9e9e9e9e', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'A2', 'available', ""),
  ('fafafafa-afaf-afaf-afaf-afafafafafaf', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'A3', 'available', ""),
  ('fbfbfbfb-bfbf-bfbf-bfbf-bfbfbfbfbfbf', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'A4', 'available', ""),
  ('fcfcfcfc-cfcf-cfcf-cfcf-cfcfcfcfcfcf', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'A5', 'available', ""),
  ('f8f8f8f8-8f8f-8f8f-8f8f-8f8f8f8f8f8f', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'B1', 'available', ""),
  ('fdfdfdfd-dfdf-dfdf-dfdf-dfdfdfdfdfdf', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'B2', 'available', ""),
  ('fefefefe-efef-efef-efef-efefefefefef', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'B3', 'available', ""),
  ('f0f0f0f0-0f0f-0f0f-0f0f-0f0f0f0f0f0f', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'B4', 'available', ""),
  ('f1f1f1f1-1f1f-1f1f-1f1f-1f1f1f1f1f1f', '5b79831a-a9d3-4538-8fb5-569494bd17a5', 'B5', 'available', ""),
  ('f2f2f2f2-2f2f-2f2f-2f2f-2f2f2f2f2f2f', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'A1', 'available', ""),
  ('f3f3f3f3-3f3f-3f3f-3f3f-3f3f3f3f3f3f', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'A2', 'available', ""),
  ('f4f4f4f4-4f4f-4f4f-4f4f-4f4f4f4f4f4f', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'A3', 'available', ""),
  ('f5f5f5f5-5f5f-5f5f-5f5f-5f5f5f5f5f5f', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'A4', 'available', ""),
  ('f6f6f6f6-6f6f-6f6f-6f6f-6f6f6f6f6f6f', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'A5', 'available', ""),
  ('af20c380-b6c8-4c99-b1d9-780871b80ab1', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'B1', 'available', ""),
  ('f7f7f7f7-7f7f-7f7f-7f7f-7f7f7f7f7f7f', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'B2', 'available', ""),
  ('cb3e9985-dec6-4c7b-9675-409dad659196', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'B3', 'available', ""),
  ('f9f9f9f9-9f9f-9f9f-9f9f-9f9f9f9f9f9f', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'B4', 'available', ""),
  ('g0g0g0g0-0g0g-0g0g-0g0g-0g0g0g0g0g0g', '8beff8fd-39e4-49ea-ae5e-a0ec9af888c5', 'B5', 'available', "")
;
  
  
  
  
  
