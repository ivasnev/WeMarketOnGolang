-- Таблица ролей пользователей
CREATE TABLE user_roles
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Описание для user_roles
COMMENT
ON TABLE user_roles IS 'Таблица для хранения ролей пользователей';
COMMENT
ON COLUMN user_roles.id IS 'ID роли';
COMMENT
ON COLUMN user_roles.name IS 'Название роли';

-- Таблица статусов заказов
CREATE TABLE order_statuses
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Описание для order_statuses
COMMENT
ON TABLE order_statuses IS 'Таблица для хранения статусов заказов';
COMMENT
ON COLUMN order_statuses.id IS 'ID статуса';
COMMENT
ON COLUMN order_statuses.name IS 'Название статуса';

-- Таблица статусов платежей
CREATE TABLE payment_statuses
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Описание для payment_statuses
COMMENT
ON TABLE payment_statuses IS 'Таблица для хранения статусов платежей';
COMMENT
ON COLUMN payment_statuses.id IS 'ID статуса';
COMMENT
ON COLUMN payment_statuses.name IS 'Название статуса';

-- Таблица статусов возвратов
CREATE TABLE return_statuses
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Описание для return_statuses
COMMENT
ON TABLE return_statuses IS 'Таблица для хранения статусов возвратов';
COMMENT
ON COLUMN return_statuses.id IS 'ID статуса';
COMMENT
ON COLUMN return_statuses.name IS 'Название статуса';

-- Таблица статусов запасов
CREATE TABLE inventory_statuses
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Описание для inventory_statuses
COMMENT
ON TABLE inventory_statuses IS 'Таблица для хранения статусов запасов';
COMMENT
ON COLUMN inventory_statuses.id IS 'ID статуса';
COMMENT
ON COLUMN inventory_statuses.name IS 'Название статуса';

-- Таблица статусов промокодов
CREATE TABLE promotion_statuses
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Описание для promotion_statuses
COMMENT
ON TABLE promotion_statuses IS 'Таблица для хранения статусов промокодов';
COMMENT
ON COLUMN promotion_statuses.id IS 'ID статуса';
COMMENT
ON COLUMN promotion_statuses.name IS 'Название статуса';

-- Таблица статусов уведомлений
CREATE TABLE notification_statuses
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Описание для notification_statuses
COMMENT
ON TABLE notification_statuses IS 'Таблица для хранения статусов уведомлений';
COMMENT
ON COLUMN notification_statuses.id IS 'ID статуса';
COMMENT
ON COLUMN notification_statuses.name IS 'Название статуса';

-- Таблица пользователей
CREATE TABLE users
(
    id                SERIAL PRIMARY KEY,
    name              VARCHAR(100)        NOT NULL,
    email             VARCHAR(100) UNIQUE NOT NULL,
    password_hash     VARCHAR(255)        NOT NULL,
    phone             VARCHAR(20),
    address           VARCHAR(255),
    role_id           INTEGER             NOT NULL,
    registration_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login        TIMESTAMP,
    account_status    BOOLEAN   DEFAULT TRUE,
    order_count       INTEGER   DEFAULT 0,
    FOREIGN KEY (role_id) REFERENCES user_roles (id)
);

-- Описание для users
COMMENT
ON TABLE users IS 'Таблица для хранения информации о пользователях';
COMMENT
ON COLUMN users.id IS 'ID пользователя';
COMMENT
ON COLUMN users.name IS 'Имя пользователя';
COMMENT
ON COLUMN users.email IS 'Электронная почта пользователя';
COMMENT
ON COLUMN users.password_hash IS 'Хэш пароля пользователя';
COMMENT
ON COLUMN users.phone IS 'Телефон пользователя';
COMMENT
ON COLUMN users.address IS 'Адрес пользователя';
COMMENT
ON COLUMN users.role_id IS 'ID роли пользователя';
COMMENT
ON COLUMN users.registration_date IS 'Дата регистрации пользователя';
COMMENT
ON COLUMN users.last_login IS 'Дата последнего входа пользователя';
COMMENT
ON COLUMN users.account_status IS 'Статус аккаунта (активен/заблокирован)';
COMMENT
ON COLUMN users.order_count IS 'Количество заказов пользователя';

