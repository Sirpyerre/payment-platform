-- we don't know how to generate root <with-no-name> (class Root) :(

comment on database postgres is 'default administrative connection database';

create table merchants
(
    id         serial
        primary key,
    name       varchar(255) not null,
    created_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table merchants
    owner to postgres;

create table customers
(
    id         serial
        primary key,
    name       varchar(255) not null,
    email      varchar(255) not null
        unique,
    created_at timestamp with time zone default CURRENT_TIMESTAMP
);

alter table customers
    owner to postgres;

create table transactions
(
    id                  serial
        primary key,
    merchant_id         integer
        references merchants,
    customer_id         integer
        references customers,
    amount              numeric(10, 2)                     not null,
    status              varchar(50)                        not null,
    created_at          timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at          timestamp with time zone default CURRENT_TIMESTAMP,
    transaction_bank_id bigint                   default 0 not null
);

alter table transactions
    owner to postgres;

create table refunds
(
    id             serial
        primary key,
    transaction_id integer
        references transactions,
    amount         numeric(10, 2) not null,
    created_at     timestamp with time zone default CURRENT_TIMESTAMP,
    status         varchar(50)    not null
);

alter table refunds
    owner to postgres;

