-- Masukkan akun default untuk testing
INSERT INTO user_credentials (username, email, password, created_at, updated_at) 
VALUES ('admin', 'admin@example.com', '$2a$10$YourHashedPasswordHere...', NOW(), NOW());

INSERT INTO user_profiles (credential_id, full_name, phone_number) 
VALUES (1, 'Super Administrator', '08123456789');