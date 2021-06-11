
CREATE TABLE IF NOT EXISTS "orders"(
    id serial PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    username varchar(255) not null,
    address text,
    total_price decimal(19,2) not null,
    status varchar(25) not null
);

CREATE TABLE IF NOT EXISTS "order_items"(
    id serial PRIMARY KEY,
    order_id integer not null,
    item_id integer not null,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    qty integer not null,
    price decimal(19,2) not null
);