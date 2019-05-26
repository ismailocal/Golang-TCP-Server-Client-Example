create table log
(
	id serial not null
		constraint log_pk
			primary key,
	server_ip varchar,
	token varchar,
	time timestamp default CURRENT_TIMESTAMP
);

alter table log owner to myuser;

