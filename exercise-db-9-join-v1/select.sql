SELECT reports.id,
students.fullname, 
students.class,
students.status,
reports.study,
reports.score
FROM reports , students 
WHERE reports.student_id = students.id AND status = 'active' AND reports.score < 70
ORDER BY reports.score ASC