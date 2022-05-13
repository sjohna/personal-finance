create table account (
  id bigint primary key,
	name text not null,
	description text
);

-- event sourcing stuff
create type event_type as enum('create', 'update', 'delete');

create type entity_type as enum('account');

create type action_origin as enum('api-call');

create table action (
  id bigint generated always as identity primary key,
  time timestamptz not null default now(),
	origin action_origin not null,
	notes text
);

create table event (
	id bigint generated always as identity primary key,
	event_type event_type not null,
	entity_type entity_type not null,
	entity_id bigint not null,
	parameters json not null,
  action_id int references action not null
);

create unique index single_create on event (entity_id) where (event_type = 'create');
create unique index single_delete on event (entity_id) where (event_type = 'delete');

create table entity_id (
	entity_type text primary key,
	id bigint not null
);

insert into entity_id(entity_type, id)
values('all', 1);

create function nextEntityId()
returns bigint as $next_id$
declare next_id bigint;
begin
  select id into next_id from entity_id
	where entity_type = 'all';
	update entity_id set id = next_id + 1
	where entity_type = 'all';
	return next_id;
end;
$next_id$ language plpgsql;

