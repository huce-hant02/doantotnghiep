package model

import (
	"time"

	"github.com/lib/pq"
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type EmailInfo struct {
	Email string     `json:"email"`
	Error ErrorLogin `json:"error"`
}

type ErrorLogin struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type OutlineStatistic struct {
	NumOfInit              int `json:"numOfInit" gorm:"num_of_init"`
	NumOfWriting           int `json:"numOfWriting" gorm:"num_of_writing"`
	NumOfDone              int `json:"numOfDone" gorm:"num_of_done"`
	NumOfTrainingApproved  int `json:"numOfTrainingApproved" gorm:"num_of_training_approved"`
	NumOfDeanApproved      int `json:"numOfDeanApproved" gorm:"num_of_dean_approved"`
	NumOfSectionApproved   int `json:"numOfSectionApproved" gorm:"num_of_section_approved"`
	NumOfAdminApproved     int `json:"numOfAdminApproved" gorm:"num_of_admin_approved"`
	NumOfCourse            int `json:"numOfCourse" gorm:"num_of_course"`
	NumOfProgram           int `json:"numOfProgram" gorm:"num_of_program"`
	NumOfMajor             int `json:"numOfMajor" gorm:"num_of_major"`
	NumOfTeacher           int `json:"numOfTeacher" gorm:"num_of_teacher"`
	NumOfDocument          int `json:"numOfDocument" gorm:"num_of_document"`
	NumOfDocumentInLibrary int `json:"numOfDocumentInLibrary" gorm:"num_of_document_in_library"`
	NumOfFaculty           int `json:"numOfFaculty" gorm:"num_of_faculty"`
	NumOfRejected          int `json:"numOfRejected" gorm:"num_of_rejected"`
}
type LoginLDAP struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

type UserPermissonBoolean struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}
type UserPemissionModelType struct {
	Create bool                            `json:"create"`
	Read   bool                            `json:"read"`
	Update bool                            `json:"update"`
	Delete bool                            `json:"delete"`
	Fields map[string]UserPermissonBoolean `json:"fields"`
}

type UserPermissionAPI struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type UserPermissionRole struct {
	ID   uint   `json:"id" gorm:"-"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type UserPermission struct {
	UserActive *bool                             `json:"userActive"`
	ActiveRole string                            `json:"activeRole"`
	Roles      []UserPermissionRole              `json:"roles"`
	Routes     []string                          `json:"routes"`
	APIs       []UserPermissionAPI               `json:"apis"`
	ModelTypes map[string]UserPemissionModelType `json:"modelTypes"`
}

type ActivatedUserRole struct {
	UserActive *bool                `json:"userActive"`
	ActiveRole string               `json:"activeRole"`
	Roles      []UserPermissionRole `json:"roles"`
}

type UserResetRole struct {
	Roles []string `json:"roles"`
}

type UserSelectRole struct {
	Role string `json:"role"`
}

type UserSelectRoleResponse struct {
	Token      TokenDetail `json:"token"`
	ActiveRole string      `json:"activeRole"`
}

// TokenLoadResponse response for login and refresh api
type TokenLoadResponse struct {
	AccessToken            string                  `json:"access_token"`
	RefreshToken           string                  `json:"refresh_token"`
	ActiveRole             string                  `json:"activeRole"` // Vai trò hiện tại đang chọn trên hệ thống
	Permission             UserPermission          `json:"permission"` // Quyền người dùng dựa trên vai trò
	Role                   string                  `json:"role"`
	UserID                 uint                    `json:"userId"`
	User                   *User                   `json:"user"`
	EmployeeInfo           *Employee               `json:"employeeInfo"`
	DepartmentInfo         *Faculty                `json:"departmentInfo"`
	ExternalAssessmentTeam *ExternalAssessmentTeam `json:"teamInfo"`
}

type Sort struct {
	Key   string `json:"key"`   // field name
	Value string `json:"value"` // asc or desc
}

type TimeFilter struct {
	ColumnName string `json:"columnName"`
	StartAt    string `json:"startAt"`
	EndAt      string `json:"endAt"`
}

type TimePayload struct {
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
}

type SubmitProgramPayload struct {
	ProgramId uint       `json:"programId"`
	Status    *string    `json:"status"`
	Note      string     `json:"note"`
	Deadline  *time.Time `json:"deadline"`
}

type SubmitOutlinePayload struct {
	OutlineId uint       `json:"outlineId"`
	Status    *string    `json:"status"`
	Note      string     `json:"note"`
	Deadline  *time.Time `json:"deadline"`
}

type ValidationError struct {
	ID    uint                   `json:"id"`
	Error map[string]interface{} `json:"error"`
}

type FilterGeneralPayloadHRM struct {
	ModelType         string         `json:"modelType"`
	Filter            string         `json:"filter"`
	Search            string         `json:"search"`
	Page              int            `json:"page"`
	PageSize          int            `json:"pageSize"`
	IsPaginateDB      bool           `json:"isPaginateDB"`
	IgnoreAssociation bool           `json:"ignoreAssociation"`
	IsExactSearch     bool           `json:"isExactSearch"`
	IsUnscoped        bool           `json:"isUnscoped"`
	SelectColumn      pq.StringArray `json:"selectColumn"`
	JoinColumn        pq.StringArray `json:"joinColumn"`
}

type ReportOutlinePayload struct {
	ScholasticSemesterId int    `json:"scholasticSemesterId"`
	Semester             int    `json:"semester"`
	Year                 int    `json:"year"`
	SchoolYear           string `json:"schoolYear"` // ex: K15
	ByEducation          bool   `json:"byEducation"`
	ByFrame              bool   `json:"byFrame"` // phân loại theo khối kiến thức
}

type ReportUsedDocumentPayload struct {
	ScholasticSemesterId int  `json:"scholasticSemesterId"`
	ByEducation          bool `json:"byEducation"`
	ByCourse             bool `json:"byCourse"`
}

type DocumentErr struct {
	StartAt        time.Time `json:"startAt"`
	EndAt          time.Time `json:"endAt"`
	Service        string    `json:"service"`
	NumOfDocuments int       `json:"numOfDocuments"`
	Error          string    `json:"error"`
}

type ScholasticCopyInfo struct {
	SourceID int      `json:"sourceId"`
	TargetID int      `json:"targetId"`
	Model    string   `json:"model"`
	ListCode []string `json:"listCode"`
	ListID   []int    `json:"listId"`
}

type SyncStatusPayload struct {
	ApiID       uint   `json:"apiId"`
	Url         string `json:"url"`
	Method      string `json:"method"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsRunning   *bool  `json:"isRunning"`

	DefaultParams  *string `json:"defaultParams"`
	DefaultPayload *string `json:"defaultPayload"`
}
