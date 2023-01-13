ALTER table staking rename to staking_contracts;
create table staking_events
(
    id              bigserial not null
                        constraint staking_events_pk
                            primary key,
    created               timestamptz default now() not null,
    updated               timestamptz default now() not null,
    txhash          text not null,
    log_index       numeric not null check ( log_index >= 0 ),
    token           text not null,
    staking_contract text not null,
    staker          text not null,
    stake_value    numeric not null check ( stake_value >= 0 ),
    block_number    numeric not null check ( block_number >= 0 ),
    constraint tokens_fk 
        foreign key(token) 
            references tokens(address),
    constraint staking_contracts_fk
        foreign key(staking_contract)
            references staking_contracts(address),
    constraint staking_events_txhash_log_index_unique
        unique (txhash, log_index)
);

---- create above / drop below ----
-- undo --
ALTER table staking_contracts rename to staking;
drop table staking_events;
