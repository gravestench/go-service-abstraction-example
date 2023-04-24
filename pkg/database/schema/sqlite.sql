CREATE TABLE user_role (
                           `id` INTEGER PRIMARY KEY AUTOINCREMENT,
                           `name` VARCHAR(100) NOT NULL,
                           `description` VARCHAR(200)
);

INSERT INTO
    user_role (name, description)
VALUES
    ('Limited User', 'read-only access'),
    ('Administrator', 'full access');

CREATE TABLE user (
                      `id` INTEGER PRIMARY KEY AUTOINCREMENT,
                      `username` VARCHAR(100) DEFAULT NULL,
                      `password` VARCHAR(100) DEFAULT NULL,
                      `email_address` VARCHAR(100) DEFAULT NULL,
                      `first_name` VARCHAR(100) DEFAULT NULL,
                      `last_name` VARCHAR(100) DEFAULT NULL,
                      `active` TINYINT(1) DEFAULT 1,
                      `last_login` datetime DEFAULT NULL,
                      `role` integer NOT NULL,
                      FOREIGN KEY(role) REFERENCES user_role(id)
);

INSERT INTO
    user (username, password, email_address, first_name, last_name, role, active)
VALUES
    ('admin', 'admin', 'admin@example.com', 'test', 'admin', 2, 1);