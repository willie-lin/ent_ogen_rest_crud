package main

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/mattn/go-sqlite3"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/ent"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/ent/ogent"
	"log"
	"net/http"
)

func main() {

	// Create ent client
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	// Run the migrations.
	if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
		log.Fatal(err)
	}

	// Start listening
	srv, _ := ogent.NewServer(ogent.NewOgentHandler(client))
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}



}
