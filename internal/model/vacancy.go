package model

type Vacancy struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	City         string `json:"city"`
	Country      string `json:"country"`
	CompanyName  string `json:"company_name"`
	SalaryFrom   *int   `json:"salary_from"`
	SalaryTo     *int   `json:"salary_to"`
	SalaryPeriod string `json:"salary_period"`
	CreatedAt    string `json:"created_at"`
}

type OneVacancy struct {
	ID               int                     `json:"id"`
	Title            string                  `json:"title"`
	City             string                  `json:"city"`
	Country          string                  `json:"country"`
	CompanyName      string                  `json:"company_name"`
	SalaryFrom       *int                    `json:"salary_from"`
	SalaryTo         *int                    `json:"salary_to"`
	SalaryPeriod     string                  `json:"salary_period"`
	WorkFormat       string                  `json:"work_format"`
	WorkSchedule     string                  `json:"work_schedule"`
	Conditions       []*VacanyCondition      `json:"coniditions"`
	Requirements     []*VacanyRequirement    `json:"requirements"`
	Responsibilities []*VacanyResponsibility `json:"responsiblities"`
	CreatedAt        string                  `json:"created_at"`
}

type VacanyCondition struct {
	ID        int    `json:"id"`
	VacancyId int    `json:"vacancy_id"`
	Icon      string `json:"icon"`
	Condition string `json:"condition"`
}

type VacanyResponsibility struct {
	ID             int    `json:"id"`
	VacancyId      int    `json:"vacancy_id"`
	Responsibility string `json:"responsibility"`
}

type VacanyRequirement struct {
	ID          int    `json:"id"`
	VacancyId   int    `json:"vacancy_id"`
	Requirement string `json:"requirement"`
}

type VacancyFilters struct {
	SpecializationID int
	Title            string
	City             string
	Country          string
	Salary           *int
}
