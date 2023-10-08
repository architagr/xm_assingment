package models

import "golang.org/x/crypto/bcrypt"

type PasswordModel string

func (pass *PasswordModel) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	*pass = PasswordModel(string(bytes))
	return nil
}

func (pass *PasswordModel) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*pass), []byte(password))
	return err == nil
}
