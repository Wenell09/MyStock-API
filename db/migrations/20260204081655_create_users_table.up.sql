create table if not exists users (
    id uuid primary key references auth.users(i)d on delete cascade,
    email varchar(255) not null unique,
    name varchar(255),
    role varchar(50) default 'admin',
    created_at timestamp default now(),
    updated_at timestamp default now()
);
