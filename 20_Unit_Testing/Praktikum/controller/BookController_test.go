package controller

import (
	"github.com/stretchr/testify/assert"
)

func TestGetBooksController_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{
		{
			name:                   "get book normal",
			path:                   "/books",
			expectedStatus:         http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}

	}
}
