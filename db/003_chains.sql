create table chains
(
     id  bigserial not null
        constraint chains_pk
            primary key,
    created               timestamptz default now() not null,
    updated               timestamptz default now() not null,
    name   text   not null unique,
    rpc_url text not null,
    snapshot_start_block numeric not null check ( snapshot_start_block >= 0 ),
    snapshot_end_block numeric not null check ( snapshot_end_block >= 0 )
);

insert into chains(name, rpc_url, snapshot_start_block, snapshot_end_block) values 
('ETH',  'https://mainnet.infura.io/v3/7e04619b10c04711b2cf8dea7a679ff4', 16078485, 16298368),
('GNO',  'https://gnosischain-rpc.gateway.pokt.network/', 25235581, 25935581),
('POLY', 'https://polygon-mainnet.infura.io/v3/7e04619b10c04711b2cf8dea7a679ff4', 36005795, 38005795);

ALTER TABLE tokens ADD CONSTRAINT chains_fk FOREIGN KEY (chain) REFERENCES chains(name);

---- create above / drop below ----
-- undo --
ALTER TABLE tokens DROP CONSTRAINT chains_fk;
drop table chains;