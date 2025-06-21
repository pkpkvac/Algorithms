-- Consecutive Numbers

-- Table: Logs
-- +-------------+---------+
-- | Column Name | Type    |
-- +-------------+---------+
-- | id          | int     |
-- | num         | varchar |
-- +-------------+---------+
-- In SQL, id is the primary key for this table.
-- id is an autoincrement column starting from 1.

-- Find all numbers that appear at least three times consecutively.
-- Return the result table in any order.

-- Example 1:
-- Input: 
-- Logs table:
-- +----+-----+
-- | id | num |
-- +----+-----+
-- | 1  | 1   |
-- | 2  | 1   |
-- | 3  | 1   |
-- | 4  | 2   |
-- | 5  | 1   |
-- | 6  | 2   |
-- | 7  | 2   |
-- +----+-----+
-- Output: 
-- +-----------------+
-- | ConsecutiveNums |
-- +-----------------+
-- | 1               |
-- +-----------------+
-- Explanation: 1 is the only number that appears consecutively for at least three times.

-- ============================================================================
-- PROBLEM ANALYSIS
-- ============================================================================

-- Your attempt had these issues:
-- 1. COUNT(l.num >= 3) - This doesn't work in SQL. COUNT counts rows, not conditions.
-- 2. Missing the "consecutive" logic - need to compare adjacent rows
-- 3. Need to check if the SAME number appears in consecutive rows

-- The key insight: We need to find rows where:
-- - Current row has the same num as the next row
-- - AND the next row has the same num as the row after it
-- - This gives us 3 consecutive identical numbers

-- ============================================================================
-- APPROACH 1: SELF-JOIN (Most Intuitive)
-- ============================================================================

-- Join the table with itself three times to get three consecutive rows
SELECT DISTINCT l1.num AS ConsecutiveNums
FROM Logs l1
JOIN Logs l2 ON l1.id = l2.id - 1  -- l2 is the next row after l1
JOIN Logs l3 ON l2.id = l3.id - 1  -- l3 is the next row after l2
WHERE l1.num = l2.num AND l2.num = l3.num;

-- ============================================================================
-- STEP-BY-STEP TRACE OF SELF-JOIN
-- ============================================================================

-- For our example data:
-- +----+-----+
-- | id | num |
-- +----+-----+
-- | 1  | 1   |
-- | 2  | 1   |
-- | 3  | 1   |
-- | 4  | 2   |
-- | 5  | 1   |
-- | 6  | 2   |
-- | 7  | 2   |
-- +----+-----+

-- The joins will create combinations like:
-- l1(id=1,num=1) + l2(id=2,num=1) + l3(id=3,num=1) → l1.num=l2.num=l3.num=1 ✅
-- l1(id=2,num=1) + l2(id=3,num=1) + l3(id=4,num=2) → l1.num=l2.num≠l3.num ❌
-- l1(id=3,num=1) + l2(id=4,num=2) + l3(id=5,num=1) → l1.num≠l2.num ❌
-- l1(id=4,num=2) + l2(id=5,num=1) + l3(id=6,num=2) → l1.num≠l2.num ❌
-- l1(id=5,num=1) + l2(id=6,num=2) + l3(id=7,num=2) → l1.num≠l2.num ❌
-- l1(id=6,num=2) + l2(id=7,num=2) + l3(id=8,num=?) → No l3 exists ❌

-- Only the first combination satisfies all conditions, so result is "1"

-- ============================================================================
-- APPROACH 2: WINDOW FUNCTIONS (More Advanced)
-- ============================================================================

-- Use LAG() to get previous values and compare
SELECT DISTINCT num AS ConsecutiveNums
FROM (
    SELECT 
        num,
        LAG(num, 1) OVER (ORDER BY id) AS prev_num,
        LAG(num, 2) OVER (ORDER BY id) AS prev_prev_num
    FROM Logs
) t
WHERE num = prev_num AND num = prev_prev_num;

-- ============================================================================
-- APPROACH 3: LEAD FUNCTION (Looking Forward)
-- ============================================================================

-- Use LEAD() to get next values and compare
SELECT DISTINCT num AS ConsecutiveNums
FROM (
    SELECT 
        num,
        LEAD(num, 1) OVER (ORDER BY id) AS next_num,
        LEAD(num, 2) OVER (ORDER BY id) AS next_next_num
    FROM Logs
) t
WHERE num = next_num AND num = next_next_num;

-- ============================================================================
-- WHY YOUR ATTEMPT DIDN'T WORK
-- ============================================================================

-- Your code: COUNT(l.num >= 3)
-- Problems:
-- 1. l.num >= 3 is a boolean expression, but COUNT doesn't work this way
-- 2. Even if it did, you're checking if the NUMBER VALUE is >= 3, not if it appears 3+ times
-- 3. You're not checking for CONSECUTIVE occurrences
-- 4. Missing the concept of comparing adjacent rows

-- What you might have been thinking:
-- - COUNT(*) WHERE condition - but this needs to be in a subquery or HAVING clause
-- - But still wouldn't solve the "consecutive" requirement

-- ============================================================================
-- ALTERNATIVE THINKING: GROUP BY APPROACH (Doesn't Work Here)
-- ============================================================================

-- You might think: "Group by num and count"
-- SELECT num, COUNT(*) 
-- FROM Logs 
-- GROUP BY num 
-- HAVING COUNT(*) >= 3;

-- But this counts TOTAL occurrences, not CONSECUTIVE ones!
-- For our example, this would return both "1" (appears 3 times) and "2" (appears 2 times)
-- But "2" doesn't appear 3 times consecutively, so it shouldn't be in the result.

-- ============================================================================
-- FINAL RECOMMENDED SOLUTION
-- ============================================================================

-- The self-join approach is most intuitive and widely supported:
SELECT DISTINCT l1.num AS ConsecutiveNums
FROM Logs l1
JOIN Logs l2 ON l1.id = l2.id - 1
JOIN Logs l3 ON l2.id = l3.id - 1
WHERE l1.num = l2.num AND l2.num = l3.num; 