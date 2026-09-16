ALTER TABLE menus ADD COLUMN permission_code VARCHAR(120) NULL AFTER icon;

UPDATE menus m
LEFT JOIN permissions p ON p.id = m.permission_id
SET m.permission_code = p.code;

ALTER TABLE menus
    DROP FOREIGN KEY fk_menus_permission,
    DROP INDEX idx_menus_permission_id,
    DROP COLUMN permission_id;
