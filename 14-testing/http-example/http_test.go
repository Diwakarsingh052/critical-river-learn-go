package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDoubleHandler(t *testing.T) {
	tt := [...]struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "http.StatusOk",
			queryParam:     "2",
			expectedStatus: http.StatusOK,
			expectedBody:   "double of number: 2 : 4",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//// ResponseRecorder is an implementation of [http.ResponseWriter] that
			//// records its mutations for later inspection in tests.
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, `/double?v=`+tc.queryParam, nil)

			doubleHandler(w, r)

			require.Equal(t, tc.expectedStatus, w.Code)
			require.Equal(t, tc.expectedBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
