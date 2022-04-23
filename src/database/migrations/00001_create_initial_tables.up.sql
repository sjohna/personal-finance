CREATE TABLE account (
  id serial PRIMARY KEY,
	account_name VARCHAR,
	account_desc VARCHAR
);

CREATE TABLE credit (
	id serial PRIMARY KEY,
	creditAmount DECIMAL(20,10),
	time timestamp,
	account_id int NOT NULL REFERENCES account
);

CREATE TABLE debit (
	id serial PRIMARY KEY,
	debitAmount DECIMAL(20,10),
	time timestamp,
	account_id int NOT NULL REFERENCES account
);

CREATE TABLE transaction (
	id serial PRIMARY KEY,
	credit_id int NOT NULL UNIQUE REFERENCES credit,
	debit_id int NOT NULL UNIQUE REFERENCES debit
);