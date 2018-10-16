create table if not exists users(
  id serial primary key,
  email varchar(255) not null,
  password varchar(255) not null,
  createdDateTime timestamp not null default NOW()
);

create table if not exists groups(
  id serial primary key,
  label varchar(255) not null,
  createdDateTime timestamp not null default NOW()
);

create table if not exists memberships(
  id serial primary key,
  user_id integer not null references users(user_id),
  group_id integer not null references groups(group_id),
  createdDateTime timestamp not null default NOW()
);
