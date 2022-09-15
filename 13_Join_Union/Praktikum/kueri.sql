Join, Union, Sub-query, Function

--nomer 1
SELECT * FROM transactions WHERE user_id = 1 OR user_id = 2;

--nomer 2
SELECT SUM(d.price * d.qty) FROM transaction t
JOIN transaction_details d ON t.id = d.transaction_id
WHERE t.user_id = 1;

--nomer 3
SELECT SUM(t.qty * t.price) FROM transaction_details t
JOIN products p ON t.product_id = p.id
WHERE p.product_type_id;

--nomer 4
SELECT p.*, pt.name FROM products p
JOIN product_types py ON p.product_type_id = pt.id;

--nomer 5
SELECT t.*, p.name, u.name FROM transactions t
JOIN transaction_details d ON t.id = d transaction_id
JOIN products p ON d.product_id = p.id
JOIN users u ON t.user_id = u.id;

--nomer 6 
CREATE TRIGGER delete_transaction_details
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = OLD.id;


--nomer 7
CREATE TRIGGER update_total_qty
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
    UPDATE transactions
    SET total_qty = total_qty - OLD.qty
    WHERE id = OLD.transaction_id;

--nomer 8
SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transaction_id);