create table if not exists hooks (
  hook_id uuid primary key
  , display_name varchar(200) not null
  , description text not null
  , url text not null
  , method varchar(20) not null
  , body jsonb
  , headers jsonb
);

create table if not exists hook_schedules (
  hook_schedule_id uuid primary key
  , hook_id uuid not null
  , display_name varchar(200) not null
  , description text not null
  -- hook execution frequency
  -- 1: minutes, 2: hours, 3: day, 4: month
  , schedule_interval_unit smallint not null
  , schedule_interval_value int not null
  -- hook execution time
  , schedule_time_month smallint
  , schedule_time_day smallint
  , schedule_time_hour smallint
  , schedule_time_minute smallint
);

alter table hook_schedules 
  add constraint fk_hook_schedules_hook_id 
  foreign key (hook_id) references hooks(hook_id) 
  on delete cascade on update cascade;


create table if not exists hook_histories (
  hook_history_id uuid primary key
  , hook_id uuid not null
  , hook_schedule_id uuid not null
  -- 1: requested, 2: processing, 3: pending, 101: succeeded, 102:failed
  , status smallint not null
  , started_at timestamp not null
  , updated_at timestamp not null
  , ended_at timestamp
  , hook_snapshot jsonb
  , schedule_snapshot jsonb
);
