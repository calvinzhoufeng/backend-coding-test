package note

import (
	"go/src/pkg"

	"github.com/kataras/iris"
	"github.com/rs/zerolog/log"
)

// GetNotes Get all notes
// @Param
// @Success 200 []Notes
// @Failure 400
// @router /notes [get]
func (c *NoteController) GetNotesByTag(ctx iris.Context) {
	log.Debug().Msg("Get all notes by a tag")
	tag := ctx.Params().Get("tag")

	page, _ := ctx.Params().GetInt("tag")
	pageSize, _ := ctx.Params().GetInt("pageSize")

	notes, err := c.repo.GetNotesByTag(tag, page, pageSize)

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

// GetNotes Get all tags
// @Param
// @Success 200 []Notes
// @Failure 400
// @router /tags [get]
func (c *NoteController) GetAllTags(ctx iris.Context) {
	log.Debug().Msg("Get all tags")

	tags, err := c.repo.GetAllTags()

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      0,
			Message: "Unknown error",
		})
		return
	}

	results := make([]string, 0)
	for _, tag := range tags {
		results = append(results, tag.Name)
	}

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(&results)
	return
}

// GetNotes Get all tags
// @Param
// @Success 200 []Notes
// @Failure 400
// @router /note/funny [get]
func (c *NoteController) GenFunnyNote(ctx iris.Context) {

	client := pkg.NewClient(
		"https://randomuser.me/api/?inc=name",
		"http://api.icndb.com/jokes/random?",
	)

	firstName, lastName, err := client.UserGenerator()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		_, _ = ctx.JSON(&Err{
			ID:      0,
			Message: "Unknown error",
		})
		return
	}

	content, err := client.ContentGenerator(firstName, lastName)

	tag := Tag{
		Name: "funny",
	}
	tags := []Tag{tag}
	n, err := c.repo.CreateNote(Note{
		Content: content,
		Tags:    tags,
	})

	log.Debug().Msgf("Create a random note %v\n", n)
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(n.ToDto())
}
