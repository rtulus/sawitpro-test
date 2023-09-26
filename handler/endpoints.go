package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/jwt_usecase"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) Hello(ctx echo.Context, params generated.HelloParams) error {

	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

// User Registration API Handler
// (POST /register)
func (s *Server) Register(ctx echo.Context) error {

	// Fetch required request body post form
	ph := ctx.Request().PostFormValue("phone_number")
	name := ctx.Request().PostFormValue("full_name")
	pass := ctx.Request().PostFormValue("password")

	// validating input
	errMsgs := validateRegistrationInput(ph, name, pass)

	// convert phone number into int64
	phoneNum, err := strconv.ParseInt(strings.TrimPrefix(ph, "+"), 10, 64)
	if err != nil {
		errMsgs = append(errMsgs, "Phone number must consist of numerical characters after +62")
	}

	if len(errMsgs) > 0 {
		var errResponse generated.ErrorResponse
		errResponse.Message = strings.Join(errMsgs, ", ")
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	// make user password hashed and salted
	hashedPass, salt, err := hashAndSaltPassword(pass)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Failed to register new user"
		return ctx.JSON(http.StatusInternalServerError, errResponse)
	}

	// insert new user to DB
	input := repository.AddNewUserInput{
		FullName:       name,
		PhoneNumber:    phoneNum,
		HashedPassword: string(hashedPass),
		Salt:           string(salt),
	}
	err = s.Repository.AddNewUser(ctx.Request().Context(), input)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Failed to register new user"
		return ctx.JSON(http.StatusInternalServerError, errResponse)
	}

	// Registration API Response
	var resp generated.RegistrationResponse
	resp.Message = fmt.Sprintf("New user %s registered", name)
	return ctx.JSON(http.StatusOK, resp)
}

// User Authentication API Handler
// (POST /login)
func (s *Server) Login(ctx echo.Context) error {

	// Fetch required request body post form
	ph := ctx.Request().PostFormValue("phone_number")
	pass := ctx.Request().PostFormValue("password")

	// convert phone number into int64
	phoneNum, err := strconv.ParseInt(strings.TrimPrefix(ph, "+"), 10, 64)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Incorrect phone number or password"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	// fetch corresponding user from database
	var inputSelectUser repository.SelectUserByPhoneNumberInput
	inputSelectUser.PhoneNumber = phoneNum
	user, err := s.Repository.SelectUserByPhoneNumber(ctx.Request().Context(), inputSelectUser)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Incorrect phone number or password"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	// authenticate user
	salted := append([]byte(pass), user.Salt...)
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), salted)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Incorrect phone number or password"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	// generate jwt token
	var inputToken jwt_usecase.GenerateTokenInput
	inputToken.UserID = user.UserID
	inputToken.FullName = user.FullName
	inputToken.PhoneNumber = user.PhoneNumber
	token, err := s.JWTUsecase.GenerateUserToken(inputToken)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Failed to generate jwt"
		return ctx.JSON(http.StatusInternalServerError, errResponse)
	}

	// increment number of successful login
	var inputIncrement repository.IncrementUserSuccessfulLoginInput
	inputIncrement.UserID = user.UserID
	err = s.Repository.IncrementUserSuccessfulLogin(ctx.Request().Context(), inputIncrement)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Failed to increment successful login"
		return ctx.JSON(http.StatusInternalServerError, errResponse)
	}

	// Login API response
	var resp generated.LoginResponse
	resp.UserId = user.UserID
	resp.Token = token
	return ctx.JSON(http.StatusOK, resp)
}

// Get User Profile API Handler
// (GET /get-profile)
func (s *Server) GetProfile(ctx echo.Context, params generated.GetProfileParams) error {

	// extract jwt token from authorization header
	reqToken := params.Authorization
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// validate jwt token
	var input jwt_usecase.ValidateTokenInput
	input.Token = reqToken
	out, err := s.JWTUsecase.ValidateToken(input)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Authorization Failed"
		return ctx.JSON(http.StatusForbidden, errResponse)
	}

	// Get User Profile API response
	var resp generated.GetProfileResponse
	resp.FullName = out.FullName
	resp.PhoneNumber = out.PhoneNumber
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) UpdateProfile(ctx echo.Context, params generated.UpdateProfileParams) error {

	// Fetch request body post form
	ph := ctx.Request().PostFormValue("phone_number")
	name := ctx.Request().PostFormValue("full_name")

	// validate parameters
	err := validateUpdateProfileInput(ph, name)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Parameter Validation Failed"
		return ctx.JSON(http.StatusBadRequest, errResponse)
	}

	// extract jwt token from authorization header
	reqToken := params.Authorization
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	// validate jwt token
	var input jwt_usecase.ValidateTokenInput
	input.Token = reqToken
	out, err := s.JWTUsecase.ValidateToken(input)
	if err != nil {
		var errResponse generated.ErrorResponse
		errResponse.Message = "Authorization Failed"
		return ctx.JSON(http.StatusForbidden, errResponse)
	}

	if name != "" {
		var inputUpdateName repository.UpdateFullNameInput
		inputUpdateName.UserID = out.UserID
		inputUpdateName.FullName = name
		err = s.Repository.UpdateFullName(ctx.Request().Context(), inputUpdateName)
		if err != nil {
			var errResponse generated.ErrorResponse
			errResponse.Message = "Failed to update user full name"
			return ctx.JSON(http.StatusInternalServerError, errResponse)
		}
	}
	if ph != "" {
		// convert phone number into int64
		phoneNum, err := strconv.ParseInt(strings.TrimPrefix(ph, "+"), 10, 64)
		if err != nil {
			var errResponse generated.ErrorResponse
			errResponse.Message = "Phone number must consist of numerical characters after +62"
			return ctx.JSON(http.StatusBadRequest, errResponse)
		}

		// update phone number
		var inputUpdatePhoneNumber repository.UpdatePhoneNumberInput
		inputUpdatePhoneNumber.UserID = out.UserID
		inputUpdatePhoneNumber.PhoneNumber = phoneNum
		err = s.Repository.UpdatePhoneNumber(ctx.Request().Context(), inputUpdatePhoneNumber)
		if err != nil {
			var errResponse generated.ErrorResponse
			errResponse.Message = "Failed to update user phone number"
			return ctx.JSON(http.StatusConflict, errResponse)
		}
	}

	var resp generated.UpdateProfileResponse
	resp.Message = "Profile updated successfully"
	return ctx.JSON(http.StatusOK, resp)
}
