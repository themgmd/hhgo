package hhgo

// Employment - тип занятости
//
// FullEmployment - Полная занятость
// PartEmployment - Частичная занятость
// ProjectEmployment - Проектная занятость
// VolunteerEmployment - Волонтество
// ProbationEmployment - Стажировки
type Employment string

const (
	FullEmployment      Employment = "full"
	PartEmployment      Employment = "part"
	ProjectEmployment   Employment = "project"
	VolunteerEmployment Employment = "volunteer"
	ProbationEmployment Employment = "probation"
)

// SearchField - область поиска
//
// SearchInName - Поиск в названии вакансии
// SearchInCompanyName - Поиск в названии компании
// SearchInDescription - Поиск в описании вакансии
type SearchField string

const (
	SearchInName        SearchField = "name"
	SearchInCompanyName SearchField = "company_name"
	SearchInDescription SearchField = "description"
)

// Experience - опыт работы
//
// NoExperience - Нет опыта
// Between1And3 - От 1 года до 3 лет
// Between3And6 -  От 3 до 6 лет
// MoreThan6 - Более 6 лет
type Experience string

const (
	NoExperience Experience = "noExperience"
	Between1And3 Experience = "between1And3"
	Between3And6 Experience = "between3And6"
	MoreThan6    Experience = "moreThan6"
)

// Schedule - график работы
//
// FullDaySchedule - Полный день
// ShiftSchedule - Сменный график
// FlexibleSchedule -  Гибкий график
// RemoteSchedule - Удаленная работа
// FlyInFlyOutSchedule - Вахтовый метод
type Schedule string

const (
	FullDaySchedule     Schedule = "fullDay"
	ShiftSchedule       Schedule = "shift"
	FlexibleSchedule    Schedule = "flexible"
	RemoteSchedule      Schedule = "remote"
	FlyInFlyOutSchedule Schedule = "flyInFlyOut"
)

// Currency код валюты
//
// AZN - Азербаджанский манат
// BYR - Белорусский рубль
// EUR - Евро
// GEL - Грузинский лари
// KGS - Киргизский сом
// KZT - Казахский тенге
// RUR - Российский рубль
// UAH - Украинская гривна
// USD - Американский доллар
// UZS - Узбекский сум
type Currency string

const (
	AZN Currency = "AZN"
	BYR Currency = "BYR"
	EUR Currency = "EUR"
	GEL Currency = "GEL"
	KGS Currency = "KGS"
	KZT Currency = "KZT"
	RUR Currency = "RUR"
	UAH Currency = "UAH"
	USD Currency = "USD"
	UZS Currency = "USZ"
)

// Label - метки вакансии
//
// WithAddressLabel - Только с адресом
// AcceptHandicappedLabel - Только доступные для людей с инвалидностью
// NotFromAgencyLabel - Без вакансий агентств
// AcceptKidsLabel - Только доступные для соискателей от 14 лет
type Label string

const (
	WithAddressLabel       Label = "with_address"
	AcceptHandicappedLabel Label = "accept_handicapped"
	NotFromAgencyLabel     Label = "not_from_agency"
	AcceptKidsLabel        Label = "accept_kids"
)

// OrderBy - сортировка
//
// OrderByPublicationTime - По дате
// OrderBySalaryDesc - По убыванию дохода
// OrderBySalaryAsc - По возрастанию дохода
// OrderByRelevance - По соответствию
// OrderByDistance - По удаленности
type OrderBy string

const (
	OrderByPublicationTime OrderBy = "publication_time"
	OrderBySalaryDesc      OrderBy = "salary_desc"
	OrderBySalaryAsc       OrderBy = "salary_asc"
	OrderByRelevance       OrderBy = "relevance"
	OrderByDistance        OrderBy = "distance"
)

// PartTimeWork - вакансии для подработки
//
// OnlySaturdayAndSunday - Работа только по субботам и воскресеньям
// From4To6HoursInDay  - Работа сменами по 4 - 6 часов
// StartAfterSixteen - Работа после 16:00
// AcceptTemporary - Вакансии с временным трудоустройством
// Project - Проектная работа
// Part - Неполный рабочий день
type PartTimeWork string

const (
	OnlySaturdayAndSunday PartTimeWork = "only_saturday_and_sunday"
	From4To6HoursInDay    PartTimeWork = "from_four_to_six_hours_in_a_day"
	StartAfterSixteen     PartTimeWork = "start_after_sixteen"
	AcceptTemporary       PartTimeWork = "accept_temporary"
	Project               PartTimeWork = "project"
	Part                  PartTimeWork = "part"
)
