api:
	docker-compose build;
	docker-compose down;
	docker-compose up -d;
cert:
	cd keys;
	sh gen.sh;
	cd ../

