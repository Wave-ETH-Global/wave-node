-- +goose Up
alter table connection add status text;
alter table connection add read boolean;

-- +goose Down
alter table connection drop status;
alter table connection drop read;
