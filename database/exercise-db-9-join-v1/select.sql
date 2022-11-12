SELECT
    reports.id,
    students.fullname,
    students.class,
    students.status,
    reports.study,
    reports.score
FROM students
INNER JOIN reports
ON students.id = reports.student_id
WHERE reports.score < 70 AND students.status = 'active'
ORDER BY reports.score ASC;