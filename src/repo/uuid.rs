use std::ops::Deref;
use std::sync::Mutex;
use std::time::{SystemTime, UNIX_EPOCH};
use uuid::{ContextV7, Timestamp, Uuid};

static CONTEXT: Mutex<ContextV7> = Mutex::new(ContextV7::new());

pub fn new_uuid() -> Uuid {
    let ctx = CONTEXT.lock().unwrap();
    let now = SystemTime::now().duration_since(UNIX_EPOCH).expect("Failed to get current unix time!");
    let ts = Timestamp::from_unix(ctx.deref(), now.as_secs(), now.subsec_nanos());

    Uuid::new_v7(ts)
}