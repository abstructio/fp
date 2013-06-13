package store

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("mongourl", "localhost/testuser")
}

func drop() {
	session := connect()

	session.DB("").DropDatabase()

	defer session.Close()
}

func TestSha256(t *testing.T) {
	admin := "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918"
	x := sha("admin")
	if admin != x {
		t.Log("hash must be equal")
		t.Fail()
	}
}

func TestInsertAndUpdateUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()

	user := NewUser("ioboi", "Yannik", "Dällenbach", "test")

	if user.Id.Hex() != "" {
		t.Log("id must be empty")
		t.Fail()
	}

	inserted, _ := InsertOrUpdateUser(user)

	if !inserted {
		t.Log("must be inserted")
		t.Fail()
	}

	if user.Id.Hex() == "" {
		t.Log("id must be something after saving")
		t.Fail()
	}

	user.Username = "druic"

	updated, _ := InsertOrUpdateUser(user)

	if updated {
		t.Log("must be updated")
		t.Fail()
	}

}

func TestGetUserById(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()

	user := NewUser("ioboi", "Yannik", "Dällenbach", "test")

	InsertOrUpdateUser(user)

	result, err := UserById(user.Id)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if user.Username != result.Username {
		t.Log("usernames must match")
		t.Fail()
	}
}

func TestIfUsernameIsAviable(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()

	user := NewUser("ioboi", "Yannik", "Dällenbach", "test")

	InsertOrUpdateUser(user)

	isAviable := IsUsernameAviable("ioboi")

	if isAviable {
		t.Log("username must be taken")
		t.Fail()
	}
}

func TestUpdatePasswordHash(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()

	user := NewUser("ioboi", "Yannik", "Dällenbach", "test")

	InsertOrUpdateUser(user)

	user.Password = NewPasswordHash("stronger password")

	InsertOrUpdateUser(user)

	result, _ := UserById(user.Id)

	if result.Password != result.Password {
		t.Log("Passwords should be identical")
		t.Fail()
	}

	if result.Password != NewPasswordHash("stronger password") {
		t.Log("Password is somehow not updated")
		t.Fail()
	}
}

func TestLogin(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()

	user := NewUser("ioboi", "Yannik", "Dällenbach", "test")

	InsertOrUpdateUser(user)

	if !CheckUsernamePassword("ioboi", "test") {
		t.Log("user password combination should be right")
		t.Fail()
	}
}
