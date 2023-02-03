with params as (
  select
    chains.name,
    'osmovaloper1xf9zpq5kpxks49cg606tzd8qstaykxgt2vs0d5' as validator,
    snapshot_start_block,
    snapshot_end_block
  from
    chains
  where
    chains.name = 'OSMO'
),
staking_events as (
  select
    id,
    txhash,
    delegator as account,
    amount as delta,
    block_number
  from
    cosmos_staking_events
  where
    chain = (
      select
        chain
      from
        params
    )
    and validator = (
      select
        validator
      from
        params
    )
  order by
    block_number
),
averageable as (
  select
    account,
    cumulative_balance,
    block_number,
    blocks_between,
    blocks_between * lag(cumulative_balance, 1, 0) over (partition by account) as avg_over_blocks
  from
    (
      select
        account,
        SUM(delta) over (
          partition by account
          order by
            block_number
        ) as cumulative_balance,
        ts.block_number,
        block_number - lag(block_number, 1, block_number) over (
          partition by account
          order by
            block_number
        ) as blocks_between
      from
        (
          select
            account,
            delta,
            block_number
          from
            staking_events
          where
            block_number >= (
              select
                snapshot_start_block
              from
                params
            )
            and block_number <= (
              select
                snapshot_end_block
              from
                params
            )
          union
          -- starting balance
          (
            select
              account,
              sum(delta),
              (
                select
                  snapshot_start_block
                from
                  params
              ) as block_number
            from
              staking_events
            where
              block_number <= (
                select
                  snapshot_start_block
                from
                  params
              )
            group by
              account
            order by
              block_number
          )
          union
          -- ending balance
          (
            select
              account,
              sum(delta),
              (
                select
                  snapshot_end_block
                from
                  params
              ) as block_number
            from
              staking_events
            where
              block_number <= (
                select
                  snapshot_end_block
                from
                  params
              )
            group by
              account
            order by
              block_number
          )
        ) as ts
      where
        ts.block_number >= (
          select
            snapshot_start_block
          from
            params
        )
        and ts.block_number <= (
          select
            snapshot_end_block
          from
            params
        )
    ) as x
  order by
    block_number
)
select
  account,
  sum(avg_over_blocks) / (
    (
      select
        snapshot_end_block
      from
        params
    ) - (
      select
        snapshot_start_block
      from
        params
    )
  ) avg_hold
from
  averageable
group by
  account
having
  sum(avg_over_blocks) / (
    (
      select
        snapshot_end_block
      from
        params
    ) - (
      select
        snapshot_start_block
      from
        params
    )
  ) > 400
order by
  avg_hold desc;