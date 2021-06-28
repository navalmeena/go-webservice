package main

import (
	"log"
	"fmt"
	"net/http"
	"go-curd-ops/pkg/routes"
	"github.com/gorilla/mux"
	"os"
	// "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {

	httpPort := os.Getenv("PORT")

	logPath := os.Getenv("LOGPATH")

	fmt.Println("Running server......")
	
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	routes.RegisterUsersRoutes(r)
	http.Handle("/", r)
	
	fmt.Printf("listening on %v\n", httpPort)
	fmt.Printf("Logging to %v\n", logPath)


	log.Fatal(http.ListenAndServe("localhost:"+httpPort, logRequest(r)))

	// err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), )
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1><div>Welcome to whereever you are</div>")
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}