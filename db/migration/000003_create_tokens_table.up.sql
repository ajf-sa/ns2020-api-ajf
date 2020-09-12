create table if not exists tokens(
    id serial primary key,
    uuid bigserial,
    userid int
);
