package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	result := u.db.Create(&session)
	return result.Error // TODO: replace this
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	result := u.db.Where("token = ?", tokenTarget).Delete(&model.Session{})
	return result.Error // TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	result := u.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session)
	return result.Error // TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	session := model.Session{} // TODO: replace this
	session, e := u.SessionAvailToken(token)
	if e != nil {
		return model.Session{}, e
	}
	if u.TokenExpired(session) {
		err := u.DeleteSessions(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return session, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	results := model.Session{}
	result := u.db.First(&results, "Username = ?", name)
	if result.Error != nil {
		return model.Session{}, result.Error
	}
	return results, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	results := model.Session{}
	result := u.db.First(&results, "token = ?", token)
	if result.Error != nil {
		return model.Session{}, result.Error
	}
	return results, nil // TODO: replace this
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
