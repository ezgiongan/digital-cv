package db

import (
	"context"
	"digital-cv-backend/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresCVRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresCVRepository(pool *pgxpool.Pool) *PostgresCVRepository {
	return &PostgresCVRepository{pool: pool}
}

func (r *PostgresCVRepository) GetCV() (*models.CV, error) {
	ctx := context.Background()

	var cv models.CV

	// cv
	err := r.pool.QueryRow(ctx,
    `SELECT id, name, title, email, phone, linkedin, about FROM cv LIMIT 1`,
).Scan(
    &cv.ID,
    &cv.Name,
    &cv.Title,
    &cv.Email,
    &cv.Phone,
    &cv.LinkedIn,
    &cv.AboutMe,
)
	if err != nil {
		return nil, err
	}

	// interests
	interestRows, err := r.pool.Query(ctx,
		`SELECT id, name, description FROM interests WHERE cv_id=$1`,
		cv.ID,
	)
	if err != nil {
		return nil, err
	}
	defer interestRows.Close()

	for interestRows.Next() {
		var interest models.Interest
		if err := interestRows.Scan(
			&interest.ID,
			&interest.Name,
			&interest.Description,
		); err != nil {
			return nil, err
		}
		cv.Interests = append(cv.Interests, interest)
	}

	// education
	educationRows, err := r.pool.Query(ctx,
		`SELECT id, institution, degree, start_year, end_year, description 
		 FROM education WHERE cv_id=$1`,
		cv.ID,
	)
	if err != nil {
		return nil, err
	}
	defer educationRows.Close()

	for educationRows.Next() {
		var education models.Education
		if err := educationRows.Scan(
			&education.ID,
			&education.Institution,
			&education.Degree,
			&education.StartYear,
			&education.EndYear,
			&education.Description,
		); err != nil {
			return nil, err
		}
		cv.Education = append(cv.Education, education)
	}

	// volunteering
	volRows, err := r.pool.Query(ctx,
		`SELECT id, organization, role, start_date, end_date, description 
		 FROM volunteering WHERE cv_id=$1`,
		cv.ID,
	)
	if err != nil {
		return nil, err
	}
	defer volRows.Close()

	for volRows.Next() {
		var volunteering models.Volunteering
		if err := volRows.Scan(
			&volunteering.ID,
			&volunteering.Organization,
			&volunteering.Role,
			&volunteering.StartDate,
			&volunteering.EndDate,
			&volunteering.Description,
		); err != nil {
			return nil, err
		}
		cv.Volunteering = append(cv.Volunteering, volunteering)
	}

	// skills
	skillRows, err := r.pool.Query(ctx,
		`SELECT id, name FROM skills WHERE cv_id=$1`,
		cv.ID,
	)
	if err != nil {
		return nil, err
	}
	defer skillRows.Close()

	for skillRows.Next() {
		var skill models.Skill
		if err := skillRows.Scan(
			&skill.ID,
			&skill.Name,
		); err != nil {
			return nil, err
		}
		cv.Skills = append(cv.Skills, skill)
	}

	// experiences
	expRows, err := r.pool.Query(ctx,
		`SELECT id, company, position, start_date, end_date, description
		 FROM experience WHERE cv_id=$1`,
		cv.ID,
	)
	if err != nil {
		return nil, err
	}
	defer expRows.Close()

	for expRows.Next() {
		var exp models.Experience
		if err := expRows.Scan(
			&exp.ID,
			&exp.Company,
			&exp.Position,
			&exp.StartDate,
			&exp.EndDate,
			&exp.Description,
		); err != nil {
			return nil, err
		}
		cv.Experiences = append(cv.Experiences, exp)
	}

	return &cv, nil
}
