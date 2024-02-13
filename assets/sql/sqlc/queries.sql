-- name: CreatePost :one
insert into posts (title, author, post_state, body)
values ($1, $2, $3, $4)
returning id;

-- name: GetHeaders :many
select id,title,author from posts 
where post_state = 'readable'
limit $1;

-- name: GetAllHeaders :many
select id,title,author from posts 
where post_state = 'readable';

-- name: Delete :one
delete from posts where id = $1 returning id;

-- name: Update :one
update posts 
set title = $1,
    author = $2,
    body = $3
where id = $4
returning id,title,author;

-- name: GetPost :one 
select * from posts 
where id = $1;
