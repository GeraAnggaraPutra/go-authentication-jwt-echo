package repository

import (
	"database/sql"
	"go-auth-jwt/entity"
)

type Repository interface {
	GetAll() ([]*entity.Biodata, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]*entity.Biodata, error) {
	biodatas := []*entity.Biodata{}

	sqlQuery := "SELECT * FROM biodata"
	rows, err := r.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		biodata := &entity.Biodata{}
		var createdAt, updatedAt sql.NullTime
		err := rows.Scan(&biodata.ID, &biodata.NAME, &biodata.AGE, &biodata.ADDRESS, &createdAt, &updatedAt)
		if err != nil {
			return nil, err
		}
		if createdAt.Valid {
			biodata.CreatedAt = createdAt.Time
		}
		if updatedAt.Valid {
			biodata.UpdatedAt = updatedAt.Time
		}
		biodatas = append(biodatas, biodata)
	}

	return biodatas, nil
}
