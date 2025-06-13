-- +goose Up
-- +goose StatementBegin
INSERT INTO agama (name) VALUES
('Agama Islam'),
('Agama Kristen'),
('Agama Katolik'),
('Agama Hindu'),
('Agama Buddha'),
('Agama Konghucu'),
('Agama Lainnya');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM agama WHERE name IN (
    'Agama Islam',
    'Agama Kristen',
    'Agama Katolik',
    'Agama Hindu',
    'Agama Buddha',
    'Agama Konghucu',
    'Agama Lainnya'
);
-- +goose StatementEnd
