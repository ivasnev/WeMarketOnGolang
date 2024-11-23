-- Вставка ролей пользователей
INSERT INTO user_roles (name) VALUES
                                  ('Admin'),
                                  ('Customer'),
                                  ('Manager');

-- Вставка статусов заказов
INSERT INTO order_statuses (name) VALUES
                                      ('Pending'),
                                      ('Shipped'),
                                      ('Delivered'),
                                      ('Cancelled');

-- Вставка статусов оплаты
INSERT INTO payment_statuses (name) VALUES
                                        ('Pending'),
                                        ('Completed'),
                                        ('Failed'),
                                        ('Refunded');

-- Вставка статусов возвратов
INSERT INTO return_statuses (name) VALUES
                                       ('Requested'),
                                       ('Approved'),
                                       ('Rejected'),
                                       ('Completed');

-- Вставка статусов инвентаря
INSERT INTO inventory_statuses (name) VALUES
                                          ('In Stock'),
                                          ('Out of Stock'),
                                          ('Pre-order'),
                                          ('Discontinued');

-- Вставка статусов акций
INSERT INTO promotion_statuses (name) VALUES
                                          ('Active'),
                                          ('Expired'),
                                          ('Upcoming');

-- Вставка статусов уведомлений
INSERT INTO notification_statuses (name) VALUES
                                             ('Unread'),
                                             ('Read');

-- Вставка производителей
INSERT INTO manufacturers (name, country, website) VALUES
                                                       ('Manufacturer1', 'USA', 'https://manufacturer1.com'),
                                                       ('Manufacturer2', 'Germany', 'https://manufacturer2.com'),
                                                       ('Manufacturer3', 'China', 'https://manufacturer3.com');

-- Вставка пользователей
INSERT INTO users (name, email, password_hash, phone, address, role_id) VALUES
                                                                            ('John Doe', 'john.doe@example.com', 'hashed_password_1', '1234567890', '123 Main St', 1),
                                                                            ('Jane Smith', 'jane.smith@example.com', 'hashed_password_2', '9876543210', '456 Oak St', 2);

-- Вставка категорий
INSERT INTO categories (name, description, parent_category_id) VALUES
                                                                   ('Electronics', 'Electronic devices', NULL),
                                                                   ('Home Appliances', 'Appliances for home', NULL),
                                                                   ('Laptops', 'Portable computers', 1);

-- Вставка продуктов
INSERT INTO products (name, description, price, category_id, stock, sku, weight, availability_status, manufacturer_id) VALUES
                                                                                                                                         ('Laptop 1', 'High-performance laptop', 999.99, 3, 100, 'SKU123', 2.5, 1, 1),
                                                                                                                                         ('Washing Machine', 'Efficient washing machine', 499.99, 2, 50, 'SKU124', 80.0, 1, 2);

-- Вставка товаров в корзину
INSERT INTO shopping_cart (user_id, status, total_amount) VALUES
                                                              (1, TRUE, 1499.98),
                                                              (2, TRUE, 499.99);

-- Вставка позиций в корзину
INSERT INTO cart_items (cart_id, product_id, quantity) VALUES
                                                           (1, 1, 1),
                                                           (2, 2, 1);

-- Вставка заказов
INSERT INTO orders (user_id, total_amount, status_id, delivery_address, payment_method) VALUES
                                                                                            (1, 1499.98, 1, '123 Main St', 'Credit Card'),
                                                                                            (2, 499.99, 1, '456 Oak St', 'PayPal');

-- Вставка товаров в заказ
INSERT INTO order_items (order_id, product_id, quantity, price) VALUES
                                                                    (1, 1, 1, 999.99),
                                                                    (2, 2, 1, 499.99);

-- Вставка платежей
INSERT INTO payments (order_id, payment_method, status_id, amount) VALUES
                                                                       (1, 'Credit Card', 2, 1499.98),
                                                                       (2, 'PayPal', 2, 499.99);

-- Вставка отзывов
INSERT INTO reviews (user_id, product_id, rating, review_text, status_id) VALUES
                                                                              (1, 1, 5, 'Great laptop!', 1),
                                                                              (2, 2, 4, 'Good washing machine.', 1);

-- Вставка возвратов
INSERT INTO returns (order_id, product_id, status_id, return_reason) VALUES
    (1, 1, 3, 'Damaged item');

-- Вставка акций
INSERT INTO promotions (code, discount, status_id) VALUES
    ('PROMO2023', 10.0, 1);

-- Вставка подписок на рассылку
INSERT INTO newsletter_subscriptions (email) VALUES
                                                 ('subscriber1@example.com'),
                                                 ('subscriber2@example.com');

-- Вставка системных логов
INSERT INTO system_logs (level, message, user_id) VALUES
                                                      ('INFO', 'System started', NULL),
                                                      ('ERROR', 'Payment gateway failed', 1);
