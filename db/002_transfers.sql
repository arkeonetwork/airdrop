
create table transfers
(
    id              bigserial not null
                        constraint transfers_pk
                            primary key,
    created               timestamptz default now() not null,
    updated               timestamptz default now() not null,
    txhash          text not null,
    log_index       numeric not null check ( log_index >= 0 ),
    token           text,
    transfer_from   text not null,
    transfer_to     text not null,
    transfer_value  bigint,
    block_number    numeric not null check ( block_number >= 0 ),
    constraint tokens_fk 
        foreign key(token) 
            references tokens(address),
    constraint transfers_txhash_log_index_unique
        unique (txhash, log_index)
);

---- create above / drop below ----
-- undo --
drop table transfers;