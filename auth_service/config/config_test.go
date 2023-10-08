package config

import (
	"strings"
	"testing"
)

func TestInitFileConfig(t *testing.T) {

	t.Run("Get error if file is not present", func(tc *testing.T) {
		filePath := "./config/nonamed.json"
		_, err := InitConfigFromFile(filePath)
		if err == nil || !strings.Contains(err.Error(), "is not found") {
			tc.Errorf("Should have got error for file not present, but got %+v\n", err)
		}
	})
	t.Run("Get error file is not haveing json data", func(tc *testing.T) {
		filePath := "./wrong_config.json"
		_, err := InitConfigFromFile(filePath)
		if err == nil || !strings.Contains(err.Error(), "File Unmarshalling error") {
			tc.Errorf("Should have got error for file not having data, but got %+v\n", err)
		}
	})

	t.Run("Get config", func(tc *testing.T) {
		filePath := "./valid_config.json"
		con, err := InitConfigFromFile(filePath)
		if err != nil {
			tc.Errorf("Should not have got error, but got err%+v\n", err)
		}
		if con.GetPortNumber() != 8080 {
			tc.Errorf("Expected port number 8080 but got %d\n", con.GetPortNumber())
		}
		if con.GetTokenExpiration() != 180 {
			tc.Errorf("Expected token expiration 180 but got %d\n", con.GetTokenExpiration())
		}
		if con.GetDbConnectionString() != "connectionString" {
			tc.Errorf("Expected token expiration connectionString but got %s\n", con.GetDbConnectionString())
		}
	})
}
