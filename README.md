# GoliPors

```bash
.
│─── cmd
│    └─── main.go
│─── api
│    │─── http
│    │    │─── server.go
│    │    └─── routes..
│    └─── service
│         └─── route services..
│─── config
│    │─── read.go   
│    └─── type.go  
│─── build   
│    │─── elk
│    │    │─── setup
│    │    │─── elasticsearch
│    │    │    │─── config
│    │    │    │    └─── elasticsearch.yml
│    │    │    └─── Dockerfile
│    │    │─── extensions
│    │    │    └─── fleet
│    │    │    │    │─── Dockerfile
│    │    │    │    └─── fleet-compose.yml
│    │    │─── logstash
│    │    │    │─── config
│    │    │    │    └─── logstash.yml
│    │    │    │─── pipeline
│    │    │    │    └─── logstash.conf
│    │    │    └─── Dockerfile
│    │    │─── kibana
│    │    │    │─── config
│    │    │    │    └─── kibana.yml
│    │    │    └─── Dockerfile
│    │    │─── docker-compose.yaml
│    │    └─── .env
│    │─── redis
│    │    ├─── docker-compose.yaml
│    │    └─── .env
│    │─── postgres
│    │    ├─── docker-compose.yaml
│    │    └─── .env
│    └─── project
│         ├─── Dockerfile
│         ├─── docker-compose.yaml
│         └─── .env
│─── internal
│    └─── services...
│─── pkg
│    ├─── adapters
│    │    ├─── cache
│    │    │    └─── redis.go
│    │    └─── storage
│    │         │─── mapper
│    │         │─── types
│    │         └─── repository files...
│    ├─── postgres
│    │    └─── gorm.go
│    ├─── cache
│    │    ├─── provider.go
│    │    └─── serialization.go
│    ├─── logger
│    └─── jwt
│         ├─── claims.go
│         └─── auth.go
│─── tests
│─── config.json
│─── sample-config.json
│─── go.mod
│─── go.sum
│─── README.md
│─── LICENSE
│─── .gitignore
└─── Makefile
```