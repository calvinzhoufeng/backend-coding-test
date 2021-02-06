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
	c.app.Get("/v1/note", c.GetAllNotes)
	c.app.Delete("/v1/note", c.DeleteAllNotes)

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

	notes, err := c.repo.GetNotes()

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      0,
			Message: "Unknown error",
		})
		return
	}

	notesDto := make([]*NoteDto, 0)
	for _, note := range notes {
		notesDto = append(notesDto, note.ToDto())
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(notesDto)
	return
}

// GetNoteById  Get a note by ID
// @Param noteId
// @Success 200 Note
// @Failure 400
// @router /notes [get]
func (c *NoteController) GetNoteById(ctx iris.Context) {
	note, _, isError := c.retrieveNoteById(ctx)
	if isError == true {
		return
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(note.ToDto())
	return
}

// DeleteNoteById  Remove a note by ID
// @Param noteId
// @Success 200 Note
// @Failure 400
// @router /notes [delete]
func (c *NoteController) DeleteNoteById(ctx iris.Context) {
	_, idInt, isError := c.retrieveNoteById(ctx)
	if isError == true {
		return
	}

	err := c.repo.DeleteNoteById(idInt)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      idInt,
			Message: "Unknown error",
		})
		return
	}

	log.Debug().Int("noteId", idInt).Msg("Removed a note")

	ctx.StatusCode(iris.StatusNoContent)
	return
}

// DeleteAllNotes it's only for unit testing
// @Param
// @Success 204
// @Failure 400
// @router /note [delete]
func (c *NoteController) DeleteAllNotes(ctx iris.Context) {

	log.Warn().Msg("Gonna remove all notes")

	err := c.repo.DeleteAllNotes()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      0,
			Message: "Unknown error",
		})
		return
	}

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
			ID:      0,
			Message: "Bad request body",
		})
		return
	}

	// TODO: Other validations are skipped due to time constraits
	if len(createNoteRequest.Content) == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Err{
			ID:      0,
			Message: "Invalid content",
		})
		return
	}

	tags := make([]Tag, 0)
	for _, t := range createNoteRequest.Tags {
		tags = append(tags, Tag{Name: t})
	}
	note := Note{
		Content: createNoteRequest.Content,
		Tags:    tags,
	}
	// log.Debug().Msgf("controller to be added %v %s", note, note.Content)

	n, err := c.repo.CreateNote(note)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      0,
			Message: "Unknown error",
		})
		return
	}

	// log.Debug().Int("id", n.ID).Msg("Added a new note")

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(n.ToDto())
	return
}

// UpdateNoteById Update a new note by Id
// @Param &Note{}
// @Success 200 Note
// @Failure 400
// @router /note/{id} [post]
func (c *NoteController) UpdateNoteById(ctx iris.Context) {

	var createNoteRequest CreateNoteRequest
	if err := ctx.ReadJSON(&createNoteRequest); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Err{
			ID:      0,
			Message: "Bad request body",
		})
		return
	}

	note, idInt, isError := c.retrieveNoteById(ctx)
	if isError == true {
		return
	}

	note.Content = createNoteRequest.Content
	log.Debug().Msgf("To update note %v\n", createNoteRequest.Tags)
	tags := make([]Tag, 0)
	for _, t := range createNoteRequest.Tags {
		tags = append(tags, Tag{Name: t})
	}
	note.Tags = tags
	log.Debug().Msgf("To update note %v\n", note.Tags)

	err := c.repo.UpdateNoteById(note)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      idInt,
			Message: "Unknown error",
		})
		return
	}
	log.Debug().Msgf("Updated note %v\n", note.ToDto())

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(note.ToDto())
	return
}

// retrieveNoteById is an internal function used by others
// @Param id
//
func (c *NoteController) retrieveNoteById(ctx iris.Context) (Note, int, bool) {
	isError := false
	idInt, isError := c.retreiveID(ctx)
	if isError == true {
		return Note{}, 0, isError
	}

	note, err := c.repo.GetNoteById(idInt)
	if err != nil {
		isError = true

		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      idInt,
			Message: "Unknown error",
		})
		return Note{}, 0, isError
	}

	if note.ID == 0 {
		isError = true

		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(&Err{
			ID:      idInt,
			Message: "Invalid ID",
		})
		return Note{}, 0, isError
	}

	return note, idInt, isError
}

func (c *NoteController) retreiveID(ctx iris.Context) (int, bool) {
	id := ctx.Params().Get("id")

	isError := false
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		isError = true

		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      idInt,
			Message: "Invalid note id",
		})
		return 0, isError
	}

	return idInt, isError
}
