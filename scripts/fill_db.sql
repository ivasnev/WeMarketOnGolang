-- Вставка ролей пользователей с использованием генерации случайных данных
INSERT INTO user_roles (name)
SELECT CONCAT('Role_', seq)
FROM generate_series(1, 1000) seq;

-- Вставка статусов заказов
INSERT INTO order_statuses (name)
SELECT CONCAT('Status_', seq)
FROM generate_series(1, 1000) seq;

-- Вставка статусов оплаты
INSERT INTO payment_statuses (name)
SELECT CONCAT('PaymentStatus_', seq)
FROM generate_series(1, 1000) seq;

-- Вставка статусов возвратов
INSERT INTO return_statuses (name)
SELECT CONCAT('ReturnStatus_', seq)
FROM generate_series(1, 1000) seq;

-- Вставка статусов инвентаря
INSERT INTO inventory_statuses (name)
SELECT CONCAT('InventoryStatus_', seq)
FROM generate_series(1, 1000) seq;

-- Вставка статусов акций
INSERT INTO promotion_statuses (name)
SELECT CONCAT('PromotionStatus_', seq)
FROM generate_series(1, 1000) seq;

-- Вставка статусов уведомлений
INSERT INTO notification_statuses (name)
SELECT CONCAT('NotificationStatus_', seq)
FROM generate_series(1, 1000) seq;

-- Вставка производителей с случайными данными
INSERT INTO manufacturers (name, country, website)
SELECT CONCAT('Manufacturer_', seq),
       CASE seq % 3 + 1
           WHEN 1 THEN 'USA'
           WHEN 2 THEN 'Germany'
           WHEN 3 THEN 'China'
           END,
       CONCAT('https://manufacturer', seq, '.com')
FROM generate_series(1, 1000) seq;

-- Вставка пользователей с случайными данными
INSERT INTO users (name, email, password_hash, phone, address, role_id)
SELECT CONCAT('User_', seq),
       CONCAT('user', seq, '@example.com'),
       '$2a$10$L5FsXxDMS2XN0dBLCJwPcObny6akUv6PdnqHvXbABkC7gt7snqmlW', -- aboba1234
       CONCAT('+79', (random() * 1e9)::int),
       CONCAT('Street ', seq),
       (random() * 5 + 1)::int
FROM generate_series(1, 1000) seq;

-- Вставка категорий
INSERT INTO categories (name, description, parent_category_id)
SELECT CONCAT('Category_', seq),
       CONCAT('Description for Category_', seq),
       NULLIF((seq - 1) / 3, 0)::int
FROM generate_series(1, 1000) seq;

-- Вставка продуктов
INSERT INTO products (name, description, price, category_id, stock, sku, weight, availability_status, manufacturer_id)
SELECT CONCAT('Product_', seq),
       CONCAT('Description for Product_', seq),
       (random() * 1000)::numeric(10, 2),
       (random() * 999 + 1)::int,
       (random() * 100)::int,
       CONCAT('SKU_', seq),
       (random() * 10)::numeric(10, 2),
       (random() * 4 + 1)::int,
       (random() * 999 + 1)::int
FROM generate_series(1, 20000) seq;

-- Вставка корзин
INSERT INTO shopping_cart (user_id, status, total_amount)
SELECT (random() * 999 + 1)::int,
       random() > 0.5,
       (random() * 1000)::numeric(10, 2)
FROM generate_series(1, 1000) seq;

-- Вставка позиций в корзину
INSERT INTO cart_items (cart_id, product_id, quantity)
SELECT (random() * 999 + 1)::int,
       (random() * 999 + 1)::int,
       (random() * 5 + 1)::int
FROM generate_series(1, 1000) seq;

-- Вставка заказов
INSERT INTO orders (user_id, total_amount, status_id, delivery_address, payment_method)
SELECT (random() * 999 + 1)::int,
       (random() * 1000)::numeric(10, 2),
       (random() * 4 + 1)::int,
       CONCAT('Delivery Address ', seq),
       CASE seq % 3 + 1
           WHEN 1 THEN 'Credit Card'
           WHEN 2 THEN 'PayPal'
           WHEN 3 THEN 'Bank Transfer'
           END
FROM generate_series(1, 1000) seq;

-- Вставка товаров в заказ
INSERT INTO order_items (order_id, product_id, quantity, price)
SELECT (random() * 999 + 1)::int,
       (random() * 999 + 1)::int,
       (random() * 5 + 1)::int,
       (random() * 1000)::numeric(10, 2)
FROM generate_series(1, 1000) seq;

-- Вставка платежей
INSERT INTO payments (order_id, payment_method, status_id, amount)
SELECT (random() * 999 + 1)::int,
       CASE seq % 3 + 1
           WHEN 1 THEN 'Credit Card'
           WHEN 2 THEN 'PayPal'
           WHEN 3 THEN 'Bank Transfer'
           END,
       (random() * 4 + 1)::int,
       (random() * 1000)::numeric(10, 2)
FROM generate_series(1, 1000) seq;

-- Вставка отзывов
INSERT INTO reviews (user_id, product_id, rating, review_text, status_id)
SELECT (random() * 999 + 1)::int,
       (random() * 999 + 1)::int,
       (random() * 5 + 1)::int,
       CONCAT('Review for Product_', seq),
       (random() * 2 + 1)::int
FROM generate_series(1, 1000) seq;

-- Вставка возвратов
INSERT INTO returns (order_id, product_id, status_id, return_reason)
SELECT (random() * 999 + 1)::int,
       (random() * 999 + 1)::int,
       (random() * 4 + 1)::int,
       CONCAT('Reason for Return_', seq)
FROM generate_series(1, 1000) seq;

-- Вставка акций
INSERT INTO promotions (code, discount, status_id)
SELECT CONCAT('PROMO_', seq),
       (random() * 50)::numeric(5, 2),
       (random() * 3 + 1)::int
FROM generate_series(1, 1000) seq;

-- Вставка подписок на рассылку
INSERT INTO newsletter_subscriptions (email)
SELECT CONCAT('subscriber', seq, '@example.com')
FROM generate_series(1, 1000) seq;

-- Вставка системных логов
INSERT INTO system_logs (level, message, user_id)
SELECT CASE seq % 3 + 1
           WHEN 1 THEN 'INFO'
           WHEN 2 THEN 'ERROR'
           WHEN 3 THEN 'DEBUG'
           END,
       CONCAT('Message_', seq),
       NULLIF((random() * 999 + 1)::int, 5)
FROM generate_series(1, 1000) seq;
