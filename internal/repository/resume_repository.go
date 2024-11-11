package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/aidosgal/ei-jobs-core/internal/model"
)

type ResumeRepository struct {
	db *sql.DB
}

func NewResumeRepository(db *sql.DB) *ResumeRepository {
	return &ResumeRepository{db}
}

func (r *ResumeRepository) GetResumesByUserID(userID int) ([]*model.Resume, error) {
	var resumes []*model.Resume

	query := `
		SELECT id, user_id, date_of_birth, gender, specialization_id, description, salary_from, salary_to, salary_period, created_at
		FROM resumes
		WHERE user_id = ?
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var resume model.Resume
		if err := rows.Scan(&resume.ID, &resume.UserID, &resume.DateOfBirth, &resume.Gender, &resume.SpecializationID, &resume.Description, &resume.SalaryFrom, &resume.SalaryTo, &resume.SalaryPeriod, &resume.CreatedAt); err != nil {
			return nil, err
		}
		resumes = append(resumes, &resume)
	}

	return resumes, nil
}

// GetSkillsByResumeID fetches all skills associated with a specific resume.
func (r *ResumeRepository) GetSkillsByResumeID(resumeID int) ([]*model.ResumeSkill, error) {
	var skills []*model.ResumeSkill

	query := `SELECT id, resume_id, skill FROM resume_skills WHERE resume_id = ?`

	rows, err := r.db.Query(query, resumeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var skill model.ResumeSkill
		if err := rows.Scan(&skill.ID, &skill.ResumeID, &skill.Skill); err != nil {
			return nil, err
		}
		skills = append(skills, &skill)
	}

	return skills, nil
}

func (r *ResumeRepository) CalculateTotalExperience(resumeID int) (string, error) {
	var totalExperience string

	query := `
		SELECT start_month, start_year, end_month, end_year
		FROM resume_organizations
		WHERE resume_id = ?
	`

	rows, err := r.db.Query(query, resumeID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var totalMonths int

	for rows.Next() {
		var startMonth, startYear string
		var endMonth, endYear sql.NullString

		if err := rows.Scan(&startMonth, &startYear, &endMonth, &endYear); err != nil {
			return "", err
		}

		startDate, err := time.Parse("January 2006", startMonth+" "+startYear)
		if err != nil {
			return "", err
		}

		endDate := time.Now()

		if endMonth.Valid && endYear.Valid {
			endDate, err = time.Parse("January 2006", endMonth.String+" "+endYear.String)
			if err != nil {
				return "", err
			}
		}
		diff := endDate.Sub(startDate).Hours() / 24 / 30
		totalMonths += int(diff)
	}

	years := totalMonths / 12
	months := totalMonths % 12
	totalExperience = fmt.Sprintf("%d years %d months", years, months)

	return totalExperience, nil
}