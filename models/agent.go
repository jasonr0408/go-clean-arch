package models

// Administrator 123
type Agent struct {
	UserName string `json:"UserName" gorm:"column:UserName;"`
	Token    string `json:"Token" gorm:"column:Token;"`
}

// TableName 123
func (Agent) TableName() string {
	return "agent"
}
