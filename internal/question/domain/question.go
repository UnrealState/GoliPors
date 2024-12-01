package domain

import (
	optionDomain "golipors/internal/option/domain"
	surveyDomain "golipors/internal/survey/domain"
	"time"
)

type Question struct {
	ID               uint
	SurveyID         uint
	Survey           *surveyDomain.Survey
	Text             string
	Type             string
	Order            int
	AttachmentURL    string
	CorrectOptionIDs string
	IsConditional    bool
	Condition        string
	Options          []*optionDomain.Option
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
