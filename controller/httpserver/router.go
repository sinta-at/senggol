package httpserver

import (
	"senggol/controller/httpserver/handler"
	"senggol/controller/httpserver/middleware"
	"senggol/service"

	"github.com/julienschmidt/httprouter"
)

func makeRouter(services service.Services) *httprouter.Router {
	router := httprouter.New()

	basicAuth := middleware.AuthMiddleware{services.GetAuthenticatedUser}

	register := handler.RegisterHandler{services.Register}
	getPeers := handler.GetPeersHandler{services.GetPeers}
	postDirectMessage := handler.PostDirectMessageHandler{services.PostDirectMessage}
	getDirectMessages := handler.GetDirectMessagesHandler{services.GetDirectMessages}
	setMessageSeen := handler.SetMessageSeenHandler{services.SetMessageSeen}

	router.POST("/register", register.HandleRegister)
	router.GET("/peers", basicAuth.BasicAuth(getPeers.HandleGetPeers))
	router.POST("/peers/:id/messages", basicAuth.BasicAuth(postDirectMessage.HandlePostDirectMessage))
	router.GET("/peers/:id/messages", basicAuth.BasicAuth(getDirectMessages.HandleGetDirectMessages))
	router.PATCH("/messages/:id/seen", basicAuth.BasicAuth(setMessageSeen.HandleSetMessageSeen))

	return router
}