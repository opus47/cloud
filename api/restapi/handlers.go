package restapi

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/opus47/cloud/api/models"
	"github.com/opus47/cloud/api/restapi/operations"
)

func handleGetPiecesSearch(
	params operations.GetPiecesSearchParams,
) middleware.Responder {
	session, err := mgo.Dial("db")
	if err != nil {
		log.Printf("dial error: %v", err)
		return operations.NewGetPiecesSearchInternalServerError()
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("opus47").C("pieces")
	result := []*models.Piece{}
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Printf("find error: %v", err)
		return operations.NewGetPiecesSearchOK().WithPayload(result)
	}

	return operations.NewGetPiecesSearchOK().WithPayload(result)
}
