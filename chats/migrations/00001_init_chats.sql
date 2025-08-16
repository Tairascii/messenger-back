-- +goose Up
-- +goose StatementBegin
create table chats (
    id uuid primary key,
    last_read_message_id bigint
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chats;
-- +goose StatementEnd
