package config

var Config = configuration{
	MONGOURL:     "mongodb://admin:admin@127.0.0.1:27017",
	DATABASENAME: "mongo-kafka-go",
	TOPIC:        "tesodev-go.orders",
}

type configuration struct {
	MONGOURL     string
	DATABASENAME string
	TOPIC        string
}
