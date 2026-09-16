CREATE TABLE menus (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    public_id CHAR(36) NOT NULL,
    parent_id BIGINT UNSIGNED NULL,
    code VARCHAR(80) NOT NULL,
    name VARCHAR(120) NOT NULL,
    route VARCHAR(180) NULL,
    icon VARCHAR(80) NULL,
    permission_code VARCHAR(120) NULL,
    sort_order INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (id),
    UNIQUE KEY uq_menus_public_id (public_id),
    UNIQUE KEY uq_menus_code (code),
    KEY idx_menus_parent_order (parent_id, sort_order),
    CONSTRAINT fk_menus_parent FOREIGN KEY (parent_id) REFERENCES menus(id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT IGNORE INTO permissions (code, name, description) VALUES
('users.read', 'View users', 'View user list and details'),
('users.write', 'Manage users', 'Create and update users'),
('roles.read', 'View roles', 'View roles and permissions'),
('roles.write', 'Manage roles', 'Create and update roles and assignments'),
('menus.read', 'View menus', 'View application menus'),
('menus.write', 'Manage menus', 'Create and update application menus');
