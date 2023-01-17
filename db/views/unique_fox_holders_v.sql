-- unique list of accounts ever holding fox
create or replace view unique_fox_holders_v as (
    select distinct transfer_to as account
    from transfers
    where token = '0xc770eefad204b5180df6a14ee197d99d808ee52d'
);
