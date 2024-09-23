package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/HaAnhTu2/ecommerce_web.git/model"
	"github.com/HaAnhTu2/ecommerce_web.git/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	UserRepoInterface repository.UserRepoInterface
	DB                *mongo.Database
}

func NewUserController(UserRepoI repository.UserRepoInterface, db *mongo.Database) *UserController {
	return &UserController{UserRepoInterface: UserRepoI, DB: db}
}

func (u *UserController) Login(c *gin.Context) {
	var auth model.LoginRequest
	if err := c.ShouldBind(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := u.UserRepoInterface.FindByEmail(c.Request.Context(), auth.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}
	if auth.Email == user.Email && auth.Password == user.Password {
		token, err := u.UserRepoInterface.SaveToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		cookie := http.Cookie{}
		cookie.Name = "Token"
		cookie.Value = token
		cookie.Expires = time.Now().Add(15 * time.Minute)
		http.SetCookie(c.Writer, &cookie)
		log.Printf("User role: %s", user.Role)

		if user.Role == "Admin" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the admin dashboard!",
				"role":    user.Role,
				"token":   token,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the home page!",
				"role":    user.Role,
				"token":   token,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
	}
}

func (u *UserController) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "Token",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Delete the Cookie
	})
	c.JSON(http.StatusOK, gin.H{
		"data": "Logout successful!",
	})
}

func (u *UserController) GetByID(c *gin.Context) {
	userId := c.Param("id")
	user, err := u.UserRepoInterface.FindByID(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (u *UserController) CreateUser(c *gin.Context) {
	// user := model.User{
	// 	UserName: c.Request.FormValue("username"),
	// 	Password: c.Request.FormValue("password"),
	// 	Email:    c.Request.FormValue("email"),
	// 	Address:  c.Request.FormValue("address"),
	// 	Phone:    c.Request.FormValue("phone"),
	// }
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Role = "User"
	// Insert the user into the database
	users, err := u.UserRepoInterface.Create(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": users,
	})
}

func (u *UserController) CreateUserAdmin(c *gin.Context) {
	// user := model.User{
	// 	UserName: c.Request.FormValue("username"),
	// 	Email:    c.Request.FormValue("email"),
	// 	Password: c.Request.FormValue("password"),
	// 	Address:  c.Request.FormValue("address"),
	// 	Phone:    c.Request.FormValue("phone"),
	// 	Role:     c.Request.FormValue("role"),
	// }
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Role = "Admin"
	// Insert the user into the database
	users, err := u.UserRepoInterface.Create(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": users,
	})
}
func (u *UserController) UpdateUser(c *gin.Context) {
	var updateuser model.User
	if err := c.ShouldBind(&updateuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userId := c.Param("id")
	user, err := u.UserRepoInterface.FindByID(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	//postman
	user.UserName = updateuser.UserName
	user.Password = updateuser.Password
	user.Email = updateuser.Email
	user.Address = updateuser.Address
	user.Phone = updateuser.Phone
	//form
	if username := c.PostForm("username"); username != "" {
		user.UserName = username
	}
	if email := c.PostForm("email"); email != "" {
		user.Email = email
	}
	if password := c.PostForm("password"); password != "" {
		user.Password = password
	}
	if address := c.PostForm("address"); address != "" {
		user.Address = address
	}
	if phone := c.PostForm("phone"); phone != "" {
		user.Phone = phone
	}
	if role := c.PostForm("role"); role != "" {
		user.Role = role
	}

	updatedUser, err := u.UserRepoInterface.Update(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update user",
			"err":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": updatedUser,
	})
}
func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid argument id"})
		return
	}
	if err := u.UserRepoInterface.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
