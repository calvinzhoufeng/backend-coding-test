package note

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"
)

// NoteController allows caller to do inject dependencies
type NoteController struct {
	app  *iris.Application
	repo Repository
}

// NewController always use dependency injection in this contructor
func NewController(app *iris.Application, repository Repository) *NoteController {

	c := &NoteController{
		app:  app,
		repo: repository,
	}

	// List down all services exposed
	c.app.Post("/v1/note", c.AddNote)
	c.app.Get("/v1/note/{id}", c.GetNoteById)
	c.app.Put("/v1/note/{id}", c.UpdateNoteById)
	c.app.Delete("/v1/note/{id}", c.DeleteNoteById)
	c.app.Get("/v1/notes", c.GetAllNotes)
	// c.app.Delete("/v1/notes", c.DeleteAllNotes)

	c.app.Get("/v1/note/tag/{tag}", c.GetNotesByTag)
	c.app.Get("/v1/tags", c.GetAllTags)

	c.app.Post("/v1/note/funny", c.GenFunnyNote)

	log.Info().Msg("Note controller initialized successfully")

	return c
}

// GetNotes Get all notes
// @Param
// @Success 200 []Notes
// @Failure 400
// @router /notes [get]
func (c *NoteController) GetAllNotes(ctx iris.Context) {
	log.Debug().Msg("Get all notes")
	page, _ := ctx.Params().GetInt("page")
	pageSize, _ := ctx.Params().GetInt("pageSize")

	notes, err := c.repo.GetNotes(page, pageSize)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      "SERVER_ERROR",
			Message: "Unknown error",
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&notes)
	return
}

// GetNoteById  Get a note by ID
// @Param noteId
// @Success 200 Note
// @Failure 400
// @router /notes [get]
func (c *NoteController) GetNoteById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      "REQ_DATA_ERROR",
			Message: "Invalid note id",
		})
		return
	}

	note, err := c.repo.GetNoteById(idInt)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      "SERVER_ERROR",
			Message: "Unknown error",
		})
		return
	}

	if note.ID == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Err{
			ID:      "1",
			Message: "Invalid ID",
		})
		return
	}

	log.Debug().Int("noteId", note.ID).Msg("Found a note")

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&note)
	return
}

// DeleteNoteById  Remove a note by ID
// @Param noteId
// @Success 200 Note
// @Failure 400
// @router /notes [delete]
func (c *NoteController) DeleteNoteById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      "REQ_DATA_ERROR",
			Message: "Invalid note id",
		})
		return
	}

	err = c.repo.DeleteNoteById(idInt)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      "SERVER_ERROR",
			Message: "Unknown error",
		})
		return
	}

	log.Debug().Int("noteId", idInt).Msg("Removed a note")

	ctx.StatusCode(iris.StatusNoContent)
	return
}

// AddNotes Add a new note
// @Param &Note{}
// @Success 200 Note
// @Failure 400
// @router /notes [post]
func (c *NoteController) AddNote(ctx iris.Context) {
	var createNoteRequest CreateNoteRequest
	if err := ctx.ReadJSON(&createNoteRequest); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Err{
			ID:      "002",
			Message: "Bad request body",
		})
		return
	}

	// TODO: Other validations are skipped due to time constraits
	if len(createNoteRequest.Content) == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Err{
			ID:      "003",
			Message: "Invalid content",
		})
		return
	}

	note := Note{
		Content: createNoteRequest.Content,
	}
	log.Debug().Msgf("controller to be added %v %s", note, note.Content)

	n, err := c.repo.CreateNote(note)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      "SERVER_ERROR",
			Message: "Unknown error",
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&n)
	return
}

// UpdateNoteById Update a new note by Id
// @Param &Note{}
// @Success 200 Note
// @Failure 400
// @router /note/{id} [post]
func (c *NoteController) UpdateNoteById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      "REQ_DATA_ERROR",
			Message: "Invalid note id",
		})
		return
	}

	var createNoteRequest CreateNoteRequest
	if err := ctx.ReadJSON(&createNoteRequest); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Err{
			ID:      "002",
			Message: "Bad request body",
		})
		return
	}

	note, err := c.repo.GetNoteById(idInt)
	if err != nil || note.ID <= 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Err{
			ID:      "SERVER ERROR",
			Message: "Invalid db error",
		})
		return
	}

	note.Content = createNoteRequest.Content
	for _, tmp := range createNoteRequest.Tags {
		note.Tags = append(note.Tags, Tag{
			Name: tmp,
		})
	}

	err = c.repo.UpdateNoteById(note)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      "SERVER_ERROR",
			Message: "Unknown error",
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&note)
	return
}
