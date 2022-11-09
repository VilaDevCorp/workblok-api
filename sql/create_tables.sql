create database if not exists sensei;

use sensei;


create table if not exists user (
  id uuid DEFAULT uuid () PRIMARY KEY,
  username varchar(100) not null,
  pass char(60) not null,
  dans int default 0,
  creation_date date default now()
);

create table if not exists activity (
  id uuid DEFAULT uuid () PRIMARY KEY,
  name varchar(150) not null,
  size int default 1,
  user_id uuid,
  creation_date date default now(),
  FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE SET NULL
);

create table if not exists task (
  id uuid DEFAULT uuid () PRIMARY KEY,
  activity_id uuid not null,
  due_date date not null,
  creation_date date default now()
);

create table if not exists planning (
  id uuid DEFAULT uuid () PRIMARY KEY,
  name varchar(150) not null,
  user_id uuid,
  creation_date date default now(),
  FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE
);

create table if not exists planning_task (
  id uuid DEFAULT uuid () PRIMARY KEY,
  activity_id uuid not null,
  planning_id uuid not null,
  weekday int not null,
  creation_date date default now(),
  FOREIGN KEY (activity_id) REFERENCES activity(id) ON DELETE CASCADE,
  FOREIGN KEY (planning_id) REFERENCES planning(id) ON DELETE CASCADE
);
