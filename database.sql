create table if not exists projects
(
    id SERIAL primary key,
    title text ,
    is_active bool,
    expire_at time,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
    );

create table if not exists endpoints
(
    id SERIAL primary key,
    data text,
    project_id int not null,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    foreign key (project_id) references projects(id)
    );

create table if not exists datacenters
(
    id SERIAL primary key,
    baseurl text not null ,
    title text not null ,
    connection_rate int ,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
    );

create table if not exists relation_datacenters
(
    id SERIAL primary key,
    endpoint_id int ,
    datacenter_id int ,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    foreign key (endpoint_id) references endpoints(id),
    foreign key (datacenter_id) references datacenters(id)
    );

create table if not exists schedulings
(
    id SERIAL primary key,
    project_id int NOT NULL,
    duration text NOT NULL,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    foreign key (project_id) references projects(id)
    );