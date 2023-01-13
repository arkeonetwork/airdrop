create table chains
(
     id  bigserial not null
        constraint chains_pk
            primary key,
    created               timestamptz default now() not null,
    updated               timestamptz default now() not null,
    name   text   not null unique,
    rpc_url text not null
);

insert into chains(name, rpc_url) values ('ETH', 'https://mainnet.infura.io/v3/7e04619b10c04711b2cf8dea7a679ff4');
insert into chains(name, rpc_url) values ('GNO', 'https://gnosischain-rpc.gateway.pokt.network/');
insert into chains(name, rpc_url) values ('POLY', 'https://polygon-mainnet.infura.io/v3/7e04619b10c04711b2cf8dea7a679ff4');

ALTER TABLE tokens ADD CONSTRAINT chains_fk FOREIGN KEY (chain) REFERENCES chains(name);

---- create above / drop below ----
-- undo --
ALTER TABLE tokens DROP CONSTRAINT chains_fk;
drop table chains;