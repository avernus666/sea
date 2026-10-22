package main

import (
 "context"
 "log"
 "net/http"
 "os"
 "os/signal"
 "syscall"
 "time"
 "sea/backend/internal/config"
 "sea/backend/internal/httpapi"
 "sea/backend/internal/repository"
)

func main() {
 cfg := config.Load()
 repo, err := repository.NewPostgresRepository(context.Background(), cfg.DatabaseURL)
 if err != nil { log.Fatal("database connection failed") }
 defer repo.Close()
 server := &http.Server{Addr: ":"+cfg.Port, Handler: httpapi.NewHandler(repo,cfg.AllowedOrigin).Routes(), ReadHeaderTimeout:5*time.Second, ReadTimeout:10*time.Second, WriteTimeout:15*time.Second, IdleTimeout:60*time.Second}
 stop:=make(chan os.Signal,1); signal.Notify(stop,syscall.SIGINT,syscall.SIGTERM)
 go func(){ <-stop; ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second); defer cancel(); _=server.Shutdown(ctx) }()
 log.Printf("API listening on :%s",cfg.Port)
 if err:=server.ListenAndServe(); err!=nil && err!=http.ErrServerClosed { log.Fatal("http server failed") }
}
