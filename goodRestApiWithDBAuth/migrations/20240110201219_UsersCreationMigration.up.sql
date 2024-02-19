use MyLocalDB;

CREATE TABLE dbo.users (
    id bigint not null primary key identity(1,1),
    login varchar(250) not null unique,
    password varchar(250) not null
);

create table dbo.articles (
    id bigint not null primary key identity(1,1),
    title varchar(250) not null unique,
    author varchar(250) not null,
    content varchar(500) not null
);

use master;