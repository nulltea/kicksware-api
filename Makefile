api:
	docker-compose -f middleware-service/docker-compose.yml down;
	docker-compose -f middleware-service/docker-compose.yml build;
	docker-compose -f middleware-service/docker-compose.yml push users-service references-service products-service search-service;
	docker-compose -f middleware-service/docker-compose.yml up -d;

mongo-backup:
	docker exec mongo mongodump -u $(MONGO_USER) -p $(MONGO_PASSWORD) --authenticationDatabase admin --db=sneakerResaleDB  --out=/backup || echo "mongo down - backup impossible";

mongo-restore:
	docker exec mongo mongorestore -u $(MONGO_USER) -p $(MONGO_PASSWORD) --authenticationDatabase admin ./backup;

cert:
	cd key;
	sh ./gen.sh;
	cd ../
