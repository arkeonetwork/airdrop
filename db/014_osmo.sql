insert into chains(name,rpc_url,lcd_url,snapshot_start_block,snapshot_end_block,decimals)
values ('OSMO','https://daemon.osmosis.shapeshift.com:443/rpc','https://daemon.osmosis.shapeshift.com:443/lcd',7117606,8049108,6);

---- create above / drop below ----
delete from chains where name = 'OSMO';
