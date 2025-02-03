package models

type Tags struct {
	List []string `json:"list"`
}

type User struct {
	UserId          int    `db:"user_id" json:"userId"`
	Name            string `db:"name" json:"name"`
	Phone           string `db:"phone" json:"phone"`
	Email           string `db: "email" json:"email"`
	Password        string `db: "passsword" json:"password"`
	Gender          string `db: "gender" json:"gender"`
	City            string `db:"city" json:"city"`
	Role            string `db:"role" json:"role"`
	RequiredVacancy int    `db:"required_vacancy" json:"required_vacancy"`
	Tags            Tags   `db:"tags" json:"tags:`
}
