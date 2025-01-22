package server

import (
	"ipadel-club/controllers"
	"ipadel-club/middleware"
	"ipadel-club/ui"
	"strings"

	"github.com/gin-gonic/gin"
)

/*
curl -i \
-H "Origin: http://localhost:3000" \
-H 'Access-Control-Request-Method: GET' \
-H 'Access-Control-Request-Headers: Content-Type, Authorization' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzcxNzQwMjAsInN1YiI6MX0.jUW3UQHWzU8DWT9Al774ozjtiYSKnd2szitBH-bc2T8' \
"https://ipadel-club-be.acesgdl.com/v1/catalogs/clubs"

https://ipadel-club-js-development.up.railway.app
*/ /*
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://ipadel-club-be.acesgdl.com")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Origin, Accept, X-Custom-Header, Range")
		c.Writer.Header().Set("Expose-Headers", "Content-Length, Content-Type, Authorization, X-Requested-With, Origin, Accept, X-Custom-Header, Range")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
} */

// CORS middleware function definition
func corsMiddleware() gin.HandlerFunc {
	// Define allowed origins as a comma-separated string
	originsString := "http://localhost:3000, https://ipadel-club-js-development.up.railway.app"
	var allowedOrigins []string
	if originsString != "" {
		// Split the originsString into individual origins and store them in allowedOrigins slice
		allowedOrigins = strings.Split(originsString, ",")
	}

	// Return the actual middleware handler function
	return func(c *gin.Context) {
		// Function to check if a given origin is allowed
		isOriginAllowed := func(origin string, allowedOrigins []string) bool {
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}
			return false
		}

		// Get the Origin header from the request
		origin := c.Request.Header.Get("Origin")

		// Check if the origin is allowed
		if isOriginAllowed(origin, allowedOrigins) {
			// If the origin is allowed, set CORS headers in the response
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		}

		// Handle preflight OPTIONS requests by aborting with status 204
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Call the next handler
		c.Next()
	}
}

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Recovery())

	// ruta para servir el front-end en ReactJS
	ui.AddRoutes(router)

	// Security
	router.POST("/auth/register", controllers.Signup)
	router.POST("/auth/login", controllers.Login)
	router.POST("/auth/logout", controllers.Logout)

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
		categoryRoutes.GET("/", controllers.GetAllCategories)
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
	}

	return router

}
