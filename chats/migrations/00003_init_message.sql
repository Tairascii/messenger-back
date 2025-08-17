-- +goose Up
-- +goose StatementBegin
create table messages (
    id bigserial primary key,
    text text,
    is_edited bool,
    created_at timestamp default now(),
    sender_id uuid,
    chat_id uuid
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table messages;
-- +goose StatementEnd
