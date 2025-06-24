create table action (
    id text primary key,
    create_time text not null,
    action_time text not null,
    description text not null
);

create table delta (
    id text primary key,
    version int not null,
    action_id text not null,
    create_time text not null,
    delta_time text not null,
    delta_type text not null,
    entity_type text not null,
    scope text not null,
    entity_id text,
    data text not null,
    foreign key (action_id) references action(id)
);

create table action_application (
    action_id text primary key,
    apply_time test not null,
    foreign key (action_id) references action(id)
);

create table delta_application (
    delta_id text primary key,
    action_application_id text not null,
    apply_time text not null,
    foreign key (action_application_id) references action_application(id),
    foreign key (delta_id) references delta(id)
);
