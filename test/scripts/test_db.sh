#!/bin/bash

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest
echo "Postgre sql starting..."
sleep 3

echo "product app veritabanı oluşturuluyor..."
docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE productapp"
sleep 3
echo "Database created successfully"

docker exec -it postgres-test psql -U postgres -d productapp  -c "
create table if not exists products (
    id bigserial primary key,
    name varchar(255) not null,
    price double precision not null,
    discount double precision ,
    store varchar(255) not null
);  "

sleep 3 

echo "Table products created"