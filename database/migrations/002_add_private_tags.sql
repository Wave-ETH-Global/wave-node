-- +goose Up

alter table connection add private_tags text[];
alter table connection drop column tags_b;
alter table connection add connected_at time;
alter table connection rename column tags_a to tags;

-- +goose Down

alter table connection drop column private_tags;
alter table connection add column tags_b text[];
alter table connection rename column tags to tags_a;