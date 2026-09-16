ALTER TABLE menus ADD COLUMN permission_id BIGINT UNSIGNED NULL AFTER icon;

UPDATE menus m
LEFT JOIN permissions p ON p.code = m.permission_code
SET m.permission_id = p.id;

ALTER TABLE menus
    ADD KEY idx_menus_permission_id (permission_id),
    ADD CONSTRAINT fk_menus_permission FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE RESTRICT,
    DROP COLUMN permission_code;
