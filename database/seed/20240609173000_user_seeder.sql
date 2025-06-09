-- +goose Up
INSERT INTO users (name, email, password) VALUES
('Admin', 'admin@example.com', '$2a$12$Q4GxQFQ/Ofcj8bqLfgS/s.dA/DFElOIWWAg1hb1vOG9JK1rzSEklW');

-- +goose Down
DELETE FROM users WHERE email = 'admin@example.com';
