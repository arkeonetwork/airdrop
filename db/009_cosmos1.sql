insert into chains(name,rpc_url,snapshot_start_block,snapshot_end_block,decimals)
values ('GAIA','https://daemon.cosmos.shapeshift.com/rpc',12940505,13714445,6);

---- create above / drop below ----
delete from chains where name = 'GAIA';
