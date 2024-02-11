package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	http.Handle("/api/users/", myHandler("User service"))
	http.Handle("api/preferences/", myHandler("Preferences service"))
	http.Handle("/api/user-interests/", myHandler("User Interests service"))
	http.Handle("/api/messages/", myHandler("Message service"))
	http.Handle("/api/matches/", myHandler("Matches service"))
	http.Handle("/api/photos/", myHandler("Photos service"))
	http.Handle("/api/swipes/", myHandler("Swipes service"))

	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))

	// serve content
	http.HandleFunc("/servecontent", func(w http.ResponseWriter, r *http.Request) {
		customerFile, err := os.Open("./customers.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer customerFile.Close()

		http.ServeContent(w, r, "customersdata.csv", time.Now(), customerFile)
	})

	// serve file to download
	http.HandleFunc("/servefile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./customers.csv")
	})

	// read file
	http.HandleFunc("/fprint", func(w http.ResponseWriter, r *http.Request) {
		customerFile, err := os.Open("customers.csv")
		if err != nil {
			log.Fatal(err)
		}
		defer customerFile.Close()
		// data, err := io.ReadAll(customerFile)
		// if err != nil {
		// log.Fatal(err)
		// }
		// fmt.Fprint(w, string(data))
		io.Copy(w, customerFile)
	})

	var handlerFunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.String())
	}

	http.HandleFunc("/api/", handlerFunc)

	s := http.Server{
		Addr: ":3000",
	}

	go func() {
		log.Fatal(s.ListenAndServeTLS("./cert.pem", "key.pem"))
	}()

	fmt.Println("Listening on port 3000..., press <Enter> to stop")
	fmt.Scanln()
	s.Shutdown(context.Background())
	fmt.Println("Server stopped")
}

type myHandler string

func (my myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Golang")
	http.SetCookie(w, &http.Cookie{
		Name:    "session-id",
		Value:   "123",
		Expires: time.Now().Add(time.Hour * 24 * 365),
	})
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, string(my))
	fmt.Fprintln(w, r.Header)
}
