package segmentationService

import 
(
	"net/http"
	"context"
	"time"
) 

type Server struct
{
	httpServer *http.Server
}

// Запуск Сервера
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server
	{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Выход из приложения
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}


