-- Customers Who Bought All the Products
-- Find customer_ids who bought ALL products from Product table

-- Strategy:
-- 1. Group by customer_id to analyze each customer's purchases
-- 2. Count DISTINCT products each customer bought (handles duplicate rows)
-- 3. Compare to total count of products in Product table
-- 4. Use HAVING to filter customers who bought all products

SELECT customer_id
FROM Customer
GROUP BY customer_id
HAVING COUNT(DISTINCT product_key) = (SELECT COUNT(*) FROM Product);

-- EXPLANATION OF THE SUBQUERY: (SELECT COUNT(*) FROM Product)
-- This subquery runs independently and returns a single number:
-- It counts ALL rows in the Product table
-- In our example:
--   Product table has: | 5 | and | 6 |
--   So COUNT(*) returns 2
-- This gives us the "target number" - how many products a customer needs to buy

-- How it works with the example:
-- Customer 1: bought products {5, 6} → COUNT(DISTINCT) = 2
-- Customer 2: bought products {6} → COUNT(DISTINCT) = 1  
-- Customer 3: bought products {5, 6} → COUNT(DISTINCT) = 2
-- Total products in Product table = 2
-- Only customers 1 and 3 have COUNT(DISTINCT) = 2, so they're returned

-- Key points:
-- - COUNT(DISTINCT product_key) handles duplicate purchases
-- - Subquery (SELECT COUNT(*) FROM Product) gets total product count
-- - HAVING clause filters after grouping (can't use WHERE with aggregates)
-- - No need for JOIN since we only need customer_id and product_key from Customer table 

-- ============================================================================
-- ALTERNATIVE APPROACH: Using JOIN
-- ============================================================================

-- This approach explicitly joins Customer and Product tables
-- Then groups and compares counts
SELECT c.customer_id
FROM Customer c
INNER JOIN Product p ON c.product_key = p.product_key
GROUP BY c.customer_id
HAVING COUNT(DISTINCT c.product_key) = (SELECT COUNT(*) FROM Product);

-- How this works:
-- 1. INNER JOIN ensures we only count purchases of valid products
-- 2. GROUP BY customer_id groups each customer's purchases
-- 3. COUNT(DISTINCT c.product_key) counts unique products each customer bought
-- 4. Still compare to total product count from subquery

-- Both approaches give the same result, but:
-- - First approach is simpler (no JOIN needed)
-- - Second approach is more explicit about the relationship
-- - First approach assumes all product_keys in Customer table are valid
-- - Second approach validates through the JOIN 