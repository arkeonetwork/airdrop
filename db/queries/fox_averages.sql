with params as (
    select name, snapshot_start_block, snapshot_end_block
    from chains
    where name = 'ETH'
),
averageable as (
select account,
       cumulative_balance,
       block_number,
       blocks_between,
       blocks_between * lag(cumulative_balance, 1, 0) over (partition by account) as avg_over_blocks
from (select account,
             SUM(delta) over (partition by account order by block_number)                         as cumulative_balance,
             ts.block_number,
             block_number -
             lag(block_number, 1, block_number) over (partition by account order by block_number) as blocks_between
      from (select account,
                   delta,
                   block_number
            from fox_transfers_v
            where block_number >= (select snapshot_start_block from params)
              and block_number <= (select snapshot_end_block from params)
            union -- starting balance
            (select account,
                    sum(delta),
                    (select snapshot_start_block from params) as block_number
             from fox_transfers_v
             where block_number <= (select snapshot_start_block from params)
             group by account
             order by block_number)
            union -- ending balance
            (select account,
                    sum(delta),
                    (select snapshot_end_block from params) as block_number
             from fox_transfers_v
             where block_number <= (select snapshot_end_block from params)
             group by account
             order by block_number)) as ts
      where ts.block_number >= (select snapshot_start_block from params)
        and ts.block_number <= (select snapshot_end_block from params)) as x
order by block_number
)
select account, sum(avg_over_blocks) / ((select snapshot_end_block from params) - (select snapshot_start_block from params)) avg_hold
from averageable
group by account
-- if we want a min/cut-off amount
-- having sum(avg_over_blocks) / ((select snapshot_end_block from params) - (select snapshot_start_block from params)) > 0.01
order by avg_hold desc
;
