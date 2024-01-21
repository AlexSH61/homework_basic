-- Active: 1705848298161@@127.0.0.1@5432@hw14_sql_db@schema_store
INSERT INTO schema_store.users (name, email, password) VALUES ('sasha','ninja@gmail.com','1234');
UPDATE schema_store.Users SET name = 'pasha' WHERE id = 1;
DELETE FROM schema_store.users WHERE id = 1;
INSERT INTO schema_store.products (name, price) VALUES ('book1', 19.99);
UPDATE schema_store.products SET price = 24.99 WHERE id = 1;
DELETE FROM schema_store.products WHERE id = 1;
INSERT INTO schema_store.orders (user_id, order_date, total_amount) VALUES (1, '2022-01-21', 100.00);
DELETE FROM schema_store.Orders WHERE id = 3;


