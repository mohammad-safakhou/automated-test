create table if not exists auth_info
(
    id SERIAL primary key ,
    private_key text
);

create table if not exists accounts
(
    id SERIAL primary key,
    first_name text,
    last_name text,
    phone_number text,
    email text,
    username text,
    password text,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

create table if not exists projects
(
    id SERIAL primary key,
    title text ,
    is_active bool,
    expire_at time,
    account_id int NOT NULL,
    notifications jsonb,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    foreign key (account_id) references accounts(id)
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

create table if not exists trace_routes
(
    id SERIAL primary key,
    data text,
    project_id int not null,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    foreign key (project_id) references projects(id)
    );

create table if not exists net_cats
(
    id SERIAL primary key,
    data text,
    project_id int not null,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    foreign key (project_id) references projects(id)
    );

create table if not exists pings
(
    id SERIAL primary key,
    data text,
    project_id int not null,

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

    foreign key (project_id) references projects(id)
    );

create table if not exists page_speeds
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