-- Таблица продуктов
CREATE TABLE products
(
    id                  SERIAL PRIMARY KEY,
    name                VARCHAR(255)   NOT NULL,
    description         TEXT,
    price               NUMERIC(10, 2) NOT NULL,
    category_id         INTEGER        NOT NULL,
    stock               INTEGER        NOT NULL,
    image_url           VARCHAR(255),
    added_date          TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    options             JSONB,
    manufacturer        VARCHAR(100),
    sku                 VARCHAR(100),
    weight              NUMERIC(10, 2),
    dimensions          JSONB,
    availability_status INTEGER        NOT NULL,
    manufacturer_id     INT,
    FOREIGN KEY (manufacturer_id) REFERENCES manufacturers (manufacturer_id),
    FOREIGN KEY (category_id) REFERENCES categories (id),
    FOREIGN KEY (availability_status) REFERENCES inventory_statuses (id)
);

-- Описание для products
COMMENT
ON TABLE products IS 'Таблица для хранения информации о продуктах';
COMMENT
ON COLUMN products.id IS 'ID продукта';
COMMENT
ON COLUMN products.name IS 'Название продукта';
COMMENT
ON COLUMN products.description IS 'Описание продукта';
COMMENT
ON COLUMN products.price IS 'Цена продукта';
COMMENT
ON COLUMN products.category_id IS 'ID категории продукта';
COMMENT
ON COLUMN products.stock IS 'Количество на складе';
COMMENT
ON COLUMN products.image_url IS 'URL изображения продукта';
COMMENT
ON COLUMN products.added_date IS 'Дата добавления продукта';
COMMENT
ON COLUMN products.options IS 'Варианты продукта (цвет, размер и т.д.)';
COMMENT
ON COLUMN products.manufacturer_id IS 'Производитель продукта';
COMMENT
ON COLUMN products.sku IS 'Артикул продукта';
COMMENT
ON COLUMN products.weight IS 'Вес продукта';
COMMENT
ON COLUMN products.dimensions IS 'Габариты продукта (длина, ширина, высота)';
COMMENT
ON COLUMN products.availability_status IS 'ID статуса наличия продукта';

-- Таблица категорий продуктов
CREATE TABLE categories
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(100) NOT NULL,
    description        TEXT,
    parent_category_id INTEGER,
    FOREIGN KEY (parent_category_id) REFERENCES categories (id)
);

-- Описание для categories
COMMENT
ON TABLE categories IS 'Таблица для хранения категорий продуктов';
COMMENT
ON COLUMN categories.id IS 'ID категории';
COMMENT
ON COLUMN categories.name IS 'Название категории';
COMMENT
ON COLUMN categories.description IS 'Описание категории';
COMMENT
ON COLUMN categories.parent_category_id IS 'ID родительской категории (для вложенных категорий)';

