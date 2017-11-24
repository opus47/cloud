package restapi

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/opus47/cloud/api/models"
	"github.com/opus47/cloud/api/restapi/operations"
)

//TODO clearly this needs to hit a memcache and not mongo directly....
func handleGetPiecesSearch(
	params operations.GetPiecesSearchParams,
) middleware.Responder {
	session, err := mgo.Dial("172.17.0.2")
	if err != nil {
		log.Printf("dial error: %v", err)
		return operations.NewGetPiecesSearchInternalServerError()
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	pc := session.DB("opus47").C("pieces")
	result := []*models.Piece{}
	err = pc.Find(bson.M{}).All(&result)
	if err != nil {
		log.Printf("find error: %v", err)
		return operations.NewGetPiecesSearchInternalServerError()
	}

	//resolve composers
	cc := session.DB("opus47").C("composers")
	for _, x := range result {
		composer := models.Composer{}
		err = cc.Find(bson.M{"@id": x.Composer}).One(&composer)
		if err != nil {
			log.Printf("find composer error: %s - %v", x.Composer, err)
			return operations.NewGetPiecesSearchInternalServerError()
		}
		x.Composer = composer.Name
	}

	//resolve keys
	kc := session.DB("opus47").C("keys")
	for _, x := range result {
		key := models.Key{}
		err = kc.Find(bson.M{"@id": x.Key}).One(&key)
		if err != nil {
			log.Printf("find key error: %s - %v", x.Key, err)
			return operations.NewGetPiecesSearchInternalServerError()
		}
		x.Key = key.Name
	}

	return operations.NewGetPiecesSearchOK().WithPayload(result)
}
