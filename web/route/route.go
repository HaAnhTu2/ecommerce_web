package route

import (
	"os"

	"github.com/HaAnhTu2/ecommerce_web.git/controller"
	"github.com/HaAnhTu2/ecommerce_web.git/db"
	"github.com/HaAnhTu2/ecommerce_web.git/middleware"
	"github.com/HaAnhTu2/ecommerce_web.git/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Route(r *gin.Engine, DB *mongo.Database) {
	client := db.Connect()
	ProductRepo := repository.NewProductRepo(client.Database(os.Getenv("DB_NAME")))
	productController := controller.NewProductController(ProductRepo, DB)
	UserRepo := repository.NewUserRepo(client.Database(os.Getenv("DB_NAME")))
	userController := controller.NewUserController(UserRepo, DB)
	authMiddleware := middleware.AuthMiddleware
	r.POST("api/login", userController.Login)
	r.DELETE("api/logout", userController.Logout)
	auth := r.Group("/")
	auth.Use(authMiddleware)
	{

		adminMiddleware := middleware.RequireAdmin
		admin := r.Group("/dashboard")
		admin.Use(adminMiddleware)
		{
			admin.POST("/api/admin/create", userController.CreateUserAdmin)
			admin.PUT("/api/user/update/:id", userController.UpdateUser)
			admin.DELETE("/api/user/delete/:id", userController.DeleteUser)

			admin.POST("/api/product/create", productController.CreateProduct)
			admin.PUT("/api/product/update/:id", productController.UpdateProduct)
			admin.DELETE("/api/product/delete/:id", productController.DeleteProduct)
		}
	}

	r.POST("/api/user/create", userController.CreateUser)
	r.GET("/api/user/get/:id", userController.GetByID)

	r.GET("/api/product/get", productController.GetAllProduct)
	r.GET("/api/product/get/:id", productController.GetByID)
}
