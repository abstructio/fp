package store

import (
	"testing"
)

func TestNewPresentation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	p := NewPresentation("Hello World")

	if p.Title != "Hello World" {
		t.Log("Title should be hello world")
		t.Fail()
	}

	if len(p.User) != 0 {
		t.Log("0 user expected")
		t.Fail()
	}

	if len(p.Slides) != 0 {
		t.Log("0 slides expected")
		t.Fail()
	}
}

func TestCreateSlide(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	p := NewPresentation("Hello World")

	p.NewSlide()

	if len(p.Slides) == 0 {
		t.Log("1 slide expected")
		t.Fail()
	}
}

func TestCreateSimplePres(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()
	p := NewPresentation("Hello World")

	s := p.NewSlide()

	s.Title = "TOC"

	isNew, err := InsertOrUpdatePresentation(p)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if !isNew {
		t.Log("should be new")
		t.Fail()
	}
}

func TestGetByIdPres(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()

	p := NewPresentation("Hello World")

	s := p.NewSlide()

	s.Title = "TOC"

	InsertOrUpdatePresentation(p)

	r, err := PresentationById(p.Id)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if r.Title != p.Title {
		t.Log("Title must match")
		t.Fail()
	}

	if len(p.Slides) != len(r.Slides) {
		t.Log("not == num of slides")
		t.Fail()
	}

}

func TestCreatePresNewNotSavedUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	p := NewPresentation("Hello Word")

	err := p.AddUser(NewUser("", "", "", ""))

	if err == nil {
		t.Log("There should be an error")
		t.Fail()
	}
}

func TestCreatePresUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()

	p := NewPresentation("Hello World")

	u := NewUser("ioboi", "", "", "")

	InsertOrUpdateUser(u)

	err := p.AddUser(u)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	isNew, err := InsertOrUpdatePresentation(p)

	if !isNew {
		t.Log("should be new")
		t.Fail()
	}

	if err != nil {
		t.Log(err)
		t.Fail()
	}

}

func TestGetPresentationsOfUser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	drop()

	p := NewPresentation("Hello World")

	u := NewUser("ioboi", "", "", "")

	InsertOrUpdateUser(u)

	err := p.AddUser(u)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	isNew, err := InsertOrUpdatePresentation(p)

	if !isNew {
		t.Log("should be new")
		t.Fail()
	}

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	res, err := PresentationByUser(u)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(res) == 0 {
		t.Log("Should be 1")
		t.Fail()
	}

	if res[0].Title != "Hello World" {
		t.Log("Title should be 'Hello World'")
		t.Fail()
	}

}
