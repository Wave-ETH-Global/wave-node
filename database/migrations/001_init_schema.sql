-- +goose Up

create table profile (
   uuid text primary key,
   eth_address text unique,
   metadata json,
   public_tags text[],
   tokens json
);

create table connection (
    id bigserial,
    vertex_a text references profile(uuid),
    vertex_b text references profile(uuid),
    tags_a text[],
    tags_b text[]
);


-- +goose Down

drop table connection;
drop table profile;