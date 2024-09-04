-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orderproducts_order_id ON orderproducts(order_id);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
