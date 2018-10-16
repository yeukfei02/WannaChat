create table if not exists users(
  userId serial primary key,
  email varchar(255) not null,
  password varchar(255) not null,
  createdDateTime timestamp not null default NOW()
);

create table if not exists groups(
  groupId serial primary key,
  label varchar(255) not null,
  createdDateTime timestamp not null default NOW()
);

create table if not exists memberships(
  membershipsId serial primary key,
  userId integer not null references users(userId),
  groupId integer not null references groups(groupId),
  createdDateTime timestamp not null default NOW()
);
