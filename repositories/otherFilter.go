package repository

import (
	"doantotnghiep/model"
	"errors"
	"strings"

	"gorm.io/gorm"
)

/* AdvanceFilter */

func (r *advanceFilterV2Repo) addFilterCourse(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}) (*gorm.DB, error) {
	item := filterModel.(model.Course)
	if len(item.OtherFilter.ListCode) > 0 {
		tempDB = tempDB.Where("code IN (?)", item.OtherFilter.ListCode)
	}
	if len(item.OtherFilter.ListId) > 0 {
		tempDB = tempDB.Where("id IN (?)", item.OtherFilter.ListId)
	}
	return tempDB, nil
}

func (r *advanceFilterV2Repo) addFilterEmployee(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}) (*gorm.DB, error) {
	item := filterModel.(model.Employee)
	if len(item.OtherFilter.ListId) > 0 {
		tempDB = tempDB.Where("id IN (?)", item.OtherFilter.ListId)
	}
	if len(item.OtherFilter.ListCode) > 0 {
		tempDB = tempDB.Where("code IN (?)", item.OtherFilter.ListCode)
	}
	if len(item.OtherFilter.ListDepartmentCode) > 0 {
		for _, r := range item.OtherFilter.ListDepartmentCode {
			tempDB = tempDB.Where("employee_departments ILIKE '%" + r + "%'")
		}
	}
	return tempDB, nil
}

/* CDIO */
func (r *advanceFilterV2Repo) addFilterOutline(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}, claims map[string]interface{}) (*gorm.DB, error) {
	outline := filterModel.(model.Outline)
	role, err := getRoleFromClaims(claims)
	if err != nil {
		return nil, err
	}
	userId, ok := (claims["user_id"].(float64))
	if !ok {
		return nil, errors.New("could not parse user_id from claims")
	}
	utype, okUserType := claims["utype"].(string)
	if !okUserType {
		return nil, errors.New("could not parse utype from claims")
	}
	var employee model.Employee
	var department model.Faculty
	var team model.ExternalAssessmentTeam
	if utype == "NHANSU" {
		if err := db.Model(&model.Employee{}).Where("user_id = ?", uint(userId)).Find(&employee).Error; err != nil {
			return nil, err
		}
	} else if utype == "PHONGBAN" {
		if err := db.Model(&model.Faculty{}).Where("user_id = ?", uint(userId)).Find(&department).Error; err != nil {
			return nil, err
		}
	} else if utype == "DOANDANHGIANGOAI" {
		if err := db.Model(&model.ExternalAssessmentTeam{}).Where("user_id = ?", uint(userId)).Find(&team).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Xác thực không hợp lệ")
	}
	// #########################################################################################################

	if len(outline.OtherFilter.ListCourseCode) > 0 {
		tempDB = tempDB.Where("course_code IN (?)", outline.OtherFilter.ListCourseCode)
	}
	if len(outline.OtherFilter.ListOutlineId) > 0 {
		tempDB = tempDB.Where("id IN (?)", outline.OtherFilter.ListOutlineId)
	}
	if len(outline.OtherFilter.ListCourseId) > 0 {
		tempDB = tempDB.Where("course_id IN (?)", outline.OtherFilter.ListCourseId)
	}
	if outline.OtherFilter.Course != nil {
		if outline.OtherFilter.Course.Code != "" && len(outline.OtherFilter.Course.Code) > 0 {
			tempDB = tempDB.Where("id IN (SELECT id from outlines WHERE course_id IN (SELECT id FROM courses c WHERE lower(c.code) LIKE '%" + strings.ToLower(outline.OtherFilter.Course.Code) + "%'))")
		}
		if outline.OtherFilter.Course.Title != "" && len(outline.OtherFilter.Course.Title) > 0 {
			tempDB = tempDB.Where("id IN (SELECT id from outlines WHERE course_id IN (SELECT id FROM courses c WHERE lower(c.title) LIKE '%" + strings.ToLower(outline.OtherFilter.Course.Title) + "%'))")
		}
		if outline.OtherFilter.Course.DepartmentCode != "" {
			tempDB = tempDB.Where("id IN (SELECT id from outlines WHERE deleted_at IS NULL AND course_id IN (SELECT id FROM courses c WHERE deleted_at IS NULL AND department_code = ?))", outline.OtherFilter.Course.DepartmentCode)
		}
	}
	if len(outline.OtherFilter.ListEduId) > 0 {
		tempDB = tempDB.Where("course_id IN (SELECT course_id FROM course_education_programs WHERE education_program_id IN (?))", outline.OtherFilter.ListEduId)
	}
	if outline.OtherFilter.IsAssigned != nil && *outline.OtherFilter.IsAssigned == model.TrueValue {
		tempDB = tempDB.Where("course_id IN (SELECT course_id FROM course_employees WHERE employee_id = ?)", employee.ID)
	}

	switch role {
	case "super-admin", "admin", "phong-dao-tao", "ban-giam-hieu", "can-bo-thu-vien":
		break
	case "truong-khoa", "pho-khoa", "giao-vu-khoa", "truong-bo-mon", "giang-vien":
		departmentCodes := []string{}
		for _, r := range employee.EmployeeDepartments {
			departmentCodes = append(departmentCodes, r.Department)
		}
		tempDB = tempDB.Where("course_id IN (SELECT id FROM courses WHERE deleted_at IS NULL AND department_code IN (?)) OR course_id IN (SELECT course_id FROM course_employees WHERE employee_id = ?)", departmentCodes, employee.ID)
		break
	case "doan-danh-gia-ngoai":
		departmentCodes := []string{}
		for _, c := range team.ListDepartment {
			departmentCodes = append(departmentCodes, c)
		}

		var courseIDs []uint
		err := db.Model(&model.CourseEducationProgram{}).Select("course_id").
			Where("education_program_id IN (SELECT id FROM education_programs WHERE deleted_at IS NULL AND major_id IN (SELECT id FROM majors WHERE deleted_at IS NULL AND department_code IN (?)))", departmentCodes).
			Find(&courseIDs).Error
		if err != nil {
			tempDB = tempDB.Where("id = 0")
			return tempDB, nil
		}

		tempDB = tempDB.Where("draft = false AND course_id IN (?)", courseIDs)
		break
	case "phong-ban-khac":
		tempDB = tempDB.Where("draft = false")
		break
	default:
		tempDB = tempDB.Where("id = 0")
		break
	}

	return tempDB, nil
}

