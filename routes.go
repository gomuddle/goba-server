package gobaserver

// routes configures s's endpoints.
func (s Server) routes() {
	s.router.GET("/images/{type}/{name}", s.handle(s.authMiddleware(s.getImage())))
	s.router.GET("/images/{type}", s.handle(s.authMiddleware(s.getAllImages())))
	s.router.POST("/images/{type}/{name}", s.handle(s.authMiddleware(s.applyImage())))
	s.router.POST("/images/{type}", s.handle(s.authMiddleware(s.createImage())))
	s.router.DELETE("/images/{type}/{name}", s.handle(s.authMiddleware(s.deleteImage())))
}
