use rusqlite::{params, version, Connection, Transaction};
use crate::repo::init;

fn insert_version(conn: &Connection, version: i64) -> init::Result<()> {
    //language=SQL
    let insert_version = r#"
        insert into version(version, apply_time)
        values(?1, datetime())
    "#;

    conn.execute(insert_version, params![version])?;

    Ok(())
}

pub fn version_1(conn: &mut Connection) -> init::Result<()> {
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

fn version_1_up(tx: &Transaction) -> init::Result<()> {
    create_action_table(tx)?;
    create_delta_table(tx)?;
    create_action_application_table(tx)?;
    create_delta_application_table(tx)?;

    Ok(())
}

fn create_action_table(tx: &Transaction) -> init::Result<()> {
    //language=SQL
    let create_action_table = r#"
        create table action (
            id text primary key,
            create_time text not null,
            action_time text not null,
            description text not null
        )
    "#;

    tx.execute(create_action_table, params![])?;

    Ok(())
}

fn create_delta_table(tx: &Transaction) -> init::Result<()> {
    //language=SQL
    let create_delta_table = r#"
        create table delta (
            id text primary key,
            action_id text not null,
            create_time text not null,
            delta_time text not null,
            delta_type text not null,
            entity_type text not null,
            scope text not null,
            entity_id text,
            data text not null,
            version int not null,
            foreign key (action_id) references action(id)
        )
    "#;

    tx.execute(create_delta_table, params![])?;

    Ok(())
}

fn create_action_application_table(tx: &Transaction) -> init::Result<()> {
    //language=SQL
    let create_action_application_table = r#"
        create table action_application (
            action_id text primary key,
            apply_time test not null,
            foreign key (action_id) references action(id)
        )
    "#;

    tx.execute(create_action_application_table, params![])?;

    Ok(())
}

fn create_delta_application_table(tx: &Transaction) -> init::Result<()> {
    //language=SQL
    let create_delta_application_table = r#"
        create table delta_application (
            delta_id text primary key,
            action_application_id text not null,
            apply_time text not null,
            foreign key (action_application_id) references action_application(id),
            foreign key (delta_id) references delta(id)
        )
    "#;

    tx.execute(create_delta_application_table, params![])?;

    Ok(())
}