UPDATE students
SET address = 'Bandung'
WHERE address is NULL AND status = 'active';