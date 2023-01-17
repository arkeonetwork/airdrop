-- just prevents adding account criteria while building out queries
create or replace view my_transfers_v as (
    select * from fox_transfers_v
    where account = '0xfea9c0446f53e9e91b108ba3c33ce21b844e4f98'
);
