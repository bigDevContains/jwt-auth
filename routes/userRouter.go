package routes

import(
	controller "github.com/bigDevContains/golang-jwt-project/controllers"
	
	"github.com/gin-gonic/gin"
)

func UserRoutes(incoming Routes *gin.Engine){
	incomingToutes.POST("users/signup", controller.Signup())
	incomingROutes.POST("users/login", controller.Login())
}