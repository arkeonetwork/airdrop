create or replace view fox_transfers_v as (
    select id,txhash,transfer_to as account,transfer_value as delta, block_number
    from unique_fox_holders_v holders join transfers on holders.account = transfers.transfer_to
    where token = '0xc770eefad204b5180df6a14ee197d99d808ee52d'
      and transfer_from != transfer_to
    union
    select id, txhash,transfer_from,-(transfer_value),block_number
    from unique_fox_holders_v holders join transfers on holders.account = transfers.transfer_from
    where token = '0xc770eefad204b5180df6a14ee197d99d808ee52d'
    and transfer_from != transfer_to
    order by block_number
);
