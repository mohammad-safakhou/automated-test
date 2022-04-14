create table if not exists endpoint
(
    id SERIAL primary key,
    data text,
    project_id int not null
);