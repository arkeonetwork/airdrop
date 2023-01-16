-- just prevents adding account criteria while building out queries
create or replace view my_transfers_v as (
    select * from fox_transfers_v
    where account = lower('0xFea9c0446F53E9E91b108Ba3C33cE21B844E4f98')
);
