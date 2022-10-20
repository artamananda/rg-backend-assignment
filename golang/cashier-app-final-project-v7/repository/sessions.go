package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	// Select target token and delete from listSessions
	for key, val := range listSessions {
		if tokenTarget == val.Token {
			listSessions = RemoveIndex(listSessions, key)
			break
		}
	}
	// TODO: answer here

	jsonData, err := json.Marshal(listSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}

	return err
}

func RemoveIndex(s []model.Session, index int) []model.Session {
	ret := []model.Session{}
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	listSessions, _ := u.ReadSessions()
	listSessions = append(listSessions, session)
	jsonData, err := json.Marshal(listSessions)
	if err != nil {
		panic(err)
	}
	err = u.db.Save("sessions", jsonData)
	if err != nil {
		panic(err)
	}
	return err // TODO: replace this
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	res, err := u.TokenExist(token)
	if err != nil {
		return res, err
	}
	if u.TokenExpired(res) {
		err = fmt.Errorf("Token is Expired!")
		u.DeleteSessions(token)
		res = model.Session{}
	}
	return res, err // TODO: replace this
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
