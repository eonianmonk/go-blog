create type state as enum('readable', 'deleted','hidden');

create table if not exists posts (
    id serial primary key,
    title text,
    author text,
    post_state state, 
    body text
);