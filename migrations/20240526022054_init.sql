-- +goose Up
create table if not exists client
(
    id         bigint generated always as identity primary key,
    first_name text      default '',
    last_name  text      default '',
    created_at timestamp default (now() at time zone 'utc')
);

create table if not exists reinvestment_period
(
    id                bigint    not null unique,
    asset_address     text      not null unique,
    reinvestment_rate bigint    not null,
    asset_price       bigint    not null,
    start_time        timestamp not null,
    end_time          timestamp default '0001-01-01 00:00:00.000000',
    created_at timestamp default (now() at time zone 'utc')
);

create table if not exists user_reinvestment
(
    client_id                bigint  not null,
    reinvestment_period_id bigint  not null,
    savings_balance        numeric not null,
    asset_amount           numeric not null,
    created_at             timestamp default (now() at time zone 'utc')
);

create unique index if not exists
    user_reinvestment_user_id_reinvestment_period_id on
    user_reinvestment (client_id, reinvestment_period_id);


-- +goose Down
