package main

import (
	"context"
	"fmt"
	"github.com/bykof/gostradamus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/config"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/database"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/database/ent"
	"github.com/willie-lin/ent_ogen_rest_crud/ents/database/ent/ogent"

	"go.uber.org/zap"
	"net/http"
)

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
}

func (h handler) MarkDone(ctx context.Context, params ogent.MarkDoneParams) (ogent.MarkDoneNoContent, error) {
	return ogent.MarkDoneNoContent{}, h.client.Todo.UpdateOneID(params.ID).SetDone(true).Exec(ctx)
}

func main() {

	log, _ := zap.NewDevelopment()

	e := echo.New()

	//e.Use(apmechov4.Middleware())
	//e.Use(logger.ZapLogger(log))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	//连接 数据库
	client, err := database.NewClient()
	//fmt.Println(client)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
		return
	}

	defer client.Close()

	//fmt.Println(client)
	dateTime := gostradamus.Now()
	fmt.Println(dateTime)

	//defer client.Close()
	ctx := context.Background()

	//autoMigration := database.AutoMigration
	autoMigration := database.AutoMigration
	autoMigration(client, ctx)
	//
	//debugMode := database.DebugMode
	debugMode := database.DebugMode
	//
	debugMode(err, client, ctx)

	//controller := &handler.Controller{Client: client}

	//fmt.Println(client)
	if err != nil {
		log.Fatal("opening ent client", zap.Error(err))
		return
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//// Create ent client
	//client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Run the migrations.
	//if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
	//	log.Fatal(err)
	//}
	//
	//h := handler{
	//	OgentHandler: ogent.NewOgentHandler(client),
	//	client:       client,
	//}
	//
	//// Start listening
	//srv, _ := ogent.NewServer(h)
	//if err := http.ListenAndServe(":8080", srv); err != nil {
	//	log.Fatal(err)
	//}

	e.Logger.Fatal(e.Start(":1323"))
}

func init() {
	config.InitConfig()
}
