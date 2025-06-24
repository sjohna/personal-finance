use std::error::Error;
use rusqlite;
use rusqlite::{params, Connection, Row};
use crate::repo::{select_zero_or_one, version_1};

use crate::common::Result;

const CURRENT_VERSION: u32 = 1;

#[derive(Clone, Default)]
pub struct Version {
    version: u32,
    time_applied: chrono::DateTime<chrono::Utc>,
}

impl TryFrom<&Row<'_>> for Version {
    type Error = Box<dyn Error>;

    fn try_from(row: &Row) -> Result<Version> {
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

