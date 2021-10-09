include .env
export

cert:
	cd keys;
	sh gen.sh;
	cd ../

elasticsearch:
	helm repo add elastic https://helm.elastic.co
	helm install elasticsearch elastic/elasticsearch -f deploy/config/elasticsearch.yaml

mongodb:
#	helm upgrade --install mongodb \
#		--set auth.password=${MONGO_PASSWORD} -f ./deploy/config/mongodb.yaml bitnami/mongodb
	kubectl create secret generic mongo-auth \
 		--from-literal=MONGO_USER=${MONGO_USER} --from-literal=MONGO_PASSWORD=${MONGO_PASSWORD} \
 		--dry-run=client -o yaml | kubectl apply -f -

build:
	docker build -f ./services/${service}/docker/Dockerfile . -t timothydockid/kicksware-${service}-service
	docker push timothydockid/kicksware-${service}-service

install:
	envsubst < services/${service}/env/config.${ENV}.yaml > config.yaml && kubectl create configmap ${service}-service.config \
			--from-file=config.yaml=config.yaml --dry-run=client -o yaml | kubectl apply -f -
	helm upgrade --install ${service} services/${service}/${service}-chart

setup-k8s:
	kubectl create secret tls grpc-tls --key=keys/server.key --cert=keys/server.crt --dry-run=client -o yaml | kubectl apply -f -
	kubectl create secret generic auth-keys --from-file=private.key=services/users/key/private.key --from-file=public.key=keys/public.key.pub
	kubectl create configmap mail-service.templates \
		--from-file=services/users/template/verify.template.html \
		--from-file=services/users/template/notify.template.html \
		--from-file=services/users/template/reset.template.html --dry-run=client -o yaml | kubectl apply -f -

sync-source:
	rsync -r -v . --exclude .git ubuntu@${REMOTE_IP}:kicksware/api

grpc-tls-gen:
	openssl genrsa \
		-out keys/ca.key 2048

	openssl req \
		-subj "/C=UA/ST=Kiev/O=Kicksware, Inc./CN=rpc.kicksware.com" \
		-new -x509 -days 365 -key keys/ca.key -out keys/ca.crt

	openssl req -newkey rsa:2048 \
		-nodes -keyout keys/server.key \
		-subj "/C=UA/ST=Kiev/O=Kicksware, Inc./CN=rpc.kicksware.com" \
		-out keys/server.csr

	openssl x509 -req \
		-in keys/server.csr \
		-CA keys/ca.crt -CAkey keys/ca.key -CAcreateserial -days 365 \
		-extfile <(printf "subjectAltName=DNS:rpc.kicksware.com,DNS:localhost") \
		-out keys/server.crt
