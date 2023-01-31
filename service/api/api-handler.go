package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	rt.router.POST("/session", rt.wrap(rt.doLogin))          // OK ok
	rt.router.PUT("/profile/:id", rt.wrap(rt.setMyUserName)) // OK ok

	rt.router.PUT("/users/:id1/follow/:id2", rt.wrap(rt.followUser))      // OK ok
	rt.router.DELETE("/users/:id1/follow/:id2", rt.wrap(rt.unfollowUser)) // OK ok

	rt.router.POST("/profile/:id/photo", rt.wrap(rt.uploadPhoto))        // OK ok
	rt.router.DELETE("/profile/:id/photo/:idp", rt.wrap(rt.deletePhoto)) // OK ok

	rt.router.PUT("/photo/:idp/comment/:idc", rt.wrap(rt.commentPhoto))      // OK ok
	rt.router.DELETE("/photo/:idp/comment/:idc", rt.wrap(rt.uncommentPhoto)) // OK
	rt.router.GET("/photo/:idp/comments", rt.wrap(rt.getComments))           // OK ok

	rt.router.PUT("/photo/:idp/like/:ido", rt.wrap(rt.likePhoto))      // OK ok
	rt.router.DELETE("/photo/:idp/like/:ido", rt.wrap(rt.unlikePhoto)) // OK ok

	rt.router.PUT("/users/:id1/ban/:id2", rt.wrap(rt.banUser))      // OK ok
	rt.router.DELETE("/users/:id1/ban/:id2", rt.wrap(rt.unbanUser)) // OK ok

	rt.router.GET("/search/:username", rt.wrap(rt.getUserProfile)) // OK ok
	rt.router.GET("/stream", rt.wrap(rt.getMyStream))              // OK

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
