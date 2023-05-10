CREATE TABLE users
(
    id uuid primary key ,
    username varchar(255) not null unique ,
    password varchar(255) not null ,
    token varchar(255) unique ,
    email varchar(255) not null unique ,
    telegram varchar(255) not null unique,
    created_at timestamp without time zone default now(),
    updated_at timestamp without time zone default now()
);

CREATE TYPE check_status AS enum ('Up','Down','Warning','Error');

CREATE TABLE checklist
(
    id uuid primary key ,
    name varchar(255) not null  unique ,
    endpoint varchar(255) not null ,
    description varchar(255) ,
    status check_status ,
    last_update timestamp without time zone default now(),
    icon varchar(255),
    user_id uuid ,
    created_at timestamp without time zone default now(),
    updated_at timestamp without time zone default now() ,
    foreign key (user_id) references users (id)
);

CREATE TYPE alert_status AS enum ('Send','New');

CREATE TABLE alert
(
    id uuid primary key ,
    check_id uuid ,
    date timestamp without time zone default now(),
    name varchar(255) ,
    description varchar(255),
    status alert_status ,
    created_at timestamp without time zone default now(),
    foreign key (check_id) references checklist (id)
);



