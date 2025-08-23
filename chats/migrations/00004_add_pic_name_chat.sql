-- +goose Up
-- +goose StatementBegin
alter table chats
add column picture text,
add column title varchar(255) not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table chats
drop column picture,
drop column title;
-- +goose StatementEnd
