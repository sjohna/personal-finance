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