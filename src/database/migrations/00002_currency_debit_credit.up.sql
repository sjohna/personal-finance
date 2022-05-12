create table currency (
	id bigint primary key,
	name text not null,
  abbreviation text not null,
  magnitude int not null constraint positive_magnitude check (magnitude >= 0)
);

create table credit (
	id bigint primary key,
	amount bigint not null,
  currency_id bigint references currency not null,
	time timestamptz not null,
	account_id bigint not null references account
);

create table debit (
	id bigint primary key,
	amount bigint not null,
  currency_id bigint references currency not null,
	time timestamptz not null,
	account_id bigint not null references account
);

-- event sourcing stuff
alter type entity_type add value 'currency';
alter type entity_type add value 'credit';
alter type entity_type add value 'debit';