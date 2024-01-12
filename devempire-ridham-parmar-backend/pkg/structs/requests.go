package structs

import "time"

// All request sturcts
// Request struct have Req prefix

type ReqRegisterUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Roles     string `json:"roles" validate:"required"`
}

type ReqLoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqCreateApp struct {
	Id             int        `json:"-"`
	AppName        string     `json:"app" validate:"required"`
	Category       string     `json:"category" validate:"required"`
	Size           string     `json:"size" validate:"required"`
	Rating         string     `json:"rating" validate:"required"`
	Reviews        string     `json:"reviews" validate:"required"`
	Types          string     `json:"type" validate:"required"`
	Installs       string     `json:"installs" validate:"required"`
	Price          string     `json:"price" validate:"required"`
	ContentRating  string     `json:"contentRating" validate:"required"`
	Generes        string     `json:"generes" validate:"required"`
	LastUpdated    string     `json:"lastUpdated,omitempty"`
	CurrentVersion string     `json:"currentVer" validate:"required"`
	AndroidVersion string     `json:"androidVer" validate:"required"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `json:"-"`
}

type ReqCreateReviews struct {
	TranslatedReview      string `json:"translatedReview" validate:"required"`
	Sentiment             string `json:"sentiment" validate:"required"`
	SentimentPolarity     string `json:"-"`
	SentimentSubjectivity string `json:"-"`
}
