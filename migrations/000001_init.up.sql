/* migrate create -ext sql -dir migrations create_users */
/* migrate  -path migrations -database "postgres://postgres:Cain-666@localhost/machclub_dev?sslmode=disable" up */
/* update schema_migrations set version='000001', dirty=false; */
-- CREATE TABLE users
-- (
--   id serial not null unique,
--   email varchar(255) not null unique,
--   username varchar(255) not null unique,
--   password_hash varchar(255) not null
-- );
/* реализовать на фронте в первую очередь */
/* реализовать image */
CREATE TABLE news (
  id          serial not null unique,
  title       varchar(255) not null,
  description varchar(255) not null,
  text        text not null
)
/*TODO на будущее. Когда будут реализоваться статьи а-ля как на хабре*/
-- CREATE TABLE users_lists
-- (
--   id serial not null unique,
--   user_id int references users (id) on delete cascade not null,
--   article_id int references article_lists (id) on delete cascade not null
-- );
-- CREATE TABLE articles_lists
-- (
--   id serial not null unique,
--   title varchar(255) not null,
--   description varchar(255)
-- );