# migrate -path ./ -database "mysql://root:root@tcp(localhost:3307)/emailqueue" up;

create table email_information
(
    id int,
    sent boolean default true null,
    created_at timestamp null,
    updated_at timestamp null,
    email varchar(150) null,
    name varchar(150) null,
    last_name varchar(150) null
);

create unique index email_information_id_uindex
    on email_information (id);

alter table email_information
    add constraint email_information_pk
        primary key (id);

alter table email_information modify id int auto_increment;

