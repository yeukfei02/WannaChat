create table if not exists users(
  id serial primary key,
  email varchar(255) not null,
  password varchar(255) not null,
  createdDateTime timestamp not null default NOW()
);
