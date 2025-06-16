-- Managers and Their Reports
-- Find employees who are managers (have direct reports) and calculate statistics

-- Strategy:
-- 1. Self-join Employees table: managers with their direct reports
-- 2. Join condition: manager.employee_id = report.reports_to
-- 3. Group by manager info to aggregate report statistics
-- 4. Use COUNT() for number of reports and ROUND(AVG()) for average age
-- 5. Order by employee_id

SELECT 
    m.employee_id,
    m.name,
    COUNT(r.employee_id) as reports_count,
    ROUND(AVG(r.age)) as average_age
FROM Employees m
INNER JOIN Employees r ON m.employee_id = r.reports_to
GROUP BY m.employee_id, m.name
ORDER BY m.employee_id;

-- How the self-join works:
-- Table alias 'm' = managers
-- Table alias 'r' = reports (employees who report to managers)
-- JOIN condition: m.employee_id = r.reports_to
--   This connects each manager to their direct reports

-- Example 1 walkthrough:
-- Original data:
-- | employee_id | name    | reports_to | age |
-- | 9           | Hercy   | null       | 43  |
-- | 6           | Alice   | 9          | 41  |
-- | 4           | Bob     | 9          | 36  |
-- | 2           | Winston | null       | 37  |

-- After self-join (m.employee_id = r.reports_to):
-- | m.employee_id | m.name | r.employee_id | r.age |
-- | 9             | Hercy  | 6             | 41    |
-- | 9             | Hercy  | 4             | 36    |

-- After GROUP BY and aggregation:
-- | employee_id | name  | reports_count | average_age |
-- | 9           | Hercy | 2             | 39          |
-- (AVG(41, 36) = 38.5 → ROUND() = 39)

-- Key points:
-- - INNER JOIN automatically filters to only managers (employees with reports)
-- - Self-join creates manager-report pairs
-- - GROUP BY manager info aggregates all reports per manager
-- - COUNT(r.employee_id) counts number of direct reports
-- - ROUND(AVG(r.age)) calculates and rounds average age of reports 