package model

type StatisticResponse struct {
	Packages map[string]string `json:"packages"`

	CdioOutlineStat OutlineStat `json:"cdioOutline"`
	AbetOutlineStat OutlineStat `json:"abetOutline"`

	CdioProgramStat ProgramStat `json:"cdioProgram"`
	AbetProgramStat ProgramStat `json:"abetProgram"`

	DepartmentStat BasicStat    `json:"department"`
	MajorStat      CourseStat   `json:"major"`
	TeacherStat    BasicStat    `json:"teacher"`
	CourseStat     CourseStat   `json:"course"`
	DocumentStat   DocumentStat `json:"document"`
}

type OutlineStat struct {
	Total      int64             `json:"total"`
	Status     OutlineStatusStat `json:"status"` // Trạng thái biên soạn
	Department map[string]int    `json:"department"`
	Major      map[string]int    `json:"major"`
	Program    map[string]int    `json:"program"`
}

type ProgramStat struct {
	Total      int64          `json:"total"`
	Status     map[string]int `json:"status"` // Trạng thái biên soạn
	Department map[string]int `json:"department"`
	Major      map[string]int `json:"major"`
}

type CourseStat struct {
	Total      int64          `json:"total"`
	Department map[string]int `json:"department"`
}

type DocumentStat struct {
	Total  int64          `json:"total"`
	Source map[string]int `json:"source"`
}

type OutlineStatusStat struct {
	REJECTED         int `json:"REJECTED"`
	WRITING          int `json:"WRITING"`
	DONE             int `json:"DONE"`
	SECTION_APPROVED int `json:"SECTION_APPROVED"`
	DEAN_APPROVED    int `json:"DEAN_APPROVED"`
}

type BasicStat struct {
	Total int64 `json:"total"`
}
