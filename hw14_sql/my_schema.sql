create table users (
    id serial primary keys,
    name varchar(255) unique not null,
    email varchar(255) unique not null,
    password varchar(100)
);
create table orders(
    id serial primary key,
    user_id int references users(id),
    order_date Date not null,
    total_amount decimal(10,2) not null
);
create table orderProducts(
id serial primary key,
order_id int references orders(id),
product_id int references products(id),
quantity int,
constraint fk_order foreign key(order_id) references orders(id),
constraint fr_product foreign key(products) references products(id)
);