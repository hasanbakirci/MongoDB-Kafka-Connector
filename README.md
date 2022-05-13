## MongoDB Kafka Connector

1 - `docker-compose up -d`

2 - `docker exec -it <container_id> /bin/bash`

3 - `confluent-hub install mongodb/kafka-connect-mongodb:1.7.0`

4 -  Get `localhost:9021`

5 -  Connect  > Connect-Default > Add Connector > Mongo Source Connector > Set MongoDB Connection Url, Database, collection

6 - `./config/config.go` Set TOPIC: `<database>.<collection>`