-- Таблица корзины покупок
CREATE TABLE shopping_cart
(
    id           SERIAL PRIMARY KEY,
    user_id      INTEGER NOT NULL,
    status       BOOLEAN   DEFAULT TRUE,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    promo_code   VARCHAR(50),
    total_amount NUMERIC(10, 2),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Описание для shopping_cart
COMMENT
ON TABLE shopping_cart IS 'Таблица для хранения информации о корзинах покупок';
COMMENT
ON COLUMN shopping_cart.id IS 'ID корзины';
COMMENT
ON COLUMN shopping_cart.user_id IS 'ID пользователя, которому принадлежит корзина';
COMMENT
ON COLUMN shopping_cart.status IS 'Статус корзины (активна/неактивна)';
COMMENT
ON COLUMN shopping_cart.created_date IS 'Дата создания корзины';
COMMENT
ON COLUMN shopping_cart.promo_code IS 'Промокод, если применён';
COMMENT
ON COLUMN shopping_cart.total_amount IS 'Итоговая сумма корзины';

-- Таблица товаров в корзине
CREATE TABLE cart_items
(
    id            SERIAL PRIMARY KEY,
    cart_id       INTEGER        NOT NULL,
    product_id    INTEGER        NOT NULL,
    quantity      INTEGER        NOT NULL,
    FOREIGN KEY (cart_id) REFERENCES shopping_cart (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);

-- Описание для cart_items
COMMENT
ON TABLE cart_items IS 'Таблица для хранения товаров в корзине покупок';
COMMENT
ON COLUMN cart_items.id IS 'ID записи';
COMMENT
ON COLUMN cart_items.cart_id IS 'ID корзины';
COMMENT
ON COLUMN cart_items.product_id IS 'ID продукта';
COMMENT
ON COLUMN cart_items.quantity IS 'Количество товара';
COMMENT

-- Таблица заказов
CREATE TABLE orders
(
    id               SERIAL PRIMARY KEY,
    user_id          INTEGER        NOT NULL,
    order_date       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount     NUMERIC(10, 2) NOT NULL,
    status_id        INTEGER        NOT NULL,
    delivery_address VARCHAR(255),
    payment_method   VARCHAR(50),
    tracking_number  VARCHAR(100),
    fulfillment_time TIMESTAMP,
    shipping_method  VARCHAR(50),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (status_id) REFERENCES order_statuses (id)
);

-- Описание для orders
COMMENT
ON TABLE orders IS 'Таблица для хранения информации о заказах';
COMMENT
ON COLUMN orders.id IS 'ID заказа';
COMMENT
ON COLUMN orders.user_id IS 'ID пользователя, который сделал заказ';
COMMENT
ON COLUMN orders.order_date IS 'Дата заказа';
COMMENT
ON COLUMN orders.total_amount IS 'Итоговая сумма заказа';
COMMENT
ON COLUMN orders.status_id IS 'ID статуса заказа';
COMMENT
ON COLUMN orders.delivery_address IS 'Адрес доставки';
COMMENT
ON COLUMN orders.payment_method IS 'Метод оплаты';
COMMENT
ON COLUMN orders.tracking_number IS 'Номер отслеживания заказа';
COMMENT
ON COLUMN orders.fulfillment_time IS 'Время выполнения заказа';
COMMENT
ON COLUMN orders.shipping_method IS 'Метод доставки';

-- Таблица товаров в заказе
CREATE TABLE order_items
(
    id         SERIAL PRIMARY KEY,
    order_id   INTEGER        NOT NULL,
    product_id INTEGER        NOT NULL,
    quantity   INTEGER        NOT NULL,
    price      NUMERIC(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);

-- Описание для order_items
COMMENT
ON TABLE order_items IS 'Таблица для хранения товаров в заказе';
COMMENT
ON COLUMN order_items.id IS 'ID записи';
COMMENT
ON COLUMN order_items.order_id IS 'ID заказа';
COMMENT
ON COLUMN order_items.product_id IS 'ID продукта';
COMMENT
ON COLUMN order_items.quantity IS 'Количество товара';
COMMENT
ON COLUMN order_items.price IS 'Цена за единицу товара';

-- Таблица платежей
CREATE TABLE payments
(
    id             SERIAL PRIMARY KEY,
    order_id       INTEGER        NOT NULL,
    payment_method VARCHAR(50)    NOT NULL,
    status_id      INTEGER        NOT NULL,
    payment_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    amount         NUMERIC(10, 2) NOT NULL,
    transaction_id VARCHAR(100),
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (status_id) REFERENCES payment_statuses (id)
);

-- Описание для payments
COMMENT
ON TABLE payments IS 'Таблица для хранения информации о платежах';
COMMENT
ON COLUMN payments.id IS 'ID платежа';
COMMENT
ON COLUMN payments.order_id IS 'ID заказа';
COMMENT
ON COLUMN payments.payment_method IS 'Метод оплаты';
COMMENT
ON COLUMN payments.status_id IS 'ID статуса платежа';
COMMENT
ON COLUMN payments.payment_date IS 'Дата платежа';
COMMENT
ON COLUMN payments.amount IS 'Сумма платежа';
COMMENT
ON COLUMN payments.transaction_id IS 'Транзакционный идентификатор';

-- Таблица отзывов
CREATE TABLE reviews
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER NOT NULL,
    product_id  INTEGER NOT NULL,
    rating      INTEGER NOT NULL,
    review_text TEXT,
    review_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status_id   INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (status_id) REFERENCES notification_statuses (id)
);

-- Описание для reviews
COMMENT
ON TABLE reviews IS 'Таблица для хранения отзывов';
COMMENT
ON COLUMN reviews.id IS 'ID отзыва';
COMMENT
ON COLUMN reviews.user_id IS 'ID пользователя, который оставил отзыв';
COMMENT
ON COLUMN reviews.product_id IS 'ID продукта, на который оставлен отзыв';
COMMENT
ON COLUMN reviews.rating IS 'Оценка отзыва (в звёздах)';
COMMENT
ON COLUMN reviews.review_text IS 'Текст отзыва';
COMMENT
ON COLUMN reviews.review_date IS 'Дата добавления отзыва';
COMMENT
ON COLUMN reviews.status_id IS 'ID статуса отзыва';

-- Таблица возвратов
CREATE TABLE returns
(
    id            SERIAL PRIMARY KEY,
    order_id      INTEGER NOT NULL,
    product_id    INTEGER NOT NULL,
    return_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status_id     INTEGER NOT NULL,
    return_reason TEXT,
    FOREIGN KEY (order_id) REFERENCES orders (id),
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (status_id) REFERENCES return_statuses (id)
);

-- Описание для returns
COMMENT
ON TABLE returns IS 'Таблица для хранения возвратов';
COMMENT
ON COLUMN returns.id IS 'ID возврата';
COMMENT
ON COLUMN returns.order_id IS 'ID заказа';
COMMENT
ON COLUMN returns.product_id IS 'ID продукта';
COMMENT
ON COLUMN returns.return_date IS 'Дата возврата';
COMMENT
ON COLUMN returns.status_id IS 'ID статуса возврата';
COMMENT
ON COLUMN returns.return_reason IS 'Причина возврата';

-- Таблица промокодов
CREATE TABLE promotions
(
    id                   SERIAL PRIMARY KEY,
    code                 VARCHAR(50)   NOT NULL,
    discount             NUMERIC(5, 2) NOT NULL,
    minimum_order_amount NUMERIC(10, 2),
    start_date           TIMESTAMP,
    end_date             TIMESTAMP,
    status_id            INTEGER       NOT NULL,
    usage_limit          INTEGER,
    FOREIGN KEY (status_id) REFERENCES promotion_statuses (id)
);

-- Описание для promotions
COMMENT
ON TABLE promotions IS 'Таблица для хранения промокодов';
COMMENT
ON COLUMN promotions.id IS 'ID промокода';
COMMENT
ON COLUMN promotions.code IS 'Код промокода';
COMMENT
ON COLUMN promotions.discount IS 'Скидка, предоставляемая промокодом';
COMMENT
ON COLUMN promotions.minimum_order_amount IS 'Минимальная сумма заказа для использования промокода';
COMMENT
ON COLUMN promotions.start_date IS 'Дата начала действия промокода';
COMMENT
ON COLUMN promotions.end_date IS 'Дата окончания действия промокода';
COMMENT
ON COLUMN promotions.status_id IS 'ID статуса промокода';
COMMENT
ON COLUMN promotions.usage_limit IS 'Лимит использования промокода';

-- Таблица сессий пользователей
CREATE TABLE user_sessions
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER NOT NULL,
    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_date   TIMESTAMP,
    ip_address VARCHAR(45),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Описание для user_sessions
COMMENT
ON TABLE user_sessions IS 'Таблица для хранения сессий пользователей';
COMMENT
ON COLUMN user_sessions.id IS 'ID сессии';
COMMENT
ON COLUMN user_sessions.user_id IS 'ID пользователя';
COMMENT
ON COLUMN user_sessions.start_date IS 'Дата начала сессии';
COMMENT
ON COLUMN user_sessions.end_date IS 'Дата окончания сессии';
COMMENT
ON COLUMN user_sessions.ip_address IS 'IP-адрес пользователя';

-- Таблица уведомлений
CREATE TABLE notifications
(
    id                SERIAL PRIMARY KEY,
    user_id           INTEGER     NOT NULL,
    notification_type VARCHAR(50) NOT NULL,
    message           TEXT,
    sent_date         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    read_status       BOOLEAN   DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Описание для notifications
COMMENT
ON TABLE notifications IS 'Таблица для хранения уведомлений';
COMMENT
ON COLUMN notifications.id IS 'ID уведомления';
COMMENT
ON COLUMN notifications.user_id IS 'ID пользователя, которому отправлено уведомление';
COMMENT
ON COLUMN notifications.notification_type IS 'Тип уведомления (заказ, акция, системное)';
COMMENT
ON COLUMN notifications.message IS 'Текст уведомления';
COMMENT
ON COLUMN notifications.sent_date IS 'Дата отправки уведомления';
COMMENT
ON COLUMN notifications.read_status IS 'Статус прочтения уведомления';

-- Таблица акций
CREATE TABLE sales
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100)  NOT NULL,
    description TEXT,
    discount    NUMERIC(5, 2) NOT NULL,
    start_date  TIMESTAMP,
    end_date    TIMESTAMP,
    products    JSONB
);

-- Описание для sales
COMMENT
ON TABLE sales IS 'Таблица для хранения акций';
COMMENT
ON COLUMN sales.id IS 'ID акции';
COMMENT
ON COLUMN sales.name IS 'Название акции';
COMMENT
ON COLUMN sales.description IS 'Описание акции';
COMMENT
ON COLUMN sales.discount IS 'Скидка по акции';
COMMENT
ON COLUMN sales.start_date IS 'Дата начала акции';
COMMENT
ON COLUMN sales.end_date IS 'Дата окончания акции';
COMMENT
ON COLUMN sales.products IS 'Список ID продуктов, участвующих в акции';

-- Таблица поддержки клиентов
CREATE TABLE customer_support
(
    id           SERIAL PRIMARY KEY,
    user_id      INTEGER NOT NULL,
    subject      VARCHAR(255),
    message      TEXT,
    status_id    INTEGER NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    response     TEXT,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (status_id) REFERENCES notification_statuses (id)
);

-- Описание для customer_support
COMMENT
ON TABLE customer_support IS 'Таблица для хранения обращений в поддержку клиентов';
COMMENT
ON COLUMN customer_support.id IS 'ID обращения';
COMMENT
ON COLUMN customer_support.user_id IS 'ID пользователя, сделавшего обращение';
COMMENT
ON COLUMN customer_support.subject IS 'Тема обращения';
COMMENT
ON COLUMN customer_support.message IS 'Текст обращения';
COMMENT
ON COLUMN customer_support.status_id IS 'ID статуса обращения';
COMMENT
ON COLUMN customer_support.created_date IS 'Дата создания обращения';
COMMENT
ON COLUMN customer_support.response IS 'Ответ сотрудника поддержки';

-- Таблица истории изменений продуктов
CREATE TABLE product_changes
(
    id            SERIAL PRIMARY KEY,
    product_id    INTEGER     NOT NULL,
    field_changed VARCHAR(50) NOT NULL,
    old_value     TEXT,
    new_value     TEXT,
    change_date   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    changed_by    INTEGER,
    FOREIGN KEY (product_id) REFERENCES products (id),
    FOREIGN KEY (changed_by) REFERENCES users (id)
);

-- Описание для product_changes
COMMENT
ON TABLE product_changes IS 'Таблица для хранения истории изменений продуктов';
COMMENT
ON COLUMN product_changes.id IS 'ID изменения';
COMMENT
ON COLUMN product_changes.product_id IS 'ID продукта';
COMMENT
ON COLUMN product_changes.field_changed IS 'Поле, которое изменилось';
COMMENT
ON COLUMN product_changes.old_value IS 'Старое значение';
COMMENT
ON COLUMN product_changes.new_value IS 'Новое значение';
COMMENT
ON COLUMN product_changes.change_date IS 'Дата изменения';
COMMENT
ON COLUMN product_changes.changed_by IS 'ID пользователя, который внёс изменение';

-- Таблица подписки на новости
CREATE TABLE newsletter_subscriptions
(
    id                SERIAL PRIMARY KEY,
    email             VARCHAR(100) NOT NULL,
    subscription_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status            BOOLEAN   DEFAULT TRUE
);

-- Описание для newsletter_subscriptions
COMMENT
ON TABLE newsletter_subscriptions IS 'Таблица для хранения подписок на новости';
COMMENT
ON COLUMN newsletter_subscriptions.id IS 'ID подписки';
COMMENT
ON COLUMN newsletter_subscriptions.email IS 'Электронная почта';
COMMENT
ON COLUMN newsletter_subscriptions.subscription_date IS 'Дата подписки';
COMMENT
ON COLUMN newsletter_subscriptions.status IS 'Статус подписки (активна/отписался)';

-- Таблица логов системы
CREATE TABLE system_logs
(
    id       SERIAL PRIMARY KEY,
    log_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    level    VARCHAR(50) NOT NULL,
    message  TEXT,
    user_id  INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- Описание для system_logs
COMMENT
ON TABLE system_logs IS 'Таблица для хранения логов системы';
COMMENT
ON COLUMN system_logs.id IS 'ID лога';
COMMENT
ON COLUMN system_logs.log_date IS 'Дата и время записи лога';
COMMENT
ON COLUMN system_logs.level IS 'Уровень лога (информация, предупреждение, ошибка)';
COMMENT
ON COLUMN system_logs.message IS 'Текст сообщения лога';
COMMENT
ON COLUMN system_logs.user_id IS 'ID пользователя, связанного с логом';

CREATE TABLE manufacturers
(
    manufacturer_id SERIAL PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    country         VARCHAR(100),
    website         VARCHAR(255),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
