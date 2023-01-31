package route

import (
	gw_handler "Go_cinema_clean_arch/gateway/handler"
	// movie_pb "Go_cinema_clean_arch/movie/proto"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	//"go.mongodb.org/mongo-driver/mongo"
)

func Private(e *echo.Echo, handler *gw_handler.GatewayHandler, conn *grpc.ClientConn) {
	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUser)
	e.POST("/users", handler.AddUser)
	e.PUT("/users/:id", handler.UpdateUser)
	e.DELETE("/users/:id", handler.DeleteUser)
	e.GET("/users/:id/watched_movies", handler.GetWatchedMovies)

	e.GET("/movies", handler.GetMovies)
	e.GET("/movies/:id", handler.GetMovie)
	e.POST("/movies", handler.AddMovie)
	e.PUT("/movies/:id", handler.UpdateMovie)
	e.DELETE("/movies/:id", handler.DeleteMovie)
}
