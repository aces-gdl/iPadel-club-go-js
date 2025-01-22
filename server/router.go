package server

import (
	"ipadel-club/controllers"
	"ipadel-club/middleware"
	"ipadel-club/ui"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Recovery())

	// ruta para servir el front-end en ReactJS
	ui.AddRoutes(router)

	// Security
	router.POST("/v1/auth/register", controllers.Signup)
	router.POST("/v1/auth/login", controllers.Login)
	router.POST("/v1/auth/logout", controllers.Logout)

	// Rutas para clubes
	clubRoutes := router.Group("/v1/catalogs/clubs")
	{
		clubRoutes.POST("/", middleware.RequireAuth, controllers.CreateClub)
		clubRoutes.GET("/", middleware.RequireAuth, controllers.GetAllClubs)
		clubRoutes.GET("/:id", middleware.RequireAuth, controllers.GetClub)
		clubRoutes.PUT("/:id", middleware.RequireAuth, controllers.UpdateClub)
		clubRoutes.DELETE("/:id", middleware.RequireAuth, controllers.DeleteClub)
	}

	// Rutas para eventos
	eventRoutes := router.Group("/v1/catalogs/events")
	{
		eventRoutes.POST("/", middleware.RequireAuth, controllers.CreateEvent)
		eventRoutes.GET("/", middleware.RequireAuth, controllers.GetAllEvents)
		eventRoutes.GET("/:id", middleware.RequireAuth, controllers.GetEvent)
		eventRoutes.PUT("/:id", middleware.RequireAuth, controllers.UpdateEvent)
		eventRoutes.DELETE("/:id", middleware.RequireAuth, controllers.DeleteEvent)
	}

	// Rutas para categorías
	categoryRoutes := router.Group("/v1/catalogs/categories")
	{
		categoryRoutes.POST("/", middleware.RequireAuth, controllers.CreateCategory)
		categoryRoutes.GET("/", middleware.RequireAuth, controllers.GetAllCategories)
		categoryRoutes.GET("/:id", middleware.RequireAuth, middleware.RequireAuth, controllers.GetCategory)
		categoryRoutes.PUT("/:id", middleware.RequireAuth, controllers.UpdateCategory)
		categoryRoutes.DELETE("/:id", middleware.RequireAuth, controllers.DeleteCategory)
	}

	// Rutas para jugadores
	playerRoutes := router.Group("/v1/players")
	{
		playerRoutes.POST("/", middleware.RequireAuth, controllers.CreatePlayer)
		playerRoutes.GET("/", middleware.RequireAuth, controllers.GetAllPlayers)
		playerRoutes.GET("/:id", middleware.RequireAuth, controllers.GetPlayer)
		playerRoutes.PUT("/:id", middleware.RequireAuth, controllers.UpdatePlayer)
		playerRoutes.DELETE("/:id", middleware.RequireAuth, controllers.DeletePlayer)
	}
	// Rutas para inscripciones
	inscriptionRoutes := router.Group("/v1/inscription")
	{
		inscriptionRoutes.POST("/", middleware.RequireAuth, controllers.CreateInscription)
		inscriptionRoutes.GET("/", middleware.RequireAuth, controllers.GetAllInscriptions)
		inscriptionRoutes.GET("/:id", middleware.RequireAuth, controllers.GetInscription)
		inscriptionRoutes.PUT("/:id", middleware.RequireAuth, controllers.UpdateInscription)
		inscriptionRoutes.DELETE("/:id", middleware.RequireAuth, controllers.DeleteInscription)
	}

	personRoutes := router.Group("/v1/person")
	{
		// ... (rutas existentes)
		personRoutes.GET("/search", middleware.RequireAuth, controllers.SearchPerson)            // Nueva ruta de búsqueda
		personRoutes.GET("/is-available", middleware.RequireAuth, controllers.IsPersonAvailable) // Nueva ruta para verificar disponibilidad de jugadores
		personRoutes.GET("/current-user", middleware.RequireAuth, controllers.GetLoggedInPerson) // Nueva ruta para obtener el usuario logueado
	}

	// Rutas para equipos
	teamRoutes := router.Group("/v1/teams")
	{
		teamRoutes.POST("/", middleware.RequireAuth, controllers.CreateTeam)
		teamRoutes.GET("/:id", middleware.RequireAuth, controllers.GetTeam)
		teamRoutes.GET("/", middleware.RequireAuth, controllers.GetAllTeams)
		teamRoutes.PUT("/:id", middleware.RequireAuth, controllers.UpdateTeam)
		teamRoutes.DELETE("/:id", middleware.RequireAuth, controllers.DeleteTeam)
	}

	// Rutas para equipos
	imagesRoutes := router.Group("/v1/images")
	{
		imagesRoutes.POST("/", middleware.RequireAuth, controllers.UploadImage)
		imagesRoutes.GET("/:id", middleware.RequireAuth, controllers.GetImage)
		imagesRoutes.GET("/thumb/:id", middleware.RequireAuth, controllers.GetImageThumb)
		imagesRoutes.GET("/list", middleware.RequireAuth, controllers.ListImages)
	}

	return router

}
