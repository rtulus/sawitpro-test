package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	s := &Server{}
	param := generated.HelloParams{
		Id: 123,
	}

	resp := `{"message":"Hello User 123"}
`
	// Assertions
	if assert.NoError(t, s.Hello(c, param)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, resp, rec.Body.String())
	}
}

// func TestRegister(t *testing.T) {
// 	// Setup
// 	e := echo.New()
// 	f := make(url.Values)
// 	f.Set("phone_number", "+628788155269")
// 	f.Set("full_name", "Tommy")
// 	f.Set("password", "Hahaha123!@#")
// 	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(f.Encode()))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	mockRepository := repository.NewMockRepositoryInterface(gomock.NewController(t))
// 	// mockRepository.EXPECT().AddNewUser(c.Request(), ).Return()
// 	s := &Server{
// 		Repository: mockRepository,
// 	}

// 	resp := `{"message":"New user Tommy registered"}`
// 	// Assertions
// 	if assert.NoError(t, s.Register(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, resp, rec.Body.String())
// 	}

// }
