package repository

import "strconv"

func getOutlineStatisticSQL(scholasticSemesterID uint) string {
	return `
	SELECT (SELECT COUNT(*) FROM outlines join scholastic_semester_outlines sso on outlines.id = sso.outline_id WHERE outlines.id > 0 AND outlines.deleted_at is null AND sso.scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + ` AND status = 'INITIAL') AS num_of_init,
		(SELECT COUNT(*) FROM outlines join scholastic_semester_outlines sso on outlines.id = sso.outline_id WHERE outlines.id > 0 AND outlines.deleted_at is null AND sso.scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + ` AND status = 'WRITING') AS num_of_writing,
		(SELECT COUNT(*) FROM outlines join scholastic_semester_outlines sso on outlines.id = sso.outline_id WHERE outlines.id > 0 AND outlines.deleted_at is null AND sso.scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + ` AND status = 'DONE') AS num_of_done,
		(SELECT COUNT(*) FROM outlines join scholastic_semester_outlines sso on outlines.id = sso.outline_id WHERE outlines.id > 0 AND outlines.deleted_at is null AND sso.scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + ` AND status = 'ADMIN_APPROVED') AS num_of_admin_approved,
		(SELECT COUNT(*) FROM outlines join scholastic_semester_outlines sso on outlines.id = sso.outline_id WHERE outlines.id > 0 AND outlines.deleted_at is null AND sso.scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + ` AND status = 'DEAN_APPROVED') AS num_of_dean_approved,
		(SELECT COUNT(*) FROM outlines join scholastic_semester_outlines sso on outlines.id = sso.outline_id WHERE outlines.id > 0 AND outlines.deleted_at is null AND sso.scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + ` AND status = 'TRAINING_APPROVED') AS num_of_training_approved,
		(SELECT COUNT(*) FROM outlines join scholastic_semester_outlines sso on outlines.id = sso.outline_id WHERE outlines.id > 0 AND outlines.deleted_at is null AND sso.scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + ` AND status = 'SECTION_APPROVED') AS num_of_section_approved,
		(SELECT COUNT(*) FROM outlines join scholastic_semester_outlines sso on outlines.id = sso.outline_id WHERE outlines.id > 0 AND outlines.deleted_at is null AND sso.scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + ` AND status = 'REJECTED') AS num_of_rejected,
		(SELECT COUNT(*) FROM courses WHERE id > 0 AND deleted_at is null AND scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + `) AS num_of_course,
		(SELECT COUNT(*) FROM employees JOIN users ON users.id = employees.user_id WHERE employees.id > 0 AND employees.deleted_at is null AND users.role NOT IN ('developer', 'dean_admin')) AS num_of_teacher,
		(SELECT COUNT(*) FROM majors WHERE id > 0 AND scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + `) AS num_of_major,
		(SELECT COUNT(*) FROM faculties WHERE id > 0 ) AS num_of_faculty,
		(SELECT COUNT(*) FROM documents WHERE id > 0 AND deleted_at is NULL AND in_library = true ) AS num_of_document_in_library,
		(SELECT COUNT(*) FROM documents WHERE id > 0 AND deleted_at is NULL ) AS num_of_document,
		(SELECT COUNT(*) FROM education_programs WHERE id > 0 AND deleted_at is null AND scholastic_semester_id = ` + strconv.Itoa(int(scholasticSemesterID)) + `) AS num_of_program
	`
}
