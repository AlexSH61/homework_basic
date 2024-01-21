-- Active: 1705848298161@@127.0.0.1@5432@hw14_sql_db@schema_store
create database hw14_sql_db;

 create table schema_store.users(
 id SERIAL primary key,
 name varchar(255) not null,
 email varchar(255) unique not null,
 password varchar(255) not null
 );
CREATE TABLE schema_store.orders (
    id SERIAL PRIMARY KEY,
    user_id int REFERENCES schema_store.users(id),
    order_date DATE,
    total_amount DECIMAL(10, 2),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

create table schema_store.products(
id SERIAL primary key,
name varchar(255) not null,
price decimal(10,2) not null
);
 create table schema_store.order_products(
    order_id INT,
    product_id INT,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES schema_store.orders(id),
    FOREIGN KEY (product_id) REFERENCES schema_store.Products(id)
);