package models

import "time"

type Question struct {
    ID               uint        `gorm:"primaryKey"`
    SurveyID         uint
    Survey           *Survey     `gorm:"foreignKey:SurveyID"`
    Text             string      `gorm:"not null"`
    Type             string      `gorm:"not null"`
    Order            int
    AttachmentURL    string
    CorrectOptionIDs string      `gorm:"type:text"`
    IsConditional    bool
    Condition        string      `gorm:"type:text"`
    Options          []*Option   `gorm:"foreignKey:QuestionID"`
    CreatedAt        time.Time
    UpdatedAt        time.Time
}