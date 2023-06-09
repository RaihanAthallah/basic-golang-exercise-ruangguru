package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	s.db.Create(&session)
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	s.db.Where("token = ?", token).Delete(&model.Session{})
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	s.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(&model.Session{
		Username: session.Username,
		Token:    session.Token,
		Expiry:   session.Expiry,
	})
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	result := s.db.Where("username = ?", name).First(&model.Session{})
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var result model.Session
	err := s.db.Where("token = ?", token).First(&model.Session{}).Scan(&result)
	if err.Error != nil {
		return model.Session{}, err.Error
	}

	return result, nil // TODO: replace this
}
