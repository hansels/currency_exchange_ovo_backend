package usecase

import "cloud.google.com/go/firestore"

type Opts struct {
	Firestore *firestore.Client
}

type Module struct {
	Firestore *firestore.Client
}

func New(opts *Opts) *Module {
	return &Module{Firestore: opts.Firestore}
}
