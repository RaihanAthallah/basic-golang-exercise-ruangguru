package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
)

type TeacherRepository interface {
	FetchByID(id int) (*model.Teacher, error)
	Store(g *model.Teacher) error
	Delete(id int) error
}

type teacherRepoImpl struct {
	db *sql.DB
}

func NewTeacherRepo(db *sql.DB) *teacherRepoImpl {
	return &teacherRepoImpl{db}
}

func (g *teacherRepoImpl) FetchByID(id int) (*model.Teacher, error) {
	row := g.db.QueryRow("SELECT id, name, address, subject FROM teachers WHERE id = $1", id)

	var Teacher model.Teacher
	err := row.Scan(&Teacher.ID, &Teacher.Name, &Teacher.Address, &Teacher.Subject)
	if err != nil {
		return nil, err
	}

	return &Teacher, nil
}

func (g *teacherRepoImpl) Store(teacher *model.Teacher) error {
	_, err := g.db.Exec("INSERT INTO teachers(name, address, subject) VALUES($1, $2, $3) RETURNING id", teacher.Name, teacher.Address, teacher.Subject)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (g *teacherRepoImpl) Delete(id int) error {
	_, err := g.db.Exec("DELETE FROM teachers WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}
