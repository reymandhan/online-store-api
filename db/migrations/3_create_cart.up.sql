
CREATE TABLE IF NOT EXISTS "carts"(
    id serial PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    username varchar(255) not null,
    total_price decimal(19,2) not null
);

CREATE TABLE IF NOT EXISTS "cart_items"(
    id serial PRIMARY KEY,
    cart_id integer not null,
    item_id integer not null,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    qty integer not null
);