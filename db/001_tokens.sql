create table tokens
(
    address text  not null 
     constraint tokens_pk
      primary key,
    name     text   not null,
    symbol   text   not null,
    chain    text   not null,
    genesis_block numeric check ( genesis_block >= 0 ),
    height    numeric not null check ( height >= 0 )
);

-- mainnet 
  -- fox
insert into tokens values ('0xc770EEfAd204B5180dF6a14Ee197D99d808ee52d', 'FOX', 'FOX', 'ETH', 7446830, 0);
  -- foxy
insert into tokens values ('0xDc49108ce5C57bc3408c3A5E95F3d864eC386Ed3', 'FOX Yieldy', 'FOXy', 'ETH', 14476930, 0);
  -- fox <> ETH LP
insert into tokens values ('0x470e8de2eBaef52014A47Cb5E6aF86884947F08c', 'UNI-V2', 'Uniswap V2', 'ETH', 10104463, 0);
  -- tfox
insert into tokens values ('0x808d3e6b23516967ceae4f17a5f9038383ed5311', 'TokemaktFOX', 'tFOX', 'ETH', 13690577, 0);
  -- scfox
insert into tokens values ('0x04979ccCC2F854E167DAAd0AE095Da49eAE4842E', 'FOX $0.80 Success Token March 2024', 'scFOX0324', 'ETH', 14174869, 0);

-- gnosis
  -- fox
insert into tokens values ('0x21a42669643f45bc0e086b8fc2ed70c23d67509d', 'FOX on xDai', 'FOX', 'GNO', 16988521, 0);

-- Polygon
  -- fox
insert into tokens values ('0x65a05db8322701724c197af82c9cae41195b0aa8', 'FOX', 'FOX', 'POLY', 14283484 , 0);

create table staking
(
    address text  not null 
     constraint staking_pk
      primary key,
    contract_name   text   not null,
    chain    text   not null,
    genesis_block numeric check ( genesis_block >= 0 ),
    height    numeric not null check ( height >= 0 )
);

-- mainnet 
  -- LP stakingrewards
insert into staking values ('0xc14eaa8284feff79edc118e06cadbf3813a7e555', 'stakingrewards', 'ETH', 15941059, 0);

---- create above / drop below ----
-- undo --
drop table tokens;
drop table staking;

