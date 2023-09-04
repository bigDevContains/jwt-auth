package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "github.com/bigDevContains/golang-jwt-project/helpers"
	"github.com/bigDevContains/golang-jwt-project/models"
	"github.com/bigDevContains/golang-jwt-project/helpers"
	"github.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.openCollection(database.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPssword()

func SignUp()

func Login()

func GetUsers()

func GetUSer() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		helper.MatchUserTypeTOUid(c, userId); 
	}
}
