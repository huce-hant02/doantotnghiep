package model

type Api struct {
	Id     int    `json:"id" gorm:"column:id"`
	URL    string `json:"url" gorm:"column:url"`
	Method string `json:"method" gorm:"column:method"`
}

type Author struct {
	ID          int    `json:"id" gorm:"column:id"`
	Role        string `json:"role" gorm:"column:role"`
	Description string `json:"description" gorm:"description"`
}

type Author_api struct {
	ID       int    `json:"id" gorm:"column:id"`
	AuthorId int    `json:"authorId" gorm:"column:author_id"`
	ApiId    int    `json:"apiId" gorm:"column:api_id"`
	Value    string `json:"value" gorm:"column:value"`
}

type Policy struct {
	ID     int    `json:"id"`
	ApiId  int    `json:"apiId" gorm:"column:api_id"`
	URL    string `json:"url" gorm:"column:url"`
	Method string `json:"method" gorm:"column:method"`

	AuthorId int    `json:"authorId" gorm:"column:author_id"`
	Role     string `json:"role" gorm:"column:role"`
	Value    string `json:"value" gorm:"column:value"`
}

type AuthorRepository interface {
	AddNewApi(URL, Method string) error
	GetAllApi(page int, pageSize int) (apis []Api, total int64, err error)
	DeleteApi(idAPi int) error
	UpdateApi(id int, newAPi Api) error

	AddNewAuthor(roleName, desc string) error
	GetAllAuthor(page int, pageSize int) (apis []Author, total int64, err error)
	DeleteAuthor(id int) error
	UpdateAuthor(id int, newAuthor Author) error

	AddAuthor_api(newPolicy Policy) error
	GetAuthor_api(page int, pageSize int) (apis []Author_api, total int64, err error)
	DeleteAuthor_api(id int) error
	DeleteAuthor_api_byauthorandapi(authorId int, apiId int) error
	UpdateAuthor_api(id int, newAuthor_api Author_api) error

	GetPolicyById(authorApiID int) (*Policy, error)
	FilterPolicy(authorApiID, authorID, ApiID int, Role, URL, Method, Value string) ([]Policy, error)
}
