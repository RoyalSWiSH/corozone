package main

import (
	"fmt"
	"testing"

	//gonfig "github.com/eduardbcom/gonfig"
	"database/sql"
	// "database/sql/driver"
	_ "encoding/json"
)

func init() {
	dbconf := &DBConfig{
		Host:     "db-postgresql-fra1-47535-do-user-1884949-0.a.db.ondigitalocean.com",
		Port:     25060,
		User:     "doadmin",
		Password: "",
		DBName:   "defaultdb",
		SSL:      "require",
	}

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		dbconf.Host, dbconf.Port, dbconf.User, dbconf.Password, dbconf.DBName, dbconf.SSL)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

}
func TestDBGetPushtoken(t *testing.T) {
	token := DBGetPushtoken("rBpyst6rNbTGn7ltJwtDgjrbU2G3")
	var testtoken pushtoken

	testtoken.FirebasePushtoken = "fPpMG6fxW8E:APA91bGi-KDyZ1jZ-I_6piBicdrbxtM6amaruWMIxWgf42AZB0R4AzbeM2fguXp4auGhmYIh9EXOK6UYd-dHolqb1mHH6HslnuK14uLVFaBWF_CHzJLP6HOdDrmayVh8Hc9Zbnk3Lv-T"
	t.Log(token)
	if token != testtoken {
		t.Errorf("Token was incorrect, got %v, want: %v", token.FirebasePushtoken, testtoken.FirebasePushtoken)
	}
}
