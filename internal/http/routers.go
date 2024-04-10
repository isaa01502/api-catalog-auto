package http

func (s *server) routers() {
	s.router.GET("/health")

	v1 := s.router.Group("/api/v1")
	{
		v1.GET("catalog", s.handlers.catalog.GetCatalog)
		v1.POST("catalog", s.handlers.catalog.AddCatalog)
		v1.PUT("catalog", s.handlers.catalog.UpdateCatalog)
		v1.DELETE("catalog", s.handlers.catalog.DeleteCatalog)
	}
}
