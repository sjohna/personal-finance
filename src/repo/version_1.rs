use rusqlite::{params, Connection, Transaction};

use crate::common::Result;

fn insert_version(conn: &Connection, version: i64) -> Result<()> {
    //language=SQL
    let insert_version = r#"
        insert into version(version, apply_time)
        values(?1, datetime())
    "#;

    conn.execute(insert_version, params![version])?;

    Ok(())
}

pub fn version_1(conn: &mut Connection) -> Result<()> {
    let tx = conn.transaction()?;

    let result = version_1_up(&tx);

    match result {
        Ok(()) => {
            tx.commit()?;
            insert_version(conn, 1)?;
            Ok(())
        }
        Err(e) => {
            tx.rollback()?;
            Err(e)
        }
    }
}

fn version_1_up(tx: &Transaction) -> Result<()> {
    let sql = include_str!("./versions/version_1.sql");

    tx.execute_batch(sql)?;

    Ok(())
}