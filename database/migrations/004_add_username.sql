-- +goose Up

alter table profile add column username text;

-- +goose Down

alter table profile drop column username;