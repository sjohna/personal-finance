CREATE TABLE account (
  id serial PRIMARY KEY,
	account_name text not null,
	account_desc text
);

CREATE TABLE credit (
	id serial PRIMARY KEY,
	creditAmount DECIMAL(20,10) not null,
	time timestamptz not null,
	account_id int NOT NULL REFERENCES account
);

CREATE TABLE debit (
	id serial PRIMARY KEY,
	debitAmount DECIMAL(20,10) not null,
	time timestamptz not null,
	account_id int NOT NULL REFERENCES account
);

CREATE TABLE transaction (
	id serial PRIMARY KEY,
	credit_id int NOT NULL UNIQUE REFERENCES credit,
	debit_id int NOT NULL UNIQUE REFERENCES debit
);