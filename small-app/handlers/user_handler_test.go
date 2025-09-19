package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"small-app/internal/users"
	"small-app/internal/users/mockusers"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestSignup(t *testing.T) {
	// sample newUser that we would get after reading JSON body
	newUser := users.NewUser{
		Name:     "John Doe",
		Age:      30,
		Email:    "d@email.com",
		Password: "abc",
	}

	// User that we would get from CreatUser method of users package
	mockUser := users.User{
		Id:           "ab49a45c-ec2c-47a5-8675-9f072e2d9216",
		Email:        "d@email.com",
		Name:         "John Doe",
		Age:          30,
		PasswordHash: "2a$10$EimVQRw4YiKIoMqh3JMwOesA9ngPGZT.chFEmPSaHzYl.mlnhLr12",
	}

	tt := [...]struct {
		name             string
		body             []byte
		expectedStatus   int
		expectedResponse string
		mockStore        func(m *mockusers.MockStore)
	}{
		{
			name: "Ok",
			body: []byte(`{
   				 "name": "John Doe",
   				 "age": 30,
                 "email": "d@email.com",
                 "password": "abc"
			}`),
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"Id":"ab49a45c-ec2c-47a5-8675-9f072e2d9216","email":"d@email.com","name":"John Doe","age":30,"password_hash":"2a$10$EimVQRw4YiKIoMqh3JMwOesA9ngPGZT.chFEmPSaHzYl.mlnhLr12"}`,
			mockStore: func(m *mockusers.MockStore) {
				m.EXPECT().CreatUser(gomock.Eq(newUser)).Return(mockUser, nil).Times(1)
			},
		},
		{
			name: "Fail_NoEmail",
			body: []byte(`{
  						"name": "John Doe",
  						"age": 30,
  						"password": "abc"
  				}`),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `{"error":"Key: 'NewUser.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`,

			mockStore: func(m *mockusers.MockStore) {
				m.EXPECT().CreatUser(gomock.Any()).Times(0)
			},
		},
	}
	// NewController need to be passed to the NewMockStore
	// it is used for testing, takes testing.T as an argument
	controller := gomock.NewController(t)

	//NewMockStore would return the implementation of the interface
	mock := mockusers.NewMockStore(controller)

	// setting the handler with the mock implementation of the interface
	h := handler{
		uc: mock,
	}
	router := gin.New()
	//we need gin router to register the /signup endpoint
	router.POST("/signup", h.Signup)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// calling MockStore to set expectations for each test case
			tc.mockStore(mock)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(tc.body))
			router.ServeHTTP(w, r)

			require.Equal(t, tc.expectedStatus, w.Code)
			require.Equal(t, tc.expectedResponse, strings.TrimSpace(w.Body.String()))

		})
	}
}
