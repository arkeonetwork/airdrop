
create table transfers
(
    id              bigserial not null
                        constraint transfers_pk
                            primary key,
    txhash          text not null,
    token           text,
    transfer_from   text not null,
    transfer_to     text not null,
    transfer_value  bigint,
    block_number    numeric not null check ( block_number >= 0 ),
    constraint tokens_fk 
        foreign key(token) 
            references tokens(address)
);

---- create above / drop below ----
-- undo --
drop table transfers;