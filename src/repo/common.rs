use std::error::Error;
use rusqlite::{Params, Row};
use crate::common::Result;

// functions needed:
// select_one_or_more
// select_zero_or_more
// select_one
// select_zero_or_one
//
// for both primitives (T: FromSql) and structures (T: FromRow)

pub fn select_zero_or_one<T: for<'a, 'b> TryFrom<&'a Row<'b>, Error = Box<dyn Error>>, P: Params>(conn: &rusqlite::Connection, query: &str, params: P) -> Result<Option<Box<T>>> { // Q: why do I need explicit lifetimes here? And what's that for<'a? deal?
    let mut statement = conn.prepare(query)?;
    let mut returned = statement.query_and_then(params, |value: &Row| T::try_from(value))?;  // Q: what's the difference between this and query_map?

    let row = returned.next();

    if row.is_none() {
        return Ok(None);
    }

    let ret = Box::new(row.unwrap()?);

    match returned.next() {
        None => Ok(Some(ret)),
        Some(_) => Err(Box::from("expected zero or one rows")),
    }
}

pub fn select_one<T: for<'a, 'b> TryFrom<&'a Row<'b>, Error = Box<dyn Error>>, P: Params>(conn: &rusqlite::Connection, query: &str, params: P) -> Result<Box<T>> {
    let returned = select_zero_or_one(conn, query, params)?;

    match returned {
        Some(ret) => Ok(ret),
        None => Err(Box::from("no rows returned")),
    }
}

pub fn select_one_or_more<T: for<'a, 'b> TryFrom<&'a Row<'b>, Error = Box<dyn Error>>, P: Params>(conn: &rusqlite::Connection, query: &str, params: P) -> Result<Vec<T>> {
    let returned = select_zero_or_more(conn, query, params)?;
    if returned.len() == 0 {
        return Err(Box::from("expected at least one row"));   // TODO: standardize these errors
    }

    Ok(returned)
}

pub fn select_zero_or_more<T: for<'a, 'b> TryFrom<&'a Row<'b>, Error = Box<dyn Error>>, P: Params>(conn: &rusqlite::Connection, query: &str, params: P) -> Result<Vec<T>>
where
{
    let mut statement = conn.prepare(query)?;
    let returned = statement.query_and_then(params, |value: &'_ Row<'_>| T::try_from(value))?;  // Q: what's the difference between this and query_map?

    let mut ret = Vec::new();
    for t in returned {
        ret.push(t?);
    }

    Ok(ret)
}