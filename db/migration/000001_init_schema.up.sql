set timezone = 'Asia/Ho_Chi_Minh';
create extension if not exists "uuid-ossp";
create table if not exists books (
    book_id uuid default uuid_generate_v4(),
    code varchar(25) not null,
    title varchar(255) not null,
    publisher varchar(255) not null,
    public_date date not null,
    version int not null default 1,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    primary key (book_id)
    );

create index on books (code);

insert into books (code, title, publisher, public_date) values ('BOOK01', 'Học làm người', 'NXB Xã hội', '2000-12-31');
insert into books (code, title, publisher, public_date, created_at, updated_at) values ('BOOK02', 'Binh pháp thầy Huấn', 'NXB Lao động', '2022-12-31', current_timestamp, current_timestamp);
insert into books (code, title, publisher, public_date, created_at, updated_at) values ('BOOK03', 'Gái nghành', 'NXB Nghành', '1999-09-06', current_timestamp, current_timestamp);
