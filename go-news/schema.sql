begin;
drop table if exists posts;

create table posts(
	id bigserial primary key,
	title text not null,
	content text not null,
	pubtime bigint not null,
	link text not null,
	UNIQUE (title, link)
);

commit;