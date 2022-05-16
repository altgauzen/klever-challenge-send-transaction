create table details(
    id serial primary key,
    address varchar,
    balance varchar,
    totalTX varchar
);

INSERT INTO details(address, balance, totalTX) VALUES
('bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r', '144011754', '17000'),
('de1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r', '288011754', '23000');