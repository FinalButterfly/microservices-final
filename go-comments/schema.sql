begin;
drop table if exists comments;

create table comments(
	id bigserial primary key,
  parentId bigint not null,
  articleId bigint not null,
  pubtime bigint not null,
  content text not null,
  profane boolean not null
);

commit;