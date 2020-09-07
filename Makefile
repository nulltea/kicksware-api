api:
	docker-compose build;
	docker-compose down;
	docker-compose up -d;
cert:
	cd key;
	sh ./gen.sh;
	cd ../

