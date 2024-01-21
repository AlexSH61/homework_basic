-- Active: 1705848298161@@127.0.0.1@5432@hw14_sql_db@schema_store
SELECT * FROM schema_store.users;
SELECT * FROM schema_store.products;
SELECT * FROM schema_store.orders WHERE user_id = 1;
SELECT
    u.id AS user_id,
    u.name AS user_name,
    SUM(o.total_amount) AS total_order_amount,
    AVG(p.price) AS average_product_price
FROM schema_store.Users u
LEFT JOIN schema_store.orders o ON u.id = o.user_id
LEFT JOIN schema_store.order_products op ON o.id = op.order_id
LEFT JOIN schema_store.products p ON op.product_id = p.id
WHERE u.id = 1
GROUP BY u.id, u.name;