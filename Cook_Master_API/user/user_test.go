package user

import (
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Remove(dbPath)

}

func cleanDB(b *testing.B) {
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:            bson.NewObjectId(),
			FirstName:     "JJ",
			Name:          "Abrams",
			Email:         "abrams@gmail.com",
			Role:          "Student",
			Address:       "1 rue de la Paix",
			Phone:         "0123456789",
			Subscribed_ID: bson.NewObjectId().Time().Day(),
			Loyality_ID:   bson.NewObjectId().Time().Day(),
		}
		b.StartTimer()
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}
		b.ResetTimer()
	}
}

func BenchmarkRead(b *testing.B) {
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:            bson.NewObjectId(),
			FirstName:     "JJ",
			Name:          "Abrams_" + strconv.Itoa(i),
			Email:         "abrams@gmail.com",
			Role:          "Student",
			Address:       "1 rue de la Paix",
			Phone:         "0123456789",
			Subscribed_ID: bson.NewObjectId().Time().Day(),
			Loyality_ID:   bson.NewObjectId().Time().Day(),
		}

		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}
		b.StartTimer()
		_, err = One(u.ID)
		if err != nil {
			b.Fatalf("Error retrieving user: %s", err)
		}
	}
}

func BenchmarkUpdate(b *testing.B) {
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:            bson.NewObjectId(),
			FirstName:     "JJ",
			Name:          "Abrams_" + strconv.Itoa(i),
			Email:         "abrams@gmail.com",
			Role:          "Student",
			Address:       "1 rue de la Paix",
			Phone:         "0123456789",
			Subscribed_ID: bson.NewObjectId().Time().Day(),
			Loyality_ID:   bson.NewObjectId().Time().Day(),
		}

		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}
		b.StartTimer()
		u.Role = "Professor"
		err = u.Save()
		if err != nil {
			b.Fatalf("Error updating user: %s", err)
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:            bson.NewObjectId(),
			FirstName:     "JJ",
			Name:          "Abrams_" + strconv.Itoa(i),
			Email:         "abrams@gmail.com",
			Role:          "Student",
			Address:       "1 rue de la Paix",
			Phone:         "0123456789",
			Subscribed_ID: bson.NewObjectId().Time().Day(),
			Loyality_ID:   bson.NewObjectId().Time().Day(),
		}
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}
		b.StartTimer()
		err = Delete(u.ID)
		if err != nil {
			b.Fatalf("Error deleting user: %s", err)
		}
	}
}

func BenchmarkCreate(b *testing.B) {
	cleanDB(b)
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:            bson.NewObjectId(),
			FirstName:     "JJ",
			Name:          "Abrams_" + strconv.Itoa(i),
			Email:         "abrams@gmail.com",
			Role:          "Student",
			Address:       "1 rue de la Paix",
			Phone:         "0123456789",
			Subscribed_ID: bson.NewObjectId().Time().Day(),
			Loyality_ID:   bson.NewObjectId().Time().Day(),
		}
		b.StartTimer()
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}
		_, err = One(u.ID)
		if err != nil {
			b.Fatalf("Error retrieving user: %s", err)
		}

		u.Role = "Professor"
		err = u.Save()
		if err != nil {
			b.Fatalf("Error updating user: %s", err)
		}

		err = Delete(u.ID)
		if err != nil {
			b.Fatalf("Error deleting user: %s", err)
		}
	}
}

func BenchmarkCRUD(b *testing.B) {
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := &User{
			ID:            bson.NewObjectId(),
			FirstName:     "JJ",
			Name:          "Abrams",
			Email:         "abrams@gmail.com",
			Role:          "Student",
			Address:       "1 rue de la Paix",
			Phone:         "0123456789",
			Subscribed_ID: bson.NewObjectId().Time().Day(),
			Loyality_ID:   bson.NewObjectId().Time().Day(),
		}

		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving user: %s", err)
		}

		_, err = One(u.ID)
		if err != nil {
			b.Fatalf("Error retrieving user: %s", err)
		}

		u.Role = "Professor"
		err = u.Save()
		if err != nil {
			b.Fatalf("Error updating user: %s", err)
		}

		err = Delete(u.ID)
		if err != nil {
			b.Fatalf("Error deleting user: %s", err)
		}
	}
}

func TestCRUD(t *testing.T) {
	t.Log("Create")
	u := &User{
		ID:            bson.NewObjectId(),
		FirstName:     "JJ",
		Name:          "Abrams",
		Email:         "abrams@gmail.com",
		Role:          "Student",
		Address:       "1 rue de la Paix",
		Phone:         "0123456789",
		Subscribed_ID: bson.NewObjectId().Time().Day(),
		Loyality_ID:   bson.NewObjectId().Time().Day(),
	}

	err := u.Save()
	if err != nil {
		t.Fatalf("Error saving user: %s", err)
	}

	t.Log("Read")
	u2, err := One(u.ID)
	if err != nil {
		t.Fatalf("Error retrieving user: %s", err)
	}
	if !reflect.DeepEqual(u2, u) {
		t.Errorf("Expected user to be %v, got %v", u, u2)
	}

	t.Log("Update")
	u.Role = "Professor"
	err = u.Save()
	if err != nil {
		t.Fatalf("Error updating user: %s", err)
	}

	u3, err := One(u.ID)
	if err != nil {
		t.Fatalf("Error retrieving user: %s", err)
	}
	if !reflect.DeepEqual(u3, u) {
		t.Errorf("Expected user to be %v, got %v", u, u2)
	}
	t.Log("Delete")
	err = Delete(u.ID)
	if err != nil {
		t.Fatalf("Error deleting user: %s", err)
	}

	_, err = One(u.ID)
	if err == nil {
		t.Fatalf("Record was not deleted, but should have been")
	}
	if err != storm.ErrNotFound {
		t.Fatalf("Error retrieving non-existing recor: %s", err)
	}

	t.Log("Read all")
	u2.ID = bson.NewObjectId()
	u3.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil {
		t.Fatalf("Error Saving a record: %s", err)
	}
	err = u2.Save()
	if err != nil {
		t.Fatalf("Error Saving a record: %s", err)
	}

	err = u3.Save()
	if err != nil {
		t.Fatalf("Error Saving a record: %s", err)
	}
	users, err := All()
	if err != nil {
		t.Fatalf("Error reading all the records: %s", err)
	}
	if len(users) != 3 {
		t.Errorf("Expected 3 records, got %d", len(users))
	}
}
