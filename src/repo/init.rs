use std::error::Error;
use rusqlite;
use rusqlite::{params, Connection, Params, Row};
use crate::repo::version_1;

const CURRENT_VERSION: u32 = 1;

// TODO: reorganize all of this
pub type Result<T> = std::result::Result<T, Box<dyn Error>>;

#[derive(Clone, Default)]
pub struct Version {
    version: u32,
    time_applied: String,   // TODO: look at chrono for time types?
}

impl FromRow for Version {
    fn from_row(row: &Row) -> Result<Version> {
        Ok(Version {
            version: row.get(0)?,
            time_applied: row.get(1)?,
        })
    }
}

pub fn connect() -> Result<Connection> {
    match Connection::open("database/test.db") {
        Ok(conn) => Ok(conn),
        Err(e) => Err(e)?,
    }
}

pub fn init(conn: &mut Connection) -> Result<()> {
    // do a number of things:
    //  - if version table does not exist, create it and init version 1
    //  - (future) if version table exists, check if we're at the latest version
    //  - if we are, great
    //  - if not, update to that version

    let version = check_or_init_version(conn)?;

    if version.version == 0 {
        version_1::version_1(conn)?;
    }

    Ok(())
}

// TODO: consider whether this should return a Box<Version>, or just a Version
fn check_or_init_version(conn: &Connection) -> Result<Box<Version>> {
    if conn.table_exists(None, "version")? {
        //language=SQL
        let get_version_query = r#"
            select version, apply_time
            from version
            order by version desc
            limit 1
        "#;

        let version: Option<Box<Version>> = select_zero_or_one(conn, get_version_query, [])?;

        match version {
            Some(version) => Ok(version),
            None => Ok(Box::new(Version::default())),
        }
    } else {
        create_version_table(conn)?;

        Ok(Box::new(Version::default()))
    }
}

fn create_version_table(conn: &Connection) -> Result<()> {
    //language=SQL
    let create_query = r#"
        create table version(
            version integer primary key,
            apply_time text
        )
    "#;

    conn.execute(create_query, params![])?;

    Ok(())
}

// functions needed:
// select_one_or_more
// select_zero_or_more
// select_one
// select_zero_or_one
//
// for both primitives (T: FromSql) and structures (T: FromRow)

pub fn select_zero_or_one<T: FromRow, P: Params>(conn: &rusqlite::Connection, query: &str, params: P) -> Result<Option<Box<T>>> {
    let mut statement = conn.prepare(query)?;
    let mut returned = statement.query_and_then(params, T::from_row)?;  // Q: what's the difference between this and query_map?

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

pub fn select_one<T: FromRow, P: Params>(conn: &rusqlite::Connection, query: &str, params: P) -> Result<Box<T>> {
    let returned = select_zero_or_one(conn, query, params)?;

    match returned {
        Some(ret) => Ok(ret),
        None => Err(Box::from("no rows returned")),
    }
}

pub fn select_one_or_more<T: FromRow, P: Params>(conn: &rusqlite::Connection, query: &str, params: P) -> Result<Vec<T>> {
    let returned = select_zero_or_more(conn, query, params)?;
    if returned.len() == 0 {
        return Err(Box::from("expected at least one row"));   // TODO: standardize these errors
    }

    Ok(returned)
}

pub fn select_zero_or_more<T: FromRow, P: Params>(conn: &rusqlite::Connection, query: &str, params: P) -> Result<Vec<T>> {
    let mut statement = conn.prepare(query)?;
    let returned = statement.query_and_then(params, T::from_row)?;  // Q: what's the difference between this and query_map?

    let mut ret = Vec::new();
    for t in returned {
        ret.push(t?);
    }

    Ok(ret)
}


pub trait FromRow {
    fn from_row(row: &rusqlite::Row) -> Result<Self> where Self: Sized; // TODO: boxed references?
}

