proxy:
	docker-compose -f proxy-service/docker-compose.yml down;
	docker-compose -f proxy-service/docker-compose.yml build;
	docker-compose -f proxy-service/docker-compose.yml up -d;

middleware:
	docker-compose -f middleware-service/docker-compose.yml down;
	docker-compose -f middleware-service/docker-compose.yml build;
	docker-compose -f middleware-service/docker-compose.yml push users-service references-service products-service search-service;
	docker-compose -f middleware-service/docker-compose.yml up -d;

web-app:
	docker-compose -f web-app-service/docker-compose.yml down;
	docker-compose -f web-app-service/docker-compose.yml build;
	docker-compose -f web-app-service/docker-compose.yml push web-app;
	docker-compose -f web-app-service/docker-compose.yml up -d;

mongo-backup:
	docker exec mongo mongodump -u $(MONGO_USER) -p $(MONGO_PASSWORD) --authenticationDatabase admin --db=sneakerResaleDB  --out=/backup || echo "mongo down - backup impossible";

mongo-restore:
	docker exec mongo mongorestore -u $(MONGO_USER) -p $(MONGO_PASSWORD) --authenticationDatabase admin ./backup;

styles:
	mkdir web-app-service/Web/wwwroot/styles/css;
	for dir in web-app-service/Web/wwwroot/styles/less/*; do \
		lessc-each web-app-service/Web/wwwroot/styles/less/$(basename ${dir}) web-app-service/Web/wwwroot/styles/css/$(basename ${dir}); \
	done;

cert:
	cd key;
	sh ./gen.sh;
	cd ../
