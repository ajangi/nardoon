package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ajangi/nardoon/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

// RegisterUserBody : this struct type is to bind register user request body
type RegisterUserBody struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required,len=11"`
	Email    string `json:"email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

var validate *validator.Validate

// RegisterUser : this method is to register users by phone number and password and name
func RegisterUser(c echo.Context) (err error) {
	body := new(RegisterUserBody)
	// this part is to check if the body is empty or a valid json
	if err = c.Bind(body); err != nil {
		emptyData := utils.ResponseData{}
		badRequestMessage := utils.ResponseMessages{utils.GetMessageByKey(utils.InputErrorMessageKey)}
		errorResponse := utils.ResponseApi{Result: "ERROR", Data: emptyData, Messages: badRequestMessage, StatusCode: http.StatusBadRequest}
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	validate = validator.New()
	err = validate.Struct(body)
	// this part is to check the request body validation
	if err != nil {
		emptyData := utils.ResponseData{}
		badRequestMessage := utils.ResponseMessages{utils.GetMessageByKey(utils.InputErrorMessageKey)}
		errorResponse := utils.ResponseApi{Result: "ERROR", Data: emptyData, Messages: badRequestMessage, StatusCode: http.StatusBadRequest}
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	// in this part we can get valid data to register user in database and return jwt token
	name := body.Name
	email := body.Email
	phone := body.Phone
	rawPassword := body.Password
	hashedPassword, _ := HashPassword(rawPassword)
	role := strings.ToLower(body.Role)
	db := utils.DbConn()
	insUser, err := db.Prepare("INSERT INTO users(name, email,phone,password) VALUES(?,?,?,?)")
	if err != nil {
		// todo : check this error here
		panic(err.Error())
	}
	res, err := insUser.Exec(name, email, phone, hashedPassword)
	if err != nil {
		fmt.Println(err.Error())
		// todo : check this error here
		panic(err.Error())
	}
	userID, _ := res.LastInsertId()
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow("select id from roles where name = ?", role)
	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		fmt.Println("ErrNoRows")
		// todo : Check this error!!!!
	case nil:
		insRoleUser, err := db.Prepare("INSERT INTO users_roles(user_id, role_id) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insRoleUser.Exec(userID, id)
	default:
		fmt.Println(err)
		// todo : Check this error!!!!
	}
	// Create token for the registered user
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 3000).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		// TODO : check this error here
		return err
	}
	emptyData := utils.ResponseData{
		map[string]string{"token": t},
	}
	healthyMessage := utils.ResponseMessages{}
	healthyResponse := utils.ResponseApi{Result: "SUCCESS", Data: emptyData, Messages: healthyMessage, StatusCode: http.StatusOK}
	return c.JSON(http.StatusOK, healthyResponse)
}

// HashPassword : this method is to make password hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash : this method is to check password and hashed one
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
