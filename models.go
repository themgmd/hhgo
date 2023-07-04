package hhgo

import (
	"github.com/guregu/null"
	"time"
)

const ApiEndpoint = "https://api.hh.ru"

// Options - Опции для создания инстанса класса HeadHunter
//
// PublicId - публичный ключ с https://dev.hh.ru
// PrivateKey - приватный ключ с https://dev.hh.ru
// Endpoint - Endpoint для запросов
type Options struct {
	PublicId   string
	PrivateKey string
	Endpoint   string
}

type BaseResponse struct {
	Status     int
	StatusText string
	Data       interface{}
}

type ResponseSchedule struct {
	Id   Schedule
	Name string
}

type ResponseVacancyData struct {
	Items   []VacancyType
	Found   int
	Pages   int
	PerPage int
	Page    int
}

type ExtendedVacancyData struct {
	ResponseVacancyData
	Clusters       interface{}
	Arguments      interface{}
	AlternativeURL string
}

type Department struct {
	Id   string
	Name string
}

type Area struct {
	Id   string
	Name string
	URL  string
}

type Salary struct {
	From     float64
	To       float64
	Currency Currency
	Gross    bool
}

type Address struct {
	City          null.String
	Street        null.String
	Building      null.String
	Description   null.String
	Lat           float64
	Lng           float64
	Raw           interface{}
	Metro         Metro
	MetroStations []Metro
	Id            string
}

type Metro struct {
	StationName string
	LineName    string
	StationId   string
	LineId      string
	Lat         float64
	Lng         float64
}

type VacancyType struct {
	Id   string
	Name string
}

type Employer struct {
	Id           string
	Name         string
	URL          string
	AlternateURL string
	LogoURLs     interface{}
	VacanciesURL string
	Trusted      bool
}

type Snippet struct {
	Requirement    string
	Responsibility string
}

type Counters struct {
	Responses int
}

type Vacancy struct {
	ID                     string
	Premium                bool
	Name                   string
	Department             *Department
	HasTest                bool
	ResponseLetterRequired bool
	Area                   Area
	Salary                 Salary
	Type                   VacancyType
	Address                null.String
	ResponseUrl            null.String
	SortPointDistance      null.String
	Archived               bool
	CreatedAt              time.Time
	PublishedAt            time.Time
	ApplyAlternateUrl      string
	Relations              []string
	Employer               Employer
	Snippet                Snippet
	Counters               Counters
	WorkingDays            interface{}
	WorkingTimeIntervals   interface{}
	WorkingTimeModels      interface{}
	AcceptTemporary        bool
}
