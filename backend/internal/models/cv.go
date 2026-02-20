package models

type CV struct {
	ID		  int        `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	LinkedIn	string `json:"linkedin"`
	AboutMe     string `json:"aboutme"`
	Skills      []Skill      `json:"skills"`
	Experiences []Experience `json:"experiences"`
	Education    []Education    `json:"education"`
	Volunteering []Volunteering `json:"volunteering"`
	Interests    []Interest     `json:"interests"`
}

type Experience struct {
	ID int `json:"id"`
	Company string `json:"company"`
	Position string `json:"position"`
	StartDate string `json:"start_date"`
	EndDate *string `json:"end_date"` //null dönebilir diye belki
	Description *string `json:"description"`
}

type Education struct {
	ID int `json:"id"`
	Institution string `json:"institution"`
	Degree string `json:"degree"`
	StartYear int `json:"start_year"`
	EndYear int `json:"end_year"`
	Description *string `json:"description"`
}


type Skill struct {
	ID int `json:"id"`
	Name string `json:"name"`
	description *string `json:"description"`
}

type Volunteering struct {
	ID int `json:"id"`
	Organization string `json:"organization"`
	Role string `json:"role"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Description *string `json:"description"`
}

type Interest struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description"`
}
	