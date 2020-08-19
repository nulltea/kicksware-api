proxy:
	cd proxy-service;
	docker-compose down;
	docker-compose build;
	docker-compose push proxy-service;
	docker-compose up -d;
	
middleware:
	cd middleware-service;
	docker-compose down;
	docker-compose build;
	docker-compose push middleware-service;
	docker-compose up -d;

web-app:
	cd web-app-service;
	docker-compose down;
	docker-compose build;
	docker-compose push web-app;
	docker-compose up -d;

mongo-backup:
	docker exec mongo mongodump -u root -p greenJordans --authenticationDatabase admin --db=sneakerResaleDB  --out=/backup || echo "mongo down - buckup;

mongo-restore:
	docker exec mongo mongorestore -u root -p greenJordans --authenticationDatabase admin ./backup;

styles:
	mkdir web-app-service/Web/wwwroot/styles/css;
	for dir in web-app-service/Web/wwwroot/styles/less/*;
		do lessc-each web-app-service/Web/wwwroot/styles/less/$(basename ${dir}) web-app-service/Web/wwwroot/styles/css/$(basename ${dir});
	done;

cert:
	cd key;
	sh ./gen.sh;
	cd ../
