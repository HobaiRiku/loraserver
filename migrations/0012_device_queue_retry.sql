-- +migrate Up

alter table device_queue
    add column  retry_confirmed       boolean,
    add column  has_retry             integer,
    add column  retry_time             integer;


-- +migrate Down

alter table device_queue
    drop column  retry_confirmed       boolean,
    drop column  has_retry             integer,
    drop column  retry_time             integer;

