SELECT
    orders.id AS order_id,
    users.fullname,
    users.email,
    orders.product_name,
    orders.unit_price,
    orders.quantity,
    orders.order_date
FROM users
INNER JOIN orders
ON users.id = orders.user_id
WHERE users.status = 'active' AND
    (orders.unit_price * orders.quantity > 500000 OR
    orders.quantity > 20);