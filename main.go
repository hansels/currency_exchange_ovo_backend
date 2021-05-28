package main

import (
	"github.com/hansels/currency_exchange_ovo_backend/src/firebase"
	"github.com/hansels/currency_exchange_ovo_backend/src/server"
	"github.com/hansels/currency_exchange_ovo_backend/src/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	firestore := firebase.InitFirestore()

	opts := &usecase.Opts{Firestore: firestore}
	modules := usecase.New(opts)

	api := server.New(&server.Opts{ListenAddress: ":3000", Modules: modules})

	go api.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-term:
		log.Println("Exiting gracefully...", s)
	}
	log.Println("👋")

	return 0
}
