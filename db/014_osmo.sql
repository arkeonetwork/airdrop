insert into chains(name,rpc_url,snapshot_start_block,snapshot_end_block,decimals)
values ('OSMO','https://daemon.osmosis.shapeshift.com:443/rpc',7455612,8049108,6);

---- create above / drop below ----
delete from chains where name = 'OSMO';
