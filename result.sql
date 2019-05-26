create table result
(
	id serial not null
		constraint result_pk
			primary key,
	peer varchar,
	content text,
	time timestamp default CURRENT_TIMESTAMP
);

alter table result owner to myuser;

