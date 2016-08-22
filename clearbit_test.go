package clearbit

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var apiKey = os.Getenv("API_KEY")

func TestPerson(t *testing.T) {
	api := New(apiKey)
	person, err := api.Person("hi@mvader.me", false)
	assert.Nil(t, err)
	assert.NotNil(t, person)
	assert.Equal(t, "José Miguel Molina", person.Name.Full)

	person, err = api.Person("hi@mvader.me", true)
	assert.Nil(t, err)
	assert.NotNil(t, person)
	assert.Equal(t, "José Miguel Molina", person.Name.Full)
}

func TestCompany(t *testing.T) {
	api := New(apiKey)
	company, err := api.Company("sourced.tech", false)
	assert.Nil(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, "source{d}", company.Name)

	company, err = api.Company("sourced.tech", true)
	assert.Nil(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, "source{d}", company.Name)
}
