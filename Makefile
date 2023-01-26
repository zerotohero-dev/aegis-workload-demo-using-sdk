#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                  <aegis.z2h.dev>
#     .\_/.
#

VERSION=0.9.34
PACKAGE=aegis-workload-demo-using-sdk
REPO=z2hdev/aegis-workload-demo-using-sdk

all: build bundle push deploy

build:
	go build -o ${PACKAGE}

run:
	go run main.go

bundle:
	go mod vendor
	docker build . -t ${PACKAGE}:${VERSION}

push:
	docker build . -t ${PACKAGE}:${VERSION}
	docker tag ${PACKAGE}:${VERSION} ${REPO}:${VERSION}
	docker push ${REPO}:${VERSION}

deploy:
	kubectl apply -f ./k8s/ServiceAccount.yaml
	kubectl apply -f ./k8s/Deployment.yaml
	kubectl apply -f ./k8s/Identity.yaml

run-in-container:
	docker run ${PACKAGE}:${VERSION}
