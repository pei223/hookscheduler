create table if not exists hooks (
  hook_id uuid primary key
  , display_name varchar(200) not null
  , description text not null
  , url text not null
  , method varchar(20) not null
  , body jsonb
  , headers jsonb
);

create type schedule_frequency_unit as enum ('every_minute', 'every_hour', 'every_day', 'every_month', 'every_year');

create table if not exists hook_schedules (
  hook_schedule_id uuid primary key
  , hook_id uuid not null
  , display_name varchar(200) not null
  , description text not null
  -- hook execution frequency
  -- 1: every minutes, 2: every hours, 3: every day, 5: every month, 6: every year
  , schedule_frequency_unit schedule_frequency_unit not null
  -- hook execution time
  , schedule_time_month smallint not null
  , schedule_time_day smallint not null
  , schedule_time_hour smallint not null
  , schedule_time_minute smallint not null
  , schedule_time_second smallint not null
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


create table if not exists hook_results (
  hook_history_id uuid primary key
  , http_status_code int not null
  , response_body jsonb
  , response_headers jsonb
);

alter table hook_results
  add constraint fk_hook_results_hook_history_id
  foreign key (hook_history_id) references hook_histories(hook_history_id)
  on delete cascade on update cascade;
