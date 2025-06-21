-- Triangle Judgement

-- Table: Triangle
-- +-------------+------+
-- | Column Name | Type |
-- +-------------+------+
-- | x           | int  |
-- | y           | int  |
-- | z           | int  |
-- +-------------+------+
-- In SQL, (x, y, z) is the primary key column for this table.
-- Each row of this table contains the lengths of three line segments.

-- Report for every three line segments whether they can form a triangle.
-- Return the result table in any order.

-- Example 1:
-- Input: 
-- Triangle table:
-- +----+----+----+
-- | x  | y  | z  |
-- +----+----+----+
-- | 13 | 15 | 30 |
-- | 10 | 20 | 15 |
-- +----+----+----+
-- Output: 
-- +----+----+----+----------+
-- | x  | y  | z  | triangle |
-- +----+----+----+----------+
-- | 13 | 15 | 30 | No       |
-- | 10 | 20 | 15 | Yes      |
-- +----+----+----+----------+

-- ============================================================================
-- TRIANGLE INEQUALITY THEOREM
-- ============================================================================

-- For three line segments with lengths a, b, c to form a triangle:
-- The sum of any two sides must be GREATER than the third side
--
-- This gives us THREE conditions that ALL must be true:
-- 1. a + b > c
-- 2. a + c > b  
-- 3. b + c > a
--
-- If ANY of these conditions is false, the segments cannot form a triangle.

-- ============================================================================
-- STEP-BY-STEP ANALYSIS
-- ============================================================================

-- Example 1: x=13, y=15, z=30
-- Check condition 1: x + y > z → 13 + 15 > 30 → 28 > 30 → FALSE ❌
-- Since one condition fails, result is "No"

-- Example 2: x=10, y=20, z=15  
-- Check condition 1: x + y > z → 10 + 20 > 15 → 30 > 15 → TRUE ✅
-- Check condition 2: x + z > y → 10 + 15 > 20 → 25 > 20 → TRUE ✅
-- Check condition 3: y + z > x → 20 + 15 > 10 → 35 > 10 → TRUE ✅
-- All conditions pass, result is "Yes"

-- ============================================================================
-- SQL SOLUTION
-- ============================================================================

SELECT 
    x, 
    y, 
    z,
    CASE 
        WHEN x + y > z AND x + z > y AND y + z > x THEN 'Yes'
        ELSE 'No'
    END AS triangle
FROM Triangle;

-- ============================================================================
-- ALTERNATIVE APPROACHES
-- ============================================================================

-- Approach 1: Using IF statement (MySQL specific)
SELECT 
    x, 
    y, 
    z,
    IF(x + y > z AND x + z > y AND y + z > x, 'Yes', 'No') AS triangle
FROM Triangle;

-- Approach 2: More explicit CASE statement
SELECT 
    x, 
    y, 
    z,
    CASE 
        WHEN x + y <= z THEN 'No'
        WHEN x + z <= y THEN 'No' 
        WHEN y + z <= x THEN 'No'
        ELSE 'Yes'
    END AS triangle
FROM Triangle;

-- ============================================================================
-- EXPLANATION OF THE LOGIC
-- ============================================================================

-- The CASE statement evaluates:
-- 1. x + y > z: Can sides x and y together "reach" beyond side z?
-- 2. x + z > y: Can sides x and z together "reach" beyond side y?
-- 3. y + z > x: Can sides y and z together "reach" beyond side x?
--
-- ALL three conditions must be true using AND operator
-- If any condition fails, the triangle cannot be formed

-- ============================================================================
-- EDGE CASES TO CONSIDER
-- ============================================================================

-- 1. Degenerate triangle: x=5, y=5, z=10
--    5 + 5 = 10, which is NOT > 10, so "No"
--
-- 2. Equilateral triangle: x=5, y=5, z=5
--    All conditions: 5 + 5 > 5 → 10 > 5 → TRUE, so "Yes"
--
-- 3. Right triangle: x=3, y=4, z=5
--    3 + 4 > 5 → 7 > 5 → TRUE
--    3 + 5 > 4 → 8 > 4 → TRUE  
--    4 + 5 > 3 → 9 > 3 → TRUE, so "Yes"

-- ============================================================================
-- FINAL SOLUTION
-- ============================================================================

SELECT 
    x, 
    y, 
    z,
    CASE 
        WHEN x + y > z AND x + z > y AND y + z > x THEN 'Yes'
        ELSE 'No'
    END AS triangle
FROM Triangle; 