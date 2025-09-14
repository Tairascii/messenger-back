-- +goose Up
-- +goose StatementBegin
alter table chats
add column last_message_id bigint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table chats
drop column last_message_id;
-- +goose StatementEnd
