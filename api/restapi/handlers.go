package restapi

import (
	"log"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"database/sql"
	_ "github.com/lib/pq"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/opus47/cloud/api/models"
	"github.com/opus47/cloud/api/restapi/operations"
)

//TODO clearly this needs to hit a memcache and not mongo directly....
func _handleGetPiecesSearch(
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

func handleGetPiecesSearch(
	params operations.GetPiecesSearchParams,
) middleware.Responder {

	connStr := "host=172.17.0.2 user=postgres dbname=opus47 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("pg-connect error: %v", err)
		return operations.NewGetPiecesSearchInternalServerError()
	}
	defer db.Close()

	result := []*models.Piece{}

	if params.Text == nil {
		return operations.NewGetPiecesSearchOK().WithPayload(result)
	}

	log.Printf("[Pieces-Search] text: %s", *params.Text)

	// build the search string in the format
	// one for the rhythm ---> one:*&for:*&the:*&rhythm:*
	search_string := ""
	fields := strings.Fields(*params.Text)
	for i := 0; i < len(fields)-1; i++ {
		search_string += fields[i] + ":*&"
	}
	search_string += fields[len(fields)-1] + ":*"
	log.Printf("[Pieces-Search] query: %s", search_string)

	rows, err := db.Query(`
	select pid, cname, ptitle, kname, pcatalog from mv_pieces
	where document @@ to_tsquery('english', '` + search_string + `')
	`)
	if err != nil {
		log.Printf("pg-querry error: %v", err)
		return operations.NewGetPiecesSearchInternalServerError()
	}

	defer rows.Close()
	for rows.Next() {
		p := &models.Piece{}
		err := rows.Scan(&p.ID, &p.Composer, &p.Title, &p.Key, &p.Catalog)
		if err != nil {
			log.Printf("pg-scan error: %v", err)
			return operations.NewGetPiecesSearchInternalServerError()
		}
		result = append(result, p)
	}

	return operations.NewGetPiecesSearchOK().WithPayload(result)
}
