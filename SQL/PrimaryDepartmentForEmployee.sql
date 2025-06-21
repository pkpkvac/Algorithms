-- Primary Department for Each Employee
-- Find each employee's primary department

-- PROBLEM UNDERSTANDING:
-- 1. Employees can belong to multiple departments
-- 2. If multiple departments: primary_flag='Y' marks the primary one
-- 3. If only ONE department: primary_flag='N' but we still want that department
-- 4. Return each employee with their primary department

-- Example analysis:
-- Employee 1: Only dept 1 (flag='N') → return dept 1 (their only dept)
-- Employee 2: Dept 1 (flag='Y'), dept 2 (flag='N') → return dept 1 (the primary)
-- Employee 3: Only dept 3 (flag='N') → return dept 3 (their only dept)  
-- Employee 4: Dept 2,3,4 (flags N,Y,N) → return dept 3 (the primary)

-- ============================================================================
-- APPROACH 1: UNION (Easiest to understand)
-- ============================================================================

-- Strategy: Get primary departments OR only departments
SELECT employee_id, department_id
FROM Employee 
WHERE primary_flag = 'Y'  -- Get employees with explicit primary departments

UNION

SELECT employee_id, department_id
FROM Employee e1
WHERE primary_flag = 'N' 
  AND NOT EXISTS (
    -- Only include if this employee has NO primary department
    SELECT 1 FROM Employee e2 
    WHERE e2.employee_id = e1.employee_id 
      AND e2.primary_flag = 'Y'
  );

-- How UNION works:
-- Part 1: Gets all rows where primary_flag='Y' (explicit primaries)
-- Part 2: Gets rows where primary_flag='N' AND no 'Y' exists for that employee
-- UNION combines them (removing duplicates)

-- ============================================================================
-- APPROACH 2: Window Function with ROW_NUMBER (More advanced)
-- ============================================================================

SELECT employee_id, department_id
FROM (
  SELECT employee_id, department_id,
         ROW_NUMBER() OVER (
           PARTITION BY employee_id 
           ORDER BY CASE WHEN primary_flag = 'Y' THEN 1 ELSE 2 END
         ) as rn
  FROM Employee
) ranked
WHERE rn = 1;

-- How this works:
-- 1. PARTITION BY employee_id: Group rows by employee
-- 2. ORDER BY CASE: Put 'Y' rows first (priority 1), 'N' rows second (priority 2)
-- 3. ROW_NUMBER(): Assign 1,2,3... within each employee group
-- 4. WHERE rn = 1: Take only the first row for each employee

-- ============================================================================
-- APPROACH 3: Conditional Aggregation (Alternative)
-- ============================================================================

SELECT 
    employee_id,
    CASE 
        WHEN SUM(CASE WHEN primary_flag = 'Y' THEN 1 ELSE 0 END) > 0
        THEN MAX(CASE WHEN primary_flag = 'Y' THEN department_id END)
        ELSE MAX(department_id)
    END as department_id
FROM Employee
GROUP BY employee_id;

-- How this works:
-- 1. GROUP BY employee_id: One row per employee
-- 2. SUM(CASE...): Count how many 'Y' flags this employee has
-- 3. If count > 0: Return the department_id where primary_flag='Y'
-- 4. If count = 0: Return any department_id (since there's only one)

-- ============================================================================
-- YOUR ORIGINAL ATTEMPT - WHY IT'S WRONG
-- ============================================================================

-- SELECT e.employee_id, e.department_id
-- FROM Employee e
-- GROUP BY e.employee_id        ← PROBLEM 1: Can't select department_id without grouping by it
-- CASE                          ← PROBLEM 2: CASE can't be used like this in FROM clause
--     WHEN                      ← PROBLEM 3: Syntax is completely wrong
--         COUNT(*) e.primary_flag = 1

-- ISSUES:
-- 1. GROUP BY employee_id but SELECT department_id (not allowed without aggregation)
-- 2. CASE statement syntax is wrong - it's not a clause, it's an expression
-- 3. Can't use CASE in FROM clause like that
-- 4. Logic doesn't match the problem requirements

-- ============================================================================
-- RECOMMENDED SOLUTION (Start with this one)
-- ============================================================================

-- The UNION approach is clearest for learning:
SELECT employee_id, department_id
FROM Employee 
WHERE primary_flag = 'Y'

UNION

SELECT employee_id, department_id
FROM Employee e1
WHERE primary_flag = 'N' 
  AND NOT EXISTS (
    SELECT 1 FROM Employee e2 
    WHERE e2.employee_id = e1.employee_id 
      AND e2.primary_flag = 'Y'
  ); 