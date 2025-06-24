use std::error::Error;
use chrono::{DateTime, Utc};
use json::JsonValue;
use json::object::Object;
use rusqlite::{Row, ToSql};
use rusqlite::types::{FromSql, FromSqlError, FromSqlResult, ToSqlOutput, Value, ValueRef};
use uuid::Uuid;

enum DeltaType {
    Create,
    Update,
    Delete,
}

impl FromSql for DeltaType {
    fn column_result(value: ValueRef<'_>) -> FromSqlResult<Self> {
        match value.as_str() {
            Ok("create") => Ok(DeltaType::Create),
            Ok("update") => Ok(DeltaType::Update),
            Ok("delete") => Ok(DeltaType::Delete),
            Ok(_) => Err(FromSqlError::Other(Box::from("invalid value for delta type"))),   // TODO: static instances of errors?
            Err(e) => Err(e),
        }
    }
}

impl ToSql for DeltaType {
    fn to_sql(&self) -> rusqlite::Result<ToSqlOutput<'_>> {
        match self {
            DeltaType::Create => Ok(ToSqlOutput::Owned(Value::Text("create".to_owned()))), // TODO: is this creating a ton of copies unnecessarily?
            DeltaType::Update => Ok(ToSqlOutput::Owned(Value::Text("update".to_owned()))),
            DeltaType::Delete => Ok(ToSqlOutput::Owned(Value::Text("delete".to_owned()))),
        }
    }
}

enum Scope {
    Single,
    Bulk,
}

impl FromSql for Scope {
    fn column_result(value: ValueRef<'_>) -> FromSqlResult<Self> {
        match value.as_str() {
            Ok("single") => Ok(Scope::Single),
            Ok("bulk") => Ok(Scope::Bulk),
            Ok(_) => Err(FromSqlError::Other(Box::from("invalid value for scope"))),   // TODO: static instances of errors?
            Err(e) => Err(e),
        }
    }
}

impl ToSql for Scope {
    fn to_sql(&self) -> rusqlite::Result<ToSqlOutput<'_>> {
        match self {
            Scope::Single => Ok(ToSqlOutput::Owned(Value::Text("single".to_owned()))), // TODO: is this creating a ton of copies unnecessarily?
            Scope::Bulk => Ok(ToSqlOutput::Owned(Value::Text("bulk".to_owned()))),
        }
    }
}

struct Delta {
    id: Uuid,
    version: u64,
    action_id: Uuid,
    create_time: DateTime<Utc>,
    delta_time: DateTime<Utc>,
    delta_type: DeltaType,
    entity_type: String, // TODO: enum
    scope: Scope,
    entity_id: Option<Uuid>,
    data: Option<Object>,
}

impl TryFrom<&Row<'_>> for Delta {
    type Error = Box<dyn Error>;

    fn try_from(row: &Row<'_>) -> Result<Self, Self::Error>
    {
        Ok(Delta{
            id: Uuid::parse_str(row.get::<usize, String>(0)?.as_str())?,    // TODO: helper function for this?
            version: row.get(1)?,
            action_id: Uuid::parse_str(row.get::<usize, String>(2)?.as_str())?,
            create_time: row.get(3)?,
            delta_time: row.get(4)?,
            delta_type: row.get(5)?,
            entity_type: row.get(6)?,
            scope: row.get(7)?,
            entity_id: match row.get::<usize, Option<String>>(8)? {         // TODO: helper function for this?
                Some(s) => Some(Uuid::parse_str(s.as_str())?),
                None => None,
            },
            data: { // TODO: is this the best way to do this kind of nested pattern matching
                let mut value: Option<Object> = None;

                if let Some(s) = row.get::<usize, Option<String>>(8)? {
                    match json::parse(&s) {
                        Ok(JsonValue::Object(obj)) => value = Some(obj),
                        Ok(JsonValue::Null) => value = None,
                        Ok(_) => return Err(Box::from("invalid kind of json value in data")),
                        Err(e) => return Err(e.into()),
                    }
                }

                value
            }
        })
    }

}