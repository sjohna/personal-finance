create view debits_and_credits as
select *
from
(
	select *, 0 as type
	from debit

	union 

	select *, 1 as type
	from credit
) debit_credit_union;
