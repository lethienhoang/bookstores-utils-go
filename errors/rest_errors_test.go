package errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("This is a bad request message")
	assert.EqualValues(t, err.Message, "This is a bad request message")
	assert.EqualValues(t,
		err.Code,
		http.StatusBadRequest,
	)
}

func TestNewInternalError(t *testing.T) {
	err := NewInternalError("This is a internal error message")
	assert.EqualValues(t, err.Message, "This is a internal error message")
	assert.EqualValues(t,
		err.Code,
		http.StatusInternalServerError,
	)
}

func TestNewNotFoundError(t *testing.T) {
	err := NewBadRequestError("This is a not found message")
	assert.EqualValues(t, err.Message, "This is a not found message")
	assert.EqualValues(t,
		err.Code,
		http.NotFound,
	)
}
