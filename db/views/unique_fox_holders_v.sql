-- unique list of accounts ever holding fox
create or replace view unique_fox_holders_v as (
    select distinct transfer_to as account
    from transfers
    where token = '0xc770EEfAd204B5180dF6a14Ee197D99d808ee52d'
);