func (r *advanceFilterV2Repo) addFilterEducationProgram(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}, claims map[string]interface{}) (*gorm.DB, error) {
	cdioProgram := filterModel.(model.EducationProgram)
	userId, okUserId := claims["user_id"].(float64)
	if !okUserId {
		return nil, errors.New("could not parse user_id from claims")
	}
	utype, okUserType := claims["utype"].(string)
	if !okUserType {
		return nil, errors.New("could not parse utype from claims")
	}
	role, okRole := claims["role"].(string)
	if !okRole {
		return nil, errors.New("could not parse role from claims")
	}
	var employee model.Employee
	var department model.Faculty
	var team model.ExternalAssessmentTeam
	if utype == "NHANSU" {
		if err := db.Model(&model.Employee{}).Where("user_id = ?", uint(userId)).Find(&employee).Error; err != nil {
			return nil, err
		}
	} else if utype == "PHONGBAN" {
		if err := db.Model(&model.Faculty{}).Where("user_id = ?", uint(userId)).Find(&department).Error; err != nil {
			return nil, err
		}
	} else if utype == "DOANDANHGIANGOAI" {
		if err := db.Model(&model.ExternalAssessmentTeam{}).Where("user_id = ?", uint(userId)).Find(&team).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Xác thực không hợp lệ")
	}

	// #########################################################################################################
	if len(cdioProgram.OtherFilter.ListId) > 0 {
		tempDB = tempDB.Where("id IN (?)", cdioProgram.OtherFilter.ListId)
	}
	if len(cdioProgram.OtherFilter.ListMajorId) > 0 {
		tempDB = tempDB.Where("major_id IN (?)", cdioProgram.OtherFilter.ListMajorId)
	}
	if cdioProgram.OtherFilter.IsAssigned != nil && *cdioProgram.OtherFilter.IsAssigned == model.TrueValue {
		tempDB = tempDB.Where("id IN (SELECT education_program_id FROM cdio_program_employees WHERE employee_id = ?)", employee.ID)
	}

	switch role {
	case "super-admin", "admin", "phong-dao-tao", "ban-giam-hieu", "can-bo-thu-vien":
		break
	case "truong-khoa", "pho-khoa", "giao-vu-khoa", "truong-bo-mon", "giang-vien":
		// departmentCodes := []string{}
		// for _, r := range employee.EmployeeDepartments {
		// 	departmentCodes = append(departmentCodes, r.Department)
		// }
		// tempDB = tempDB.Where("major_id IN (SELECT id FROM majors WHERE deleted_at IS NULL AND department_code IN (?)) OR id IN (SELECT education_program_id FROM cdio_program_employees WHERE employee_id = ?)", departmentCodes, employee.ID)
		break
	case "doan-danh-gia-ngoai":
		departmentCodes := []string{}
		for _, c := range team.ListDepartment {
			departmentCodes = append(departmentCodes, c)
		}
		tempDB = tempDB.Where("draft = false AND major_id IN (SELECT id FROM majors WHERE deleted_at IS NULL AND department_code IN (?))", departmentCodes)
		break
	case "phong-ban-khac":
		tempDB = tempDB.Where("draft = false")
		break
	default:
		tempDB = tempDB.Where("id = 0")
		break
	}

	return tempDB, nil
}

