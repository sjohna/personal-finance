CREATE TABLE account (
  id bigint PRIMARY KEY,
	account_name text not null,
	account_desc text
);

CREATE TABLE credit (
	id bigint PRIMARY KEY,
	creditAmount DECIMAL(20,10) not null,
	time timestamptz not null,
	account_id int NOT NULL REFERENCES account
);

CREATE TABLE debit (
	id bigint PRIMARY KEY,
	debitAmount DECIMAL(20,10) not null,
	time timestamptz not null,
	account_id int NOT NULL REFERENCES account
);

CREATE TABLE transaction (
	id bigint PRIMARY KEY,
	credit_id int NOT NULL UNIQUE REFERENCES credit,
	debit_id int NOT NULL UNIQUE REFERENCES debit
);

-- event sourcing stuff
create type event_type as enum('create', 'update', 'delete');

create type entity_type as enum('account', 'debit', 'credit', 'transaction');

create type action_origin as enum('api-call');

CREATE TABLE action (
  id serial PRIMARY KEY,
  time timestamptz not null default now(),
	action_origin action_origin not null,
	notes text
);

CREATE TABLE event (
	id serial PRIMARY KEY,
	event_type event_type not null,
	entity_type entity_type not null,
	parameters json not null,
  action_id int references action not null
);

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