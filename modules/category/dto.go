package category

import "time"

type DTOCategory struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	ModifiedAt  time.Time `db:"modified_at"`
	ModifiedBy  string    `db:"modified_by"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedBy   string    `db:"created_by"`
}
