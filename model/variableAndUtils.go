package model

import "reflect"

const (
	DefaultPassword = "1Abcde@2023"

	// outline status
	OUTLINE_INIT_STATUS       = "INITIAL"
	OUTLINE_WRITING_STATUS    = "WRITING"
	OUTLINE_DONE_STATUS       = "DONE"
	OUTLINE_SECTION_APPROVED  = "SECTION_APPROVED"
	OUTLINE_DEAN_APPROVED     = "DEAN_APPROVED"
	OUTLINE_TRAINING_APPROVED = "TRAINING_APPROVED"
	OUTLINE_ADMIN_APPROVED    = "ADMIN_APPROVED"
	OUTLINE_REJECTED          = "REJECTED"

	// user role
	ROLE_DEV              = "super-admin"
	ROLE_ADMIN            = "admin"
	ROLE_HEAD_OF_TRAINING = "phong-dao-tao"
	ROLE_HEAD_OF_UNI      = "ban-giam-hieu"
	ROLE_DEAN             = "truong-khoa"
	ROLE_HEAD_OF_SECTION  = "truong-bo-mon"
	ROLE_TEACHER          = "giang-vien"
	ROLE_LIB              = "can-bo-thu-vien"
	ROLE_GUEST            = "guest"

	// error
	NOT_FOUND = "not_found"

	// operation
	OperationAdd    = "Add"
	OperationUpdate = "Update"
	OperationDelete = "Delete"
	OperationReject = "Reject"

	// library url
	// DLIB_BASE_URL        = "https://dlib.phenikaa-uni.edu.vn/"
	DLIB_BASE_URL          = "https://10.20.2.101/" // "https://dlib.phenikaa-uni.edu.vn/" //"https://piditi.com/dlib/" // http://10.20.2.101/
	DSPACE_LIBRARY_ACCOUNT = "contact@piditi.com"
	DSPACE_LIBRARY_PASS    = "pdt2023$"

	// ELIB_BASE_URL        = "https://elib.phenikaa-uni.edu.vn/api/v1/"
	ELIB_BASE_URL        = "https://10.20.2.100/" // https://piditi.com/elib/api/v1/" // http://10.20.2.100/
	KOHA_LIBRARY_ACCOUNT = "apikoha"
	KOHA_LIBRARY_PASS    = "Ph3n1k44@2020"
	// KOHA_CLIENT_ID       = "81390aca-6665-4531-9e35-8c013fff5ab7"
	// KOHA_CLIENT_SECRET   = "fb2cc2d8-6c6f-4310-8eb1-eadd5494b700"

	// document type
	DOCTYPE_BOOK         = "book"
	DOCTYPE_INTERNAL_DOC = "internal_document"
	DOCTYPE_JOURNAL      = "journal"
	DOCTYPE_THESIS       = "thesis"
	DOCTYPE_OTHER_DOC    = "other_document"
	DOCTYPE_E_DOC        = "electronic_document"
	DOCTYPE_SR_DOC       = "scientific_research_document"
	// other constant

	BaseDN           string = "dc=phenikaa-uni,dc=edu,dc=vn"
	BindDN           string = "cn=App Thu vien,ou=Services_Account,ou=PNK,dc=phenikaa-uni,dc=edu,dc=vn"
	LDAPPassword     string = "1Abcde@2022"
	LDAPUserFilter   string = "(&(objectClass=*)"
	LDAPGroupFilter  string = "(memberUid=%s)"
	LDAPHostUsername string = "App Thu vien"
)

type SortItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var (
	TmpSort        = []SortItem{}
	TrueValue bool = true

	AccessTokenTime  int64 = 24
	RefreshTokenTime int64 = 72
	LDAPTokenTime    int64 = 180

	EditHistoryApprovedStatus string = "approved"

	APILoggerIgnoreRoutes = []string{
		"/imports",
		"/upload",
		"/apis",
		"/students/anticipate-graduation-reward",
	}

	MODEL_MAP = map[string]interface{}{
		/* CORE */
		"systemEmails":             SystemEmail{},
		"users":                    User{},
		"roles":                    Role{},
		"routes":                   Route{},
		"apis":                     API{},
		"modelTypes":               ModelType{},
		"modelTypeFields":          ModelTypeField{},
		"userRoles":                UserRole{},
		"roleRoutes":               RoleRoute{},
		"roleApis":                 RoleAPI{},
		"roleModelTypePermissions": RoleModelTypePermission{},
		"roleFieldPermissions":     RoleFieldPermission{},

		"systems":             System{},
		"scholasticSemesters": ScholasticSemester{},
		"courses":             Course{},
		"courseTypes":         CourseType{},
		"documents":           Document{},
		"employees":           Employee{},
		"faculties":           Faculty{},
		"schoolYears":         SchoolYear{},
		"majors":              Major{},
		"majorGroups":         MajorGroup{},
		"subjects":            Subject{},
		"knowledgeGroups":     KnowledgeGroup{},
		"logs":                Log{},
		"bloomWords":          BloomWord{},
		"bloomGroups":         BloomGroup{},
		"levels":              Level{},
		"keywords":            Keyword{},
		"examTypes":           ExamType{},
		"specialties":         Specialty{},
		"trainingTimes":       TrainingTime{},

		"defaultITUs":            DefaultITU{},
		"defaultStandardOutputs": DefaultStandardOutput{},
		/* ENTITY */
		"outlines":                     Outline{},
		"outlineEmployees":             OutlineEmployee{},
		"outlineVersions":              OutlineVersion{},
		"outlineDocuments":             OutlineDocument{},
		"outlineUpdateProcesses":       OutlineUpdateProcess{},
		"courseEmployees":              CourseEmployee{},
		"courseTargets":                CourseTarget{},
		"courseStandardOutputs":        CourseStandardOutput{},
		"courseEducationPrograms":      CourseEducationProgram{},
		"courseTargetEducationOutputs": CourseTargetEducationOutput{},
		"courseOutputEducationOutputs": CourseOutputEducationOutput{},
		"rubrics":                      Rubric{},
		"rubricItems":                  RubricItem{},
		"resultEvaluates":              ResultEvaluate{},
		"resultEvaluateRubrics":        ResultEvaluateRubric{},
		"resultEvaluateExamTypes":      ResultEvaluateExamType{},
		"postOutputs":                  PostOutput{},
		"teachingPlans":                TeachingPlan{},
		"teachingPlanDocuments":        TeachingPlanDocument{},
		"teachingPlanStandardOutputs":  TeachingPlanStandardOutput{},
		"teachingPlanResultEvaluates":  TeachingPlanResultEvaluate{},
		"educationPrograms":            EducationProgram{},
		"educationTargets":             EducationTarget{},
		"educationStandardOutputs":     EducationStandardOutput{},
		"orientations":                 Orientation{},
		"educationLogs":                EducationLog{},
		"ituTables":                    ITUTable{},
		"schoolYearEducationPrograms":  SchoolYearEducationProgram{},
		// "signatures":               Signature{},
		"editHistories":           EditHistory{},
		"cdioProgramEmployees":    CdioProgramEmployee{},
		"externalAssessmentTeams": ExternalAssessmentTeam{},
		"syncLogs":                SyncLog{},
		"relatedCourses":          RelatedCourse{},
	}

	LIST_MODEL_MAP = map[string]interface{}{
		"systemEmails":             []SystemEmail{},
		"users":                    []User{},
		"roles":                    []Role{},
		"routes":                   []Route{},
		"apis":                     []API{},
		"modelTypes":               []ModelType{},
		"modelTypeFields":          []ModelTypeField{},
		"userRoles":                []UserRole{},
		"roleRoutes":               []RoleRoute{},
		"roleApis":                 []RoleAPI{},
		"roleModelTypePermissions": []RoleModelTypePermission{},
		"roleFieldPermissions":     []RoleFieldPermission{},

		"systems":             []System{},
		"scholasticSemesters": []ScholasticSemester{},
		"courses":             []Course{},
		"courseTypes":         []CourseType{},
		"documents":           []Document{},
		"employees":           []Employee{},
		"faculties":           []Faculty{},
		"schoolYears":         []SchoolYear{},
		"majors":              []Major{},
		"majorGroups":         []MajorGroup{},
		"subjects":            []Subject{},
		"knowledgeGroups":     []KnowledgeGroup{},
		"logs":                []Log{},
		"bloomWords":          []BloomWord{},
		"bloomGroups":         []BloomGroup{},
		"levels":              []Level{},
		"keywords":            []Keyword{},
		"examTypes":           []ExamType{},
		"specialties":         []Specialty{},
		"trainingTimes":       []TrainingTime{},

		"defaultITUs":            []DefaultITU{},
		"defaultStandardOutputs": []DefaultStandardOutput{},
		/* ENTITY */
		"outlines":                     []Outline{},
		"outlineEmployees":             []OutlineEmployee{},
		"outlineVersions":              []OutlineVersion{},
		"outlineDocuments":             []OutlineDocument{},
		"outlineUpdateProcesses":       []OutlineUpdateProcess{},
		"courseEmployees":              []CourseEmployee{},
		"courseTargets":                []CourseTarget{},
		"courseStandardOutputs":        []CourseStandardOutput{},
		"courseEducationPrograms":      []CourseEducationProgram{},
		"courseTargetEducationOutputs": []CourseTargetEducationOutput{},
		"courseOutputEducationOutputs": []CourseOutputEducationOutput{},
		"rubrics":                      []Rubric{},
		"rubricItems":                  []RubricItem{},
		"resultEvaluates":              []ResultEvaluate{},
		"resultEvaluateRubrics":        []ResultEvaluateRubric{},
		"resultEvaluateExamTypes":      []ResultEvaluateExamType{},
		"postOutputs":                  []PostOutput{},
		"teachingPlans":                []TeachingPlan{},
		"teachingPlanDocuments":        []TeachingPlanDocument{},
		"teachingPlanStandardOutputs":  []TeachingPlanStandardOutput{},
		"teachingPlanResultEvaluates":  []TeachingPlanResultEvaluate{},
		"educationPrograms":            []EducationProgram{},
		"educationTargets":             []EducationTarget{},
		"educationStandardOutputs":     []EducationStandardOutput{},
		"orientations":                 []Orientation{},
		"educationLogs":                []EducationLog{},
		"ituTables":                    []ITUTable{},
		"schoolYearEducationPrograms":  []SchoolYearEducationProgram{},

		// "signatures":               Signature{},
		"editHistories":           []EditHistory{},
		"cdioProgramEmployees":    []CdioProgramEmployee{},
		"loeStudents":             []LoeStudent{},
		"loeCourseClasses":        []LoeCourseClass{},
		"loeCourseClassStudents":  []LoeCourseClassStudent{},
		"externalAssessmentTeams": []ExternalAssessmentTeam{},
		"syncLogs":                []SyncLog{},
		"relatedCourses":          []RelatedCourse{},
	}

	// map for handling association
	MODEL_ASSOCIATION_NAME_MAP = map[string]string{
		//course
		"Course_Document": "Documents",
		// "Course_KnowledgeGroup":   "knowledgeGroups",
		"Course_Employee":         "Teachers",
		"Course_EducationProgram": "EducationPrograms",
		"Course_Outline":          "Outlines",

		//outline
		"Outline_CourseTarget":         "CourseTargets",
		"Outline_Rubric":               "Rubrics",
		"Outline_TeachingPlan":         "TeachingPlans",
		"Outline_CourseStandardOutput": "CourseStandardOutputs",
		"Outline_Employee":             "Teachers",
		"Outline_ResultEvaluate":       "ResultEvaluates",
		"Outline_Assignee":             "AssigneeInfo",
		"Outline_Document":             "Documents",

		// teachingPlan
		"TeachingPlan_Document":       "Documents",
		"EducationProgram_Course":     "Courses",
		"TeachingPlan_ResultEvaluate": "AssessmentPost",
		"TeachingPlan_StandardOutput": "StandardOutputs",

		// result evaluate
		"ResultEvaluate_Rubric":       "Rubrics",
		"ResultEvaluate_PostOutput":   "PostOutputs",
		"ResultEvaluate_EvaluateForm": "EvaluateForm",

		// bloom word
		"BloomWord_Level": "Levels",

		// level
		"Level_Keyword": "Keywords",
	}
	//
	// map for handling association
	MODEL_ASSOCIATION_TYPE_MAP = map[string]reflect.Type{
		//course
		// "Course_KnowledgeGroup":   reflect.TypeOf(KnowledgeGroup{}),
		"Course_Document":         reflect.TypeOf(Document{}),
		"Course_Employee":         reflect.TypeOf(Employee{}),
		"Course_EducationProgram": reflect.TypeOf(EducationProgram{}),
		"Course_Outline":          reflect.TypeOf(Outline{}),

		//outline
		"Outline_CourseTarget":         reflect.TypeOf(CourseTarget{}),
		"Outline_Rubric":               reflect.TypeOf(Rubric{}),
		"Outline_TeachingPlan":         reflect.TypeOf(TeachingPlan{}),
		"Outline_CourseStandardOutput": reflect.TypeOf(CourseStandardOutput{}),
		"Outline_Employee":             reflect.TypeOf(Employee{}),
		"Outline_ResultEvaluate":       reflect.TypeOf(ResultEvaluate{}),
		"Outline_Assignee":             reflect.TypeOf(Employee{}),
		"Outline_Document":             reflect.TypeOf(Document{}),
	}
	// map for preload in advance filter function
	MODEL_LIST_ASSOCIATION_MAP_V2 = map[string]map[string]interface{}{
		"users": {
			"UserRoles": "",
		},
		"modelTypes": {
			"Fields": "",
		},
		"roleModelTypePermissions": {
			"Fields": "",
		},
		"editHistories": {
			"Modifier": "",
		},
		"abetPrograms": {
			"Constituents.Peos":                  "",
			"Major":                              "",
			"Courses.Orientations":               "", // .Course.Faculty
			"Courses.PerformanceIndicators":      "",
			"Courses.PSOs":                       "",
			"PSOs.PerformanceIndicators.Keyword": "",
			"PSOs.Peos":                          "",
			"PSOs.Keyword":                       "",
			"PEOs":                               "",
			"SchoolYears":                        "",
			"SchoolYears.SchoolYear":             "",
			"Logs":                               "",
			"Employees.Employee":                 "",
			"Comparisions":                       "",
			"Orientations":                       "",
		},
		"abetProgramOrientations": {},
		"abetProgramLogs": {
			"Employee": "",
		},
		"abetProgramEmployees": {
			"Employee": "",
		},
		"abetProgramComparisions": {},
		"abetPeos":                {},
		"abetPsos": {
			"PerformanceIndicators": "",
			"Peos":                  "",
			"Keyword":               "",
		},
		"abetPerformanceIndicators": {
			"Keyword": "",
		},
		"abetProgramConstituents": {
			"Peos": "",
		},
		"abetRelationshipConstituentAndPeos": {},
		"abetRelationshipCourseAndPrograms": {
			"Course":                "",
			"Course.RelatedCourses": "",
			"Orientations":          "",
			"PSOs.PSO":              "",
			// "PerformanceIndicators.PerformanceIndicator": "",
		},
		"abetRelationshipProgramCourseAndOrientations": {},
		"abetRelationshipProgramCourseAndSos":          {},
		"abetRelationshipProgramCourseSoAndLos":        {},
		"abetRelationshipProgramCourseAndPis":          {},
		"abetRelationshipProgramCoursePiAndLos":        {},
		"abetRelationshipPsoAndPeos":                   {},
		"abetRelationshipSchoolYearAndPrograms": {
			"SchoolYear": "",
		},
		"abetOutlines": {
			"Course.CourseType":        "",
			"Course.RelatedCourses":    "",
			"Course.Teachers":          "",
			"Course.Teachers.Employee": "",
			"Textbooks.Document":       "",
			"LearningOutcomes.PSOs.AbetRelationshipProgramCourseAndSO.PSO.Keyword": "",
			"LearningOutcomes.Keyword": "",
			"LearningOutcomes.PIs":     "",
			// "LearningOutcomes.PIs.AbetRelationshipProgramCourseAndPI.PerformanceIndicator.Pso": "",
			"AssessmentPosts":                  "",
			"AssessmentPosts.Rubrics":          "",
			"AssessmentPosts.LearningOutcomes": "",
			"Topics":                           "",
			"TeachingPlans.Documents":          "",
			"TeachingPlans.LearningOutcomes":   "",
			"TeachingPlans.AssessmentPosts":    "",
			"TeachingPlans.Topics":             "",
			"Rubrics.RubricItems":              "",
			"Logs.Employee":                    "",
			"UpdateProcesses":                  "",
		},
		"abetOutlineLogs": {
			"Employee": "",
		},
		"abetOutlineUpdateProcesses": {
			"Employee": "",
		},
		"abetOutlineLearningOutcomes": {
			"Keyword": "",
			// "PIs.AbetRelationshipProgramCourseAndPI.PerformanceIndicator": "",
			"PSOs.AbetRelationshipProgramCourseAndSO.PSO.Keyword": "",
		},
		"abetOutlineDocuments": {
			"Document": "",
		},
		"abetOutlineRubrics": {
			"RubricItems": "",
		},
		"abetOutlineRubricItems": {},
		"abetOutlineAssessmentPosts": {
			"AssessmentForms":  "", // .ExamType
			"Rubrics":          "",
			"LearningOutcomes": "",
		},
		"abetOutlineTopics": {},
		"abetOutlineTeachingPlans": {
			"Documents":        "",
			"LearningOutcomes": "",
			"AssessmentPosts":  "",
			"Topics":           "",
		},
		"abetRelationshipTeachingPlanAndDocuments":          {},
		"abetRelationshipTeachingPlanAndLearningOutcomes":   {},
		"abetRelationshipTeachingPlanAndAssessmentPosts":    {},
		"abetRelationshipTeachingPlanAndTopics":             {},
		"abetRelationshipAssessmentPostAndLearningOutcomes": {},
		"abetRelationshipAssessmentPostAndExamTypes":        {},
		"abetRelationshipAssessmentPostAndRubrics":          {},
		"faculties": {
			"Subjects": "",
		},
		"majors": {
			"Faculty":    "",
			"MajorGroup": "",
		},
		"courses": {
			"CourseType":                         "",
			"RelatedCourses":                     "",
			"Faculty":                            "",
			"Teachers":                           "",
			"Teachers.Employee":                  "",
			"Subject":                            "",
			"EducationPrograms":                  "",
			"EducationPrograms.EducationProgram": "",
			"EducationPrograms.EducationProgram.Major":            "",
			"EducationPrograms.EducationProgram.EducationTargets": "",
		},
		"bloomWords": {
			"Groups.Levels.Keywords": "",
		},
		"employees": {
			"User":           "",
			"User.UserRoles": "",
			// "Specialties":      "",
		},
		"externalAssessmentTeams": {
			"User":           "",
			"User.UserRoles": "",
		},
		"educationPrograms": {
			"EducationTargets":                                  "",
			"EducationStandardOutputs":                          "",
			"EducationStandardOutputs.Keyword.Level.Group.Area": "",
			"Courses":                       "",
			"Courses.Course":                "",
			"Courses.Course.CourseType":     "",
			"Courses.Course.RelatedCourses": "",
			"Courses.Course.Faculty":        "",
			"Courses.Course.Subject":        "",
			"ITUTables":                     "",
			"Major":                         "",
			"Major.Faculty":                 "",
			"Major.MajorGroup":              "",
			"SchoolYears":                   "",
			"SchoolYears.SchoolYear":        "",
			"Orientations":                  "",
			"Employees.Employee":            "",
			"Logs":                          "",
			"Logs.Employee":                 "",
		},
		"educationLogs": {
			"Employee": "",
		},
		"courseEducationPrograms": {},
		"cdioProgramEmployees": {
			"Employee": "",
		},
		"outlines": {
			"Documents":                              "",
			"Documents.Document":                     "",
			"CourseTargets":                          "",
			"CourseTargets.EducationStandardOutputs": "",
			"CourseTargets.EducationStandardOutputs.EducationStandardOutput":                 "",
			"CourseTargets.EducationStandardOutputs.EducationStandardOutput.Keyword":         "",
			"CourseTargets.EducationStandardOutputs.EducationStandardOutput.EducationTarget": "",
			"CourseTargets.Keyword": "",
			"CourseStandardOutputs.EducationStandardOutputs.EducationStandardOutput":                 "",
			"CourseStandardOutputs.EducationStandardOutputs.EducationStandardOutput.Keyword":         "",
			"CourseStandardOutputs.EducationStandardOutputs.EducationStandardOutput.EducationTarget": "",
			"CourseStandardOutputs.Keyword":                      "",
			"Rubrics":                                            "",
			"Rubrics.RubricItems":                                "",
			"TeachingPlans":                                      "",
			"TeachingPlans.Documents":                            "",
			"TeachingPlans.Documents.OutlineDocument":            "",
			"TeachingPlans.StandardOutputs.CourseStandardOutput": "",
			// "TeachingPlans.StandardOutputs.CourseStandardOutput": "",
			"TeachingPlans.AssessmentPosts": "",
			// "TeachingPlans.AssessmentPosts.ResultEvaluate":        "",
			"ResultEvaluates.Rubrics.Rubric":                      "",
			"ResultEvaluates.PostOutputs":                         "",
			"ResultEvaluates.EvaluateForms":                       "",
			"ResultEvaluates.EvaluateForms.ExamType":              "",
			"Logs.Employee":                                       "",
			"Course":                                              "",
			"Course.CourseType":                                   "",
			"Course.RelatedCourses":                               "",
			"Course.Faculty":                                      "",
			"Course.Faculty.Majors":                               "",
			"Course.Faculty.Subjects":                             "",
			"Course.Teachers":                                     "",
			"Course.Teachers.Employee":                            "",
			"Course.EducationPrograms":                            "", // .Courses
			"Course.EducationPrograms.EducationProgram":           "",
			"Course.EducationPrograms.EducationProgram.Courses":   "",
			"Course.EducationPrograms.EducationProgram.ITUTables": "",
			"Course.EducationPrograms.EducationProgram.EducationStandardOutputs": "",
			"Course.EducationPrograms.EducationProgram.EducationTargets":         "",
			"Course.EducationPrograms.EducationProgram.Orientations":             "",
			"Course.Subject": "",
			// "Course.Specialties":                        "",
			"OutlineUpdateProcesses.Employee": "",
		},
		"teachingPlans": {
			"Documents":                 "",
			"Documents.OutlineDocument": "",
		},
		"teachingPlanDocuments": {
			"OutlineDocument": "",
		},
		"loeStudents": {},
		"loeCourseClasses": {
			"Students": "",
		},
		"loeCourseClassStudents": {},
		"loeExamPoints":          {},
	}

	MAP_FOREIGN_KEY_DELETE = map[string]map[string]string{
		"users": {
			"employee": "user_id",
		},
	}
	MAP_ENUM_MODEL_MAP = map[string]map[string][]string{
		"students": {
			"Gender": {"", "F", "M", "O"},
		},

		"profiles": {
			"Gender":     {"", "female", "male", "other"},
			"Academic":   {"", "pgs", "gs"},
			"Degree":     {"", "ts", "ths", "bsck1", "bsck2", "cd", "dh", "dsck1", "ptth", "tc", "tskh"},
			"WorkStatus": {"", "work", "contractTerminate", "stop", "maternityLeave"},
			// "Ethnic":   {"kinh", "không", "mường", "tày"},
			// "Religion": {"phật giáo", "không", "thiên chúa giáo"},
		},
	}
	MAP_OUTLINE_STATUS_NUM = map[string]int{
		"REJECTED":          0,
		"WRITING":           1,
		"DONE":              2,
		"SECTION_APPROVED":  3,
		"DEAN_APPROVED":     4,
		"TRAINING_APPROVED": 5,
		"ADMIN_APPROVED":    6,
	}
	ListDefaultValue = []string{"0", "", "false"}

	SyncLogRunning  = "Running"
	SyncLogFinished = "Finished"
	SyncLogStopped  = "Stopped"
	SyncLogHalted   = "Halted"
)

type DeleteListId struct {
	ID []int `json:"id"`
}
