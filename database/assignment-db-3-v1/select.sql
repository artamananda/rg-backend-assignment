SELECT id, CONCAT(first_name, ' ', last_name) as fullname, SUBSTRING(exam_id, 1, 2) as class, (bahasa_indonesia + bahasa_inggris + matematika + ipa)/4 as average_score
FROM final_scores
WHERE exam_status = 'pass' AND fee_status != 'not paid'
ORDER BY average_score DESC
LIMIT 5;