use std::error::Error;
use chrono::{DateTime, Utc};
use rusqlite::Row;
use uuid::Uuid;

struct ActionApplication {
    action_id: Uuid,
    apply_time: DateTime<Utc>,
}

impl TryFrom<&Row<'_>> for ActionApplication {
    type Error = Box< dyn Error>;

    fn try_from(row: &Row<'_>) -> Result<Self, Self::Error> {
        Ok(ActionApplication{
            action_id: Uuid::parse_str(row.get::<usize, String>(0)?.as_str())?,
            apply_time: row.get(1)?,
        })
    }
}