
CREATE INDEX idx_user_id ON schema_store.Orders(user_id);
CREATE INDEX idx_order_id ON schema_store.order_products(order_id);
CREATE INDEX dx_product_id ON schema_store.order_products(product_id);
CREATE INDEX idx_email ON SCHEMA_store.users(email)