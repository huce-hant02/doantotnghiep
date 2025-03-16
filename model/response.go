package model

import (
	"time"

	"github.com/lib/pq"
)

type OutlineReport struct {
	Type               string   `json:"type"` // ABET | CDIO
	Code               string   `json:"code"`
	CourseID           uint     `json:"courseID"`
	CourseCode         string   `json:"courseCode"`
	CourseName         string   `json:"courseName"`
	NumOfCredit        float32  `json:"numOfCredit"`
	NumOfCreditLt      float32  `json:"numOfCreditLt"`
	NumOfCreditTh      float32  `json:"numOfCreditTh"`
	Semester           []string `json:"semester"`
	Major              []string `json:"major"`
	MajorGroup         []string `json:"majorGroup"`
	SchoolYear         []string `json:"schoolYear"`
	Faculty            string   `json:"faculty"`
	PrerequisiteCourse []string `json:"prerequisiteCourse"`
	FirstlyCourse      []string `json:"firstlyCourse"`
	ParallelCourse     []string `json:"parallelCourse"`
	EvaluateForm       string   `json:"evaluateForm"`
	ExamTime           int      `json:"examTime"`
	NumOfTestLt        float32  `json:"numOfTestLt"`
	NumOfTestTh        float32  `json:"numOfTestTh"`
	Status             string   `json:"status"`
	Teachers           []string `json:"teachers"`
	ScoringWeight      []string `json:"scoringWeight"`

	Programs []OutlineProgramReport `json:"programs" gorm:"type:text"`

	/* For QLDTbeta */
	AssessmentPosts string `json:"assessmentPosts"` // Công thức điểm
}

type DocumentReport struct {
	Code          string         `json:"code"`                      // Mã
	Type          string         `json:"type"`                      // Phân loại GT chính/ Tham khảo
	DocumentType  string         `json:"documentType"`              // Phân loại trên thư viện
	Title         string         `json:"title"`                     // Tiêu đề
	Author        string         `json:"author"`                    // Tác giả / dịch giả
	Year          string         `json:"year"`                      // Năm xuất bản
	Publisher     string         `json:"publisher"`                 // Nhà xuất bản
	Edition       string         `json:"edition"`                   // Phiên bản
	Isbn          string         `json:"isbn"`                      // ISBN
	Issn          string         `json:"issn"`                      // ISSN
	Link          string         `json:"link"`                      // Link
	Doi           string         `json:"doi"`                       // DOI
	Source        string         `json:"source"`                    // Nguồn
	NumOfItems    int            `json:"numOfItems"`                // Số bản ghi hiện có
	DocumentNote  string         `json:"documentNote"`              // Ghi chú
	OutlineCode   string         `json:"outlineCode"`               // Mã Đề cương
	OutlineStatus string         `json:"outlineStatus"`             // TT Đề cương
	CourseCode    string         `json:"courseCode"`                // Mã HP
	CourseName    string         `json:"courseName"`                // Tên học phần
	CourseFaculty string         `json:"courseFaculty"`             // Khoa quản lý
	Status        string         `json:"status"`                    // Trạng thái tài liệu [Cần mua?]
	UpdatedAt     *time.Time     `json:"updatedAt"`                 // Thời điểm cập nhật gần nhất
	Programs      pq.StringArray `json:"programs" gorm:"type:text"` // Chương trình

	CourseID uint `json:"courseID"`
}

type OutlineProgramReport struct {
	ProgramCode     string `json:"programCode"`     // Mã CTĐT
	ProgramName     string `json:"programName"`     // Tên CTĐT
	KnowledgeGroup  string `json:"knowledgeGroup"`  // Khối kiến thức
	KnowledgeType   string `json:"knowledgeType"`   // Bắt buộc/tự chọn
	OrientationCode string `json:"orientationCode"` // any
	OrientationName string `json:"orientationName"` // any
	OrientationID   uint   `json:"orientationId"`   // any
	OrientationType string `json:"orientationType"` // any
}

type ConnectCongThucDiem struct {
	MAHP     string `json:"MAHP"`
	TENHP    string `json:"TENHP"`
	CONGTHUC string `json:"CONGTHUC"`
	HEDAOTAO string `json:"HEDAOTAO"`
	NAMHOC   string `json:"NAMHOC"`
	HOCKY    string `json:"HOCKY"`
	DOTHOC   string `json:"DOTHOC"`
	MACTDT   string `json:"MACTDT"`
	KHOAHOC  string `json:"KHOAHOC"`
}

type ConnectListCourse struct {
	MAHP            string  `json:"MAHP"`
	TENHP           string  `json:"TENHP"`
	TONGSTC         float32 `json:"TONGSTC"`
	TONGSTCLT       float32 `json:"TONGSTCLT"`
	TONGSTCTH       float32 `json:"TONGSTCTH"`
	HOCKY           string  `json:"HOCKY"`
	MAKHOIKIENTHUC  string  `json:"MAKHOIKIENTHUC"`
	MACTDT          string  `json:"MACTDT"`
	MAKHOADAOTAO    string  `json:"MAKHOADAOTAO"`
	TENCTDT         string  `json:"TENCTDT"`
	KHOACHUYENMON   string  `json:"KHOACHUYENMON"`
	HEDAOTAO        string  `json:"HEDAOTAO"`
	TENKHOIKIENTHUC string  `json:"TENKHOIKIENTHUC"`
	PHANLOAI        string  `json:"PHANLOAI"`
	MADINHHUONG     string  `json:"MADINHHUONG"`
	TENDINHHUONG    string  `json:"TENDINHHUONG"`
	// PHANLOAIDINHHUONG string   `json:"PHANLOAIDINHHUONG"`
}

type ConnectListProgramCredit struct {
	MAKHOIKIENTHUC  string `json:"MAKHOIKIENTHUC"`
	MACTDT          string `json:"MACTDT"`
	MAKHOADAOTAO    string `json:"MAKHOADAOTAO"`
	TENCTDT         string `json:"TENCTDT"`
	HEDAOTAO        string `json:"HEDAOTAO"`
	TENKHOIKIENTHUC string `json:"TENKHOIKIENTHUC"`
	PHANLOAI        string `json:"PHANLOAI"`
	MADINHHUONG     string `json:"MADINHHUONG"`
	TENDINHHUONG    string `json:"TENDINHHUONG"`

	TONGSTC   float32 `json:"TONGSTC"`
	TONGSTCBB float32 `json:"TONGSTCBB"`
}
