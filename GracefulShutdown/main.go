package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
    "syscall"
    "fmt"
)

var (
    listenAddr string
)

func main() {

    listenAddr := ":8080"
    logger := log.New(os.Stdout, "http: ", log.LstdFlags)

    //logger.Println("pid: %dn", os.Getpid())

    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    router := http.NewServeMux() 
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      w.WriteHeader(http.StatusOK)
      fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    })

    server := &http.Server{
      Addr:         listenAddr,
      Handler:      router,
      ErrorLog:     logger,
      ReadTimeout:  5 * time.Second,
      WriteTimeout: 10 * time.Second,
      IdleTimeout:  15 * time.Second,
    }

    go gracefullShutdown(server, logger, sigs, done)


    shutdownTimer()
    
    //syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

    logger.Println("Server is ready to handle requests at", listenAddr)
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
    }
    logger.Println("Server stopped")
}

func raise(sig os.Signal) error {
  p, err := os.FindProcess(os.Getpid())
  if err != nil {
    return err
  }
  return p.Signal(sig)
}

// raises SIGTERM signal after 10 seconds, which will cause the web server to shut down gracefully
func shutdownTimer() {
        time.AfterFunc(10*time.Second, func() { 
            raise(syscall.SIGTERM) 
        })
    }

func gracefullShutdown(server *http.Server, logger *log.Logger, sigs <-chan os.Signal, done chan<- bool) {
  <-sigs
  logger.Println("Server is shutting down...")

  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  server.SetKeepAlivesEnabled(false)
  if err := server.Shutdown(ctx); err != nil {
    logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
  }
  close(done)
}