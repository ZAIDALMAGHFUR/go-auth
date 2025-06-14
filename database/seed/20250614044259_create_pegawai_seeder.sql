-- +goose Up
-- +goose StatementBegin
INSERT INTO pegawai (nama, jabatan, email, alamat) VALUES
('Andrian Didan', 'Software Engineer', 'didan@example.com', 'Jakarta'),
('Budi Santoso', 'Manager', 'budi@example.com', 'Bandung'),
('Siti Aminah', 'HR', 'siti@example.com', 'Surabaya');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM pegawai WHERE email IN ('didan@example.com', 'budi@example.com', 'siti@example.com');
-- +goose StatementEnd