func (r *advanceFilterV2Repo) addFilterCourseEducationProgram(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}) (*gorm.DB, error) {
	courseEducationProgram := filterModel.(model.CourseEducationProgram)
	if len(courseEducationProgram.OtherFilter.ListCourseId) > 0 {
		tempDB = tempDB.Where("course_id IN (?)", courseEducationProgram.OtherFilter.ListCourseId)
	}
	if len(courseEducationProgram.OtherFilter.ListEducationProgramId) > 0 {
		tempDB = tempDB.Where("education_program_id IN (?)", courseEducationProgram.OtherFilter.ListEducationProgramId)
	}
	return tempDB, nil
}

/* ABET */
func (r *advanceFilterV2Repo) addFilterAbetOutline(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}, claims map[string]interface{}) (*gorm.DB, error) {
	outline := filterModel.(model.AbetOutline)
	role, err := getRoleFromClaims(claims)
	if err != nil {
		return nil, err
	}
	userId, ok := (claims["user_id"].(float64))
	if !ok {
		return nil, errors.New("could not parse user_id from claims")
	}
	utype, okUserType := claims["utype"].(string)
	if !okUserType {
		return nil, errors.New("could not parse utype from claims")
	}
	var employee model.Employee
	var department model.Faculty
	var team model.ExternalAssessmentTeam
	if utype == "NHANSU" {
		if err := db.Model(&model.Employee{}).Where("user_id = ?", uint(userId)).Find(&employee).Error; err != nil {
			return nil, err
		}
	} else if utype == "PHONGBAN" {
		if err := db.Model(&model.Faculty{}).Where("user_id = ?", uint(userId)).Find(&department).Error; err != nil {
			return nil, err
		}
	} else if utype == "DOANDANHGIANGOAI" {
		if err := db.Model(&model.ExternalAssessmentTeam{}).Where("user_id = ?", uint(userId)).Find(&team).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Xác thực không hợp lệ")
	}
	// #########################################################################################################
	if len(outline.OtherFilter.ListCode) > 0 {
		tempDB = tempDB.Where("code IN (?)", outline.OtherFilter.ListCode)
	}
	if len(outline.OtherFilter.ListId) > 0 {
		tempDB = tempDB.Where("id IN (?)", outline.OtherFilter.ListId)
	}
	if len(outline.OtherFilter.ListCourseId) > 0 {
		tempDB = tempDB.Where("course_id IN (?)", outline.OtherFilter.ListCourseId)
	}
	if outline.OtherFilter.Course != nil {
		if outline.OtherFilter.Course.Code != "" && len(outline.OtherFilter.Course.Code) > 0 {
			tempDB = tempDB.Where("id IN (SELECT id from abet_outlines WHERE course_id IN (SELECT id FROM courses c WHERE lower(c.code) LIKE '%" + strings.ToLower(outline.OtherFilter.Course.Code) + "%'))")
		}
		if outline.OtherFilter.Course.Title != "" && len(outline.OtherFilter.Course.Title) > 0 {
			tempDB = tempDB.Where("id IN (SELECT id from abet_outlines WHERE course_id IN (SELECT id FROM courses c WHERE lower(c.title) LIKE '%" + strings.ToLower(outline.OtherFilter.Course.Title) + "%'))")
		}
		if outline.OtherFilter.Course.DepartmentCode != "" {
			tempDB = tempDB.Where("id IN (SELECT id from abet_outlines WHERE deleted_at IS NULL AND course_id IN (SELECT id FROM courses c WHERE deleted_at IS NULL AND department_code = ?))", outline.OtherFilter.Course.DepartmentCode)
		}
	}
	if outline.OtherFilter.IsAssigned != nil && *outline.OtherFilter.IsAssigned == model.TrueValue {
		tempDB = tempDB.Where("course_id IN (SELECT course_id FROM course_employees WHERE employee_id = ?)", employee.ID)
	}

	switch role {
	case "super-admin", "admin", "phong-dao-tao", "ban-giam-hieu", "can-bo-thu-vien":
		break
	case "truong-khoa", "pho-khoa", "giao-vu-khoa", "giang-vien", "truong-bo-mon":
		departmentCodes := []string{}
		for _, r := range employee.EmployeeDepartments {
			departmentCodes = append(departmentCodes, r.Department)
		}
		tempDB = tempDB.Where("course_id IN (SELECT id FROM courses WHERE deleted_at IS NULL AND department_code IN (?)) OR course_id IN (SELECT course_id FROM course_employees WHERE employee_id = ?)", departmentCodes, employee.ID)
		break
	case "doan-danh-gia-ngoai":
		departmentCodes := []string{}
		for _, c := range team.ListDepartment {
			departmentCodes = append(departmentCodes, c)
		}
		var courseIDs []uint
		err := db.Model(&model.AbetRelationshipCourseAndProgram{}).Select("course_id").
			Where("abet_program_id IN (SELECT id FROM abet_programs WHERE deleted_at IS NULL AND major_id IN (SELECT id FROM majors WHERE deleted_at IS NULL AND department_code IN (?)))", departmentCodes).
			Find(&courseIDs).Error
		if err != nil {
			tempDB = tempDB.Where("id = 0")
			return tempDB, nil
		}

		tempDB = tempDB.Where("draft = false AND course_id IN (?)", courseIDs)
		break
	case "phong-ban-khac":
		tempDB = tempDB.Where("draft = false")
		break
	default:
		tempDB = tempDB.Where("id = 0")
		break
	}

	return tempDB, nil
}

func (r *advanceFilterV2Repo) addFilterAbetProgram(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}, claims map[string]interface{}) (*gorm.DB, error) {
	abetProgram := filterModel.(model.AbetProgram)
	userId, okUserId := claims["user_id"].(float64)
	if !okUserId {
		return nil, errors.New("could not parse user_id from claims")
	}
	utype, okUserType := claims["utype"].(string)
	if !okUserType {
		return nil, errors.New("could not parse utype from claims")
	}
	role, okRole := claims["role"].(string)
	if !okRole {
		return nil, errors.New("could not parse role from claims")
	}
	var employee model.Employee
	var department model.Faculty
	var team model.ExternalAssessmentTeam
	if utype == "NHANSU" {
		if err := db.Model(&model.Employee{}).Where("user_id = ?", uint(userId)).Find(&employee).Error; err != nil {
			return nil, err
		}
	} else if utype == "PHONGBAN" {
		if err := db.Model(&model.Faculty{}).Where("user_id = ?", uint(userId)).Find(&department).Error; err != nil {
			return nil, err
		}
	} else if utype == "DOANDANHGIANGOAI" {
		if err := db.Model(&model.ExternalAssessmentTeam{}).Where("user_id = ?", uint(userId)).Find(&team).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Xác thực không hợp lệ")
	}
	// #########################################################################################################
	if len(abetProgram.OtherFilter.ListCode) > 0 {
		tempDB = tempDB.Where("code IN (?)", abetProgram.OtherFilter.ListCode)
	}
	if len(abetProgram.OtherFilter.ListId) > 0 {
		tempDB = tempDB.Where("id IN (?)", abetProgram.OtherFilter.ListId)
	}
	if len(abetProgram.OtherFilter.ListMajorId) > 0 {
		tempDB = tempDB.Where("major_id IN (?)", abetProgram.OtherFilter.ListMajorId)
	}
	if abetProgram.OtherFilter.IsAssigned != nil && *abetProgram.OtherFilter.IsAssigned == model.TrueValue {
		tempDB = tempDB.Where("id IN (SELECT abet_program_id FROM abet_program_employees WHERE employee_id = ?)", employee.ID)
	}

	switch role {
	case "super-admin", "admin", "phong-dao-tao", "ban-giam-hieu", "can-bo-thu-vien":
		break
	case "truong-khoa", "pho-khoa", "giao-vu-khoa", "truong-bo-mon", "giang-vien":
		// departmentCodes := []string{}
		// for _, r := range employee.EmployeeDepartments {
		// 	departmentCodes = append(departmentCodes, r.Department)
		// }
		// tempDB = tempDB.Where("major_id IN (SELECT id FROM majors WHERE deleted_at IS NULL AND department_code IN (?)) OR id IN (SELECT abet_program_id FROM abet_program_employees WHERE employee_id = ?)", departmentCodes, employee.ID)
		break
	case "doan-danh-gia-ngoai":
		departmentCodes := []string{}
		for _, c := range team.ListDepartment {
			departmentCodes = append(departmentCodes, c)
		}
		tempDB = tempDB.Where("draft = false AND major_id IN (SELECT id FROM majors WHERE deleted_at IS NULL AND department_code IN (?))", departmentCodes)
		break
	case "phong-ban-khac":
		tempDB = tempDB.Where("draft = false")
		break
	default:
		tempDB = tempDB.Where("id = 0")
		break
	}
	return tempDB, nil
}

func (r *advanceFilterV2Repo) addFilterAbetProgramCourse(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}) (*gorm.DB, error) {
	item := filterModel.(model.AbetRelationshipCourseAndProgram)
	if len(item.OtherFilter.ListId) > 0 {
		tempDB = tempDB.Where("id IN (?)", item.OtherFilter.ListId)
	}
	if len(item.OtherFilter.ListProgramId) > 0 {
		tempDB = tempDB.Where("abet_program_id IN (?)", item.OtherFilter.ListProgramId)
	}
	if len(item.OtherFilter.ListCourseId) > 0 {
		tempDB = tempDB.Where("course_id IN (?)", item.OtherFilter.ListCourseId)
	}
	return tempDB, nil
}

/* LOE */
func (r *advanceFilterV2Repo) addFilterAbetExamMatrix(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}) (*gorm.DB, error) {
	record := filterModel.(model.AbetExamMatrix)
	if len(record.OtherFilter.ListCLOId) > 0 {
		tempDB = tempDB.Where("program_course_pi_lo_id IN (SELECT id FROM abet_relationship_program_course_pi_and_los WHERE abet_outline_learning_outcome_id IN (?))", record.OtherFilter.ListCLOId)
	}
	return tempDB, nil
}

func (r *advanceFilterV2Repo) addFilterLoeStudent(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}) (*gorm.DB, error) {
	item := filterModel.(model.LoeStudent)
	if len(item.OtherFilter.ListId) > 0 {
		tempDB = tempDB.Where("id IN (?)", item.OtherFilter.ListId)
	}
	if len(item.OtherFilter.ListCode) > 0 {
		tempDB = tempDB.Where("code IN (?)", item.OtherFilter.ListCode)
	}
	return tempDB, nil
}

func (r *advanceFilterV2Repo) addFilterLoeCourseClass(db *gorm.DB, tempDB *gorm.DB, filterModel interface{}) (*gorm.DB, error) {
	item := filterModel.(model.LoeCourseClass)
	if len(item.OtherFilter.ListCourseId) > 0 {
		tempDB = tempDB.Where("course_id IN (?)", item.OtherFilter.ListCourseId)
	}
	return tempDB, nil
}
