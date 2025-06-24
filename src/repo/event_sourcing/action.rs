use chrono::{DateTime, Utc};
use rusqlite::Row;
use uuid::Uuid;

// TODO: action base vs. action with deltas
struct Action {
    id: Uuid,
    create_time: DateTime<Utc>,
    action_time: DateTime<Utc>,
    description: String,
}

impl TryFrom<&Row<'_>> for Action {  // Q: why do I need an explicit lifetime here?
    type Error = Box<dyn std::error::Error>;

    fn try_from(row: &Row<'_>) -> Result<Self, Self::Error> {
        Ok(Action{
            id: Uuid::parse_str(row.get::<usize, String>(0)?.as_str())?,
            create_time: row.get(1)?,
            action_time: row.get(2)?,
            description: row.get(3)?,
        })
    }
}

