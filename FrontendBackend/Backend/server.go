
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
    "io/ioutil"
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

    // send should only be used for POST requests
    router.HandleFunc("/send", formHandler)

    server := &http.Server{
      Addr:         listenAddr,
      Handler:      router,
      ErrorLog:     logger,
      ReadTimeout:  5 * time.Second,
      WriteTimeout: 10 * time.Second,
      IdleTimeout:  15 * time.Second,
    }

    go gracefullShutdown(server, logger, sigs, done)
    
    //syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

    logger.Println("Server is ready to handle requests at", listenAddr)
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
    }
    logger.Println("Server stopped")
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

func formHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
		case "GET":
			for k, v := range r.URL.Query() {
				fmt.Printf("%s: %s\n", k, v)
			}
			w.Write([]byte("Received a GET request\n"))
		case "POST":
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Received a POST Request")
			fmt.Printf("%s\n", reqBody)
			w.Write([]byte("POST Sucessfully recieved\n"))
		default:
			w.WriteHeader(http.StatusNotImplemented)
			w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}