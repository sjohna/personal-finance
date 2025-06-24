#![allow(dead_code)]    // TODO: disable once project is more fleshed out

mod repo;
mod common;

use tracing::{event, span, Level};
use common::Result;

#[derive(Debug)]
struct Person {
    id: i32,
    name: String,
    data: Option<Vec<u8>>,
}

fn main() -> Result<()> {
    // TODO: log to file
    let _stdout_subscriber = tracing_subscriber::fmt::init();

    tracing::info!("Look ma, I'm tracing!");

    let mut conn = repo::connect()?;

    let span = span!(Level::INFO, "Application is running");
    let _guard = span.enter();

    event!(Level::INFO, "Connected to database.");

    dbg!(repo::init(&mut conn))?;

    event!(Level::INFO, "Initialized database.");

    //drop(guard);

    println!("{:?}", repo::new_uuid());

    Ok(())
}


// let conn = Connection::open_in_memory()?;
//
// conn.execute(
//     "CREATE TABLE person (
//         id   INTEGER PRIMARY KEY,
//         name TEXT NOT NULL,
//         data BLOB
//     )",
//     (), // empty list of parameters.
// )?;
// let me = Person {
//     id: 0,
//     name: "Steven".to_string(),
//     data: None,
// };
// conn.execute(
//     "INSERT INTO person (name, data) VALUES (?1, ?2)",
//     (&me.name, &me.data),
// )?;
//
// let mut stmt = conn.prepare("SELECT id, name, data FROM person")?;
// let person_iter = stmt.query_map([], |row| {
//     Ok(Person {
//         id: row.get(0)?,
//         name: row.get(1)?,
//         data: row.get(2)?,
//     })
// })?;
//
// for person in person_iter {
//     println!("Found person {:?}", person?);
// }
// Ok(())