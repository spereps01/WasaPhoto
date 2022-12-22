package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	rt.router.POST("/session", rt.wrap(rt.doLogin))          // OK
	rt.router.PUT("/profile/:id", rt.wrap(rt.setMyUserName)) // OK

	rt.router.PUT("/users/:id1/follow/:id2", rt.wrap(rt.followUser))      // OK
	rt.router.DELETE("/users/:id1/follow/:id2", rt.wrap(rt.unfollowUser)) // OK

	rt.router.POST("/photo", rt.wrap(rt.addPhoto))                       // OK
	rt.router.DELETE("/profile/:id/photo/:idp", rt.wrap(rt.deletePhoto)) // OK

	rt.router.PUT("/photo/:idp/comment/:idc", rt.wrap(rt.commentPhoto))      // OK
	rt.router.DELETE("/photo/:idp/comment/:idc", rt.wrap(rt.uncommentPhoto)) // OK

	rt.router.PUT("/photo/:idp/like/:ido", rt.wrap(rt.likePhoto))      // OK
	rt.router.DELETE("/photo/:idp/like/:ido", rt.wrap(rt.unlikePhoto)) // OK

	rt.router.PUT("/b_users/:idu/ban/:idb", rt.wrap(rt.banUser))      // OK
	rt.router.DELETE("/b_users/:idu/ban/:idb", rt.wrap(rt.unbanUser)) // OK

	rt.router.GET("/search/:username", rt.wrap(rt.getUserProfile)) // OK
	// rt.router.GET("/stream", rt.wrap(rt.getMyStream))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
