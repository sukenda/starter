DELETE FROM permissions WHERE code IN ('users.read','users.write','roles.read','roles.write','menus.read','menus.write');
DROP TABLE IF EXISTS menus;
