
create table if not exists person
(
    id              uuid not null
        constraint pk_person primary key,
    name            varchar,
    normalized_name varchar,
    token           varchar
);

create unique index if not exists uk_person_normalized_name
    on person (normalized_name);

create unique index if not exists uk_person_token
    on person (token);

create table if not exists item
(
    id   uuid    not null
        constraint pk_item primary key,
    name varchar not null
);
create unique index if not exists uk_item_name
    on item (name);

create table if not exists transaction
(
    id        uuid        not null
        constraint pk_transaction primary key,
    created   timestamptz not null,
    person_id uuid        not null
        constraint fk_transaction_person references person (id),
    item_id   uuid        not null
        constraint fk_transaction_item references item (id)
);
create index if not exists idx_transaction_person
    on transaction (person_id);
