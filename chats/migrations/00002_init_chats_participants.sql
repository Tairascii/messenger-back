-- +goose Up
-- +goose StatementBegin
create table chats_participants (
    chat_id uuid,
    user_id uuid,
    primary key (chat_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table chats_participants;
-- +goose StatementEnd
