package store

import (
	"errors"
	"labix.org/v2/mgo/bson"
)

type Presentation struct {
	Id     bson.ObjectId   `bson:"_id"`
	Title  string          //Title of the whole presentation
	User   []bson.ObjectId //the user who are owning the presentation
	Slides []*Slide        //the slides of the presentation
}

type Slide struct {
	Title   string   //title of the slide
	Class   []string //html class for the slide
	Sclass  []string //special class for the background
	Content string   //content of the slide
	Notes   string   //notes for presenter
}

//Create a new presentation with a specific title.
//Attention, presentations are deffered by id not by title!!!
func NewPresentation(title string) *Presentation {
	return &Presentation{
		Title:  title,
		User:   make([]bson.ObjectId, 0),
		Slides: make([]*Slide, 0),
	}
}

//Create a new slide. The slide is added to the presentation
func (p *Presentation) NewSlide() *Slide {
	s := &Slide{
		Class:  make([]string, 0),
		Sclass: make([]string, 0),
	}

	p.Slides = append(p.Slides, s)
	return s
}

//Add a user to the presentation.
//If the user is not saved in the database(the id is empty),
//an error is returned
func (p *Presentation) AddUser(u *User) error {
	if u.Id.Hex() == "" {
		return errors.New("User is not saved yet")
	}

	p.User = append(p.User, u.Id)
	return nil
}

//Insert or Update a presentation.
//This method finaly writes the presentation object into the database.
func InsertOrUpdatePresentation(p *Presentation) (bool, error) {
	session := connect()
	defer session.Close()

	if p.Id.Hex() == "" {
		p.Id = bson.NewObjectId()
	}

	info, err := session.DB("").C("pres").UpsertId(p.Id, p)

	if err != nil {
		panic(err)
	}

	return info.Updated == 0, err
}

//Get a presentation by an id.
func PresentationById(id bson.ObjectId) (*Presentation, error) {
	session := connect()
	defer session.Close()

	result := Presentation{}

	c := session.DB("").C("pres")

	err := c.Find(bson.M{"_id": id}).One(&result)

	return &result, err
}

//Get all presentations of a specific user.
func PresentationByUser(u *User) ([]*Presentation, error) {
	session := connect()
	defer session.Close()

	result := make([]*Presentation, 0)

	c := session.DB("").C("pres")

	err := c.Find(bson.M{"user": bson.M{"$in": []bson.ObjectId{u.Id}}}).All(&result)

	return result, err
}
