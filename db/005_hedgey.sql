insert into staking_contracts(address, contract_name, chain, genesis_block, height) values 
 ('0x2aa5d15eb36e5960d056e8fea6e7bb3e2a06a351', 'hedgeyNFT', 'ETH', 14496728, 0),
 ('0x2aa5d15eb36e5960d056e8fea6e7bb3e2a06a351', 'hedgeyNFT', 'GNO', 21401247 , 0);

---- create above / drop below ----
-- undo --
delete from staking_events where staking_contract = '0x2aa5d15eb36e5960d056e8fea6e7bb3e2a06a351';
delete from staking_contracts where address = '0x2aa5d15eb36e5960d056e8fea6e7bb3e2a06a351' and contract_name = 'hedgeyNFT';