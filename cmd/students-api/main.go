package main

import (
	"context"
	
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Saiteja93/Students-Api-Go/internal/config"
)

func main() {
    // load config
    cfg:= config.MustLoad()

    //database setup

    //setup router
    router:= http.NewServeMux()
    router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("welcome to student api, it is running")) 
    })

    //setup server
    server := http.Server{
        Addr: cfg.Httpserver.Addr,
        Handler: router,
    }
    slog.Info("Server started", "address", cfg.Httpserver.Addr)

    done := make(chan os.Signal, 1)

    signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)


    go func ()  {
        err := server.ListenAndServe()
        if err != nil{
            log.Fatal("failed to start server")
        
        }
    }()

    <- done

    slog.Info("SHUTTING DOWN THE SERVER")

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()


    if err := server.Shutdown(ctx); err != nil {
        slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
    }

  

    slog.Info("Server shutdown successfully")
   

  
 }

    

