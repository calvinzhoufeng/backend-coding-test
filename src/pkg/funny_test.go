package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserGenerator(t *testing.T) {
	client := NewClient(
		"https://randomuser.me/api/?inc=name",
		"http://api.icndb.com/jokes/random?",
	)

	firstName, lastName, err := client.UserGenerator()
	if err != nil {
		assert.FailNow(t, "Failed test")
	}

	fmt.Println(firstName)
	fmt.Println(lastName)
	assert.NotEmpty(t, firstName)
}

func TestContentGenerator(t *testing.T) {
	client := NewClient(
		"https://randomuser.me/api/?inc=name",
		"http://api.icndb.com/jokes/random?",
	)

	joke, err := client.ContentGenerator("Calvin", "Zero")
	if err != nil {
		assert.FailNow(t, "Failed test")
	}

	assert.NotEmpty(t, joke)
}
