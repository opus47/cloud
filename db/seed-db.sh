#!/bin/sh

docker exec opus47-db mongoimport --db opus47 --jsonArray --collection keys --file /data47/keys.json
docker exec opus47-db mongoimport --db opus47 --jsonArray --collection composers --file /data47/composers.json
docker exec opus47-db mongoimport --db opus47 --jsonArray --collection pieces --file /data47/brahms-pieces.json
docker exec opus47-db mongoimport --db opus47 --jsonArray --collection pieces --file /data47/dvorak-pieces.json
docker exec opus47-db mongoimport --db opus47 --jsonArray --collection musicians --file /data47/musicians.json
docker exec opus47-db mongoimport --db opus47 --jsonArray --collection performances --file /data47/olyfest-perf.json
docker exec opus47-db mongoimport --db opus47 --jsonArray --collection recordings --file /data47/olyfest-rec.json
docker exec opus47-db mongoimport --db opus47 --jsonArray --collection parts --file /data47/parts.json
