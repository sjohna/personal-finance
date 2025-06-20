mod repo;

#[derive(Debug)]
struct Person {
    id: i32,
    name: String,
    data: Option<Vec<u8>>,
}

fn main() -> repo::Result<()> {
    let mut conn = repo::connect()?;

    dbg!(repo::init(&mut conn));

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