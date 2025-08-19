-- +goose Up
-- +goose StatementBegin
create table users (
    id uuid primary key not null default gen_random_uuid(),
    email varchar(255) not null unique,
    password text not null,
    created_at timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
