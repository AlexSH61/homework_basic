INSERT INTO users (name, email, password) VALUES ('petya', 'petya_ninja@example.com', '123456');
INSERT INTO products (name, price) VALUES ('book1', 19.99);
INSERT INTO orders (user_id, order_date, total_amount) VALUES (1, '2024-01-20', 39.98);
INSERT INTO order_products (order_id, product_id, quantity) VALUES (1, 1, 2);
UPDATE users SET name = 'pasha' WHERE id = 1;
UPDATE users SET name = 'sasha' WHERE id = 2;

UPDATE products SET price = 29.99 WHERE id = 1;
UPDATE products SET price = 9.99 WHERE id = 2;
DELETE FROM users WHERE id = 2;

SELECT * FROM users;
SELECT * FROM products;
SELECT * FROM orders WHERE user_id = 1;
SELECT
    u.id AS user_id,
    u.name AS user_name,
    COUNT(o.id) AS total_orders,
    SUM(o.total_amount) AS total_order_amount,
    AVG(p.price) AS average_product_price
FROM users u
JOIN orders o ON u.id = o.user_id
JOIN order_products op ON o.id = op.order_id
JOIN products p ON op.product_id = p.id
WHERE u.id = 1
GROUP BY u.id, u.name;

CREATE INDEX idxorderuser_id ON Orders(user_id);
CREATE INDEX idxorder_date ON orders (order_date);
CREATE INDEX idxorders_user_id_order_date ON Orders (user_id, order_date);
CREATE INDEX idxusers_id ON Users (id);
CREATE INDEX idxproducts_id ON Products (id);



