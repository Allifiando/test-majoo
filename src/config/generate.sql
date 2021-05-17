CREATE TABLE users (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(60) NOT NULL,
    is_active bool DEFAULT 1,
    role_id int NOT NULL
)
ALTER TABLE roles ADD CONSTRAINT users_fk FOREIGN KEY (role_id) REFERENCES roles(id);

CREATE TABLE roles (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name CHAR(25) NOT NULL,    
    is_active bool DEFAULT 1    
)

CREATE TABLE merchants (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id int NOT NULL,
    company_name VARCHAR(200) NOT NULL,    
    trademark VARCHAR(200) NOT NULL,    
    is_active bool DEFAULT 1    
)
ALTER TABLE merchants ADD CONSTRAINT merchants_fk FOREIGN KEY (user_id) REFERENCES users(id);

CREATE TABLE outlets (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    marchant_id int NOT NULL,
    name VARCHAR(200) NOT NULL,    
    location VARCHAR(200) NOT NULL,    
    is_active bool DEFAULT 1    
)
ALTER TABLE outlets ADD CONSTRAINT outlets_fk FOREIGN KEY (merchant_id) REFERENCES marchants(id);

CREATE TABLE products (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    outlet_id int NOT NULL,
    name VARCHAR(200) NOT NULL,    
    price int NOT NULL,    
    qty VARCHAR(10) NOT NULL,    
    filename VARCHAR(100) NOT NULL,    
    is_active bool DEFAULT 1    
)
ALTER TABLE products ADD CONSTRAINT products_fk FOREIGN KEY (outlet_id) REFERENCES outlets(id);

