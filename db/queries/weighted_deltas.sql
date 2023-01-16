select cumulative_balance,
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
            from my_transfers_v
            where block_number >= 16078485
              and block_number <= 16298368
            union -- starting balance
            (select account,
                    sum(delta),
                    16078485 as block_number
             from my_transfers_v
             where block_number <= 16078485
             group by account
             order by block_number)
            union -- ending balance
            (select account,
                    sum(delta),
                    16298368 as block_number
             from my_transfers_v
             where block_number <= 16298368
             group by account
             order by block_number)) as ts
      where ts.block_number >= 16078485
        and ts.block_number <= 16298368) as x
order by block_number;
