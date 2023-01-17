{{ template "views/unique_fox_holders_v.sql" . }}
{{ template "views/fox_transfers_v.sql" . }}
{{ template "views/my_transfers_v.sql" . }}
---- create above / drop below ----
drop view my_transfers_v;
drop view fox_transfers_v;
drop view unique_fox_holders_v;
