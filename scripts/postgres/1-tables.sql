create table if not exists public.settings
(
    ID                  serial
        constraint settings_pk
            primary key,
    limits_trigger      boolean default false not null,
    water_level_limit   float                 not null,
    water_amount_limit  float                 not null,
    moist_limit         float                 not null,
    scheduled_trigger   boolean default false not null,
    hour_range          int                   not null,
    location            varchar               not null,
    irrigation_duration boolean default false not null,
    chart_type          boolean default false not null,
    language            boolean default false not null,
    theme               boolean default false not null,
    lat                 float                 not null,
    lon                 float                 not null
);

create unique index settings_id_uindex
    on settings (ID);

create table if not exists public.measurements
(
    ID              serial
        constraint measurements_pk
            primary key,
    timestamp       timestamp             not null,
    hum             float                 not null,
    temp            float                 not null,
    moist           float                 not null,
    with_irrigation boolean default false not null
);

create unique index measurements_id_uindex
    on measurements (ID);

create table if not exists public.irrigation_history
(
    ID              serial
        constraint irrigation_history_pk
            primary key,
    timestamp       timestamp not null,
    water_level     float     not null,
    water_amount    float     not null,
    water_overdrawn float     not null
);

create unique index irrigation_history_id_uindex
    on irrigation_history (ID);