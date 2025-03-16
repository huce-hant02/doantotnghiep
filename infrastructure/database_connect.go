package infrastructure

import (
	"doantotnghiep/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openConnection() (*gorm.DB, error) {
	connectSQL := "host=" + dbHost +
		" user=" + dbUser +
		" dbname=" + dbName +
		" password=" + dbPassword +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectSQL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		ErrLog.Printf("Not connect to database: %+v\n", err)
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

// InitDatabase open connection and migrate database
func InitDatabase(allowMigrate bool) error {
	var err error
	db, err = openConnection()
	if err != nil {
		return err
	}

	if allowMigrate {
		log.Println("Migrating database...")

		err := db.AutoMigrate(
			&model.SystemEmail{},
			&model.Role{},
			&model.Route{},
			&model.API{},
			&model.RoleRoute{},
			&model.RoleAPI{},
			&model.ModelType{},
			&model.ModelTypeField{},
			&model.RoleModelTypePermission{},
			&model.RoleFieldPermission{},
			&model.User{},
			&model.UserRole{},
			&model.ApiLog{},
			&model.SyncLog{},
			&model.System{},
			&model.ScholasticSemester{},
			&model.Course{},
			&model.CourseEmployee{},
			&model.RelatedCourse{},
			&model.Document{},
			&model.Employee{},
			&model.Faculty{},
			&model.SchoolYear{},
			&model.MajorGroup{},
			&model.Major{},
			&model.Specialty{},
			&model.KnowledgeGroup{},
			&model.BloomWord{},
			&model.BloomGroup{},
			&model.Level{},
			&model.Keyword{},
			&model.ExamType{},
			&model.TrainingTime{},
			&model.EditHistory{},
			&model.ExternalAssessmentTeam{},

			&model.EducationProgram{},
			&model.CourseEducationProgram{},
			&model.EducationTarget{},
			&model.EducationStandardOutput{},
			&model.ITUTable{},
			&model.EducationLog{},
			&model.Orientation{},

			&model.Outline{},
			&model.CourseStandardOutput{},
			&model.CourseTarget{},
			&model.CourseTargetEducationOutput{},
			&model.CourseOutputEducationOutput{},
			&model.CourseDocument{},
			&model.Rubric{},
			&model.RubricItem{},
			&model.TeachingPlan{},
			&model.TeachingPlanDocument{},
			&model.TeachingPlanStandardOutput{},
			&model.TeachingPlanResultEvaluate{},
			&model.OutlineEmployee{},
			&model.ResultEvaluate{},
			&model.ResultEvaluateExamType{},
			&model.ResultEvaluateRubric{},
			&model.PostOutput{},
			&model.Log{},
			&model.DefaultStandardOutput{},
			&model.DefaultITU{},
			&model.OutlineVersion{},
			&model.OutlineUpdateProcess{},
			&model.CdioProgramEmployee{},
			&model.LoeStudent{},
			&model.LoeCourseClass{},
			&model.LoeCourseClassStudent{},
		)
		if err == nil {
			log.Println("Done migrating database")
		} else {
			log.Println(err.Error())
		}
	}

	return nil
}
