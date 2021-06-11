truncate items;
truncate carts;
truncate cart_items;
truncate orders;
truncate order_items;

ALTER SEQUENCE items_id_seq RESTART;
ALTER SEQUENCE carts_id_seq RESTART;
ALTER SEQUENCE cart_items_id_seq RESTART;
ALTER SEQUENCE orders_id_seq RESTART;
ALTER SEQUENCE order_items_id_seq RESTART;

INSERT INTO public.items
    (created_at, updated_at, "name", sku, price, qty, deleted_at)
VALUES  
    (now(), now(), 'Product A', 'SKU000001', 1000, 10, null),
    (now(), now(), 'Product B', 'SKU000002', 1000, 10, null),
    (now(), now(), 'Product C', 'SKU000003', 1000, 10, null),
    (now(), now(), 'Product D', 'SKU000004', 1000, 10, null),
    (now(), now(), 'Product E', 'SKU000005', 1000, 10, null);

INSERT INTO public.carts
(created_at, updated_at, username, total_price)
VALUES
    (now(), now(), 'user_1', 8000),
    (now(), now(), 'user_2', 2000),
    (now(), now(), 'user_3', 3000),
    (now(), now(), 'user_4', 5000),
    (now(), now(), 'user_5', 7000);

INSERT INTO public.cart_items
(cart_id, item_id, created_at, updated_at, qty, price)
VALUES
    (1, 1, now(), now(), 5, 1000),
    (1, 2, now(), now(), 3, 1000),
    (2, 2, now(), now(), 1, 1000),
    (2, 5, now(), now(), 1, 1000),
    (3, 1, now(), now(), 3, 1000),
    (4, 1, now(), now(), 4, 1000),
    (4, 4, now(), now(), 1, 1000),
    (5, 1, now(), now(), 4, 1000),
    (4, 2, now(), now(), 1, 1000),
    (4, 3, now(), now(), 1, 1000);
