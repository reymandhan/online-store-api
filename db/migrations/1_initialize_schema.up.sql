  
CREATE TABLE IF NOT EXISTS "items"(
    id serial PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    name varchar(255) not null,
    sku varchar(255) not null,
    price decimal(19,2) not null,
    qty integer not null
);