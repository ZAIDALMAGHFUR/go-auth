-- +goose Up
-- +goose StatementBegin
INSERT INTO agama (name) VALUES
('Islam'),
('Kristen'),
('Katolik'),
('Hindu'),
('Buddha'),
('Konghucu'),
('Lainnya');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM agama WHERE name IN (
    'Islam',
    'Kristen',
    'Katolik',
    'Hindu',
    'Buddha',
    'Konghucu',
    'Lainnya'
);
-- +goose StatementEnd