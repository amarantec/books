CREATE TABLE IF NOT EXISTS categories (
  id    serial primary key,
  name  text not null,
  url   text not null
);

CREATE TABLE IF NOT EXISTS users (
  id        serial primary key,
  name      text not null unique,
  email     text not null unique,
  password text not null
);

CREATE TABLE IF NOT EXISTS books (
  id          serial primary key,
  title       text not null unique,
  description text not null,
  genre       text[]  not null,
  author      text[]  not null,
  image_url   text not null,
  category_id integer references categories(id),
  user_id     integer references users(id)
);

