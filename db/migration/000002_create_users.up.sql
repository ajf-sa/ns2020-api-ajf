create table if not exists users(
id serial  primary key,
username varchar ,
pass varchar,
email text 

);

alter table users add first_name varchar;

alter table users add last_name varchar;

