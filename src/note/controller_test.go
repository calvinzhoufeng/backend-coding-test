package note

import (
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

type RepositoryMock struct{}

var createNoteRequest = &CreateNoteRequest{
	Content: "Test",
	Tags:    []string{"test", "work"},
}

var tag1 = Tag{Name: "test"}
var tag2 = Tag{Name: "work"}
var tags = []Tag{tag1, tag2}
var tag = &Tag{
	Name: "test",
}
var note = &Note{
	ID:      1,
	Content: "Test",
	Tags:    tags,
}

/**
 * A bunch of mock functions to mock dependencies
 */
func (r *RepositoryMock) CreateNote(note Note) (Note, error) {
	note.ID = 1
	return note, nil
}

func (r *RepositoryMock) GetRideById(id string) (Note, error) {
	note.ID = 1
	return *note, nil
}
func (r *RepositoryMock) GetNotes() ([]Note, error)        { return []Note{*note}, nil }
func (r *RepositoryMock) GetNoteById(id int) (Note, error) { return *note, nil }
func (r *RepositoryMock) UpdateNoteById(note Note) error   { return nil }
func (r *RepositoryMock) DeleteNoteById(id int) error      { return nil }
func (r *RepositoryMock) DeleteAllNotes() error            { return nil }

func (r *RepositoryMock) GetNotesByTag(tag string, page int, pageSize int) ([]Note, error) {
	return []Note{}, nil
}
func (r *RepositoryMock) GetAllTags() ([]Tag, error) { return []Tag{}, nil }

func TestController(t *testing.T) {
	repositoryMock := &RepositoryMock{}

	controller := NewController(iris.New(), repositoryMock)
	e := httptest.New(t, controller.app)

	// redirects to /v1/note without basic app
	e.POST("/v1/note").WithJSON(createNoteRequest).Expect().Status(httptest.StatusOK)

	obj := e.GET("/v1/note").Expect().Status(httptest.StatusOK).JSON().Array()

	// obj.JSON().Array().Element(0)
	t.Logf("returned data %v\n", obj.Element(0).Object())
}
