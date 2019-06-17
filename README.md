[![Build Status](https://travis-ci.org/nammn/node-aggregation.svg?branch=master)](https://travis-ci.org/nammn/node-aggregation)
# node-aggregation

## General Description
This go program is made to get fast started in creating a node aggregation service in kubernetes.
Which includes testing on travis, letting it run on kubernetes as a deployment behind a load balancer and a kube status checks.

## Design
- Kubernetes 
- Stateless application 
- Horizontal scalable with k8 load balancing and readyness probes 
- State is handled by database service, which makes the service scalable 
- for now no proper security mechanism is deferred 
## Tradeoffs
- liveness probe currently uses the endpoint `/health` which has no proper code except returning a good health

## How to run
- Clone project
- Set minikube using local Docker/deploy on global docker
    - `minikube start`
    - `eval $(minikube docker-env)`
- Build Dockerimages and run Images
    - Redis DB
        - https://hub.docker.com/_/redis
        - `docker pull redis`
        - standard local port `6379`
        - `docker run -p 6379:6379 --name some-redis redis`
    - Webservice build Image
        - `docker build -t node-aggregation .`
- Deploy Webservice on Minikube 
    - go to `resources`
        `cd ./resources/`
    - Redis: 
        - `kubectl apply -f ./deployment_database.yaml`
    - Webservice: Deploy deployment File on Kubernetes/Minikube
        - `kubectl apply -f ./deployment.yaml` 
- Test using the provided ClientAPI 

## Making sure that it is running 
- using the integrationtests
    - `make integration-test`
- using the unit-test
    - `make unit-test`
- using the cobra client
    - `go build ./client`
    - `./client/client helloworld` -> should return `hello world`
    - `./client/client post` -> should return 
    `
    {
	"timeslice": 100,
	"cpu": 20,
	"mem": 20
	}
    ` as used in `POST`
    

### Packages
- Dep for dependency management
- Mux for router management
### Overall Architecture
[Picture](https://app.mural.co/t/icdretro7302/m/icdretro7302/1560497471665/f9a5f4324db66f10d3ff2991f8d60dbd8201e69d)
[Testing](https://app.mural.co/t/icdretro7302/m/icdretro7302/1560498264807/d2413ef2ad16a8a05fe019c4ade929077f9c26da)

## Tests
- Travis for unit/integrationstest
- Validate that the service is actually running

## Missing
- Proper security mechanism (usage of secrets and security indirections)
- Proper relation mapping between deployment of webservice and redis the DB
- Usage of service discovery in webservice to actually discover redis
    - currently the url is hardcoded as `localhost` 
    - further improvement would be the usage of an env
    - further improvement would even be service discovery
- Proper DB mapping/user handling
- Cobra integration is currently nothing more than a skeleton
- Integrationtest is currently only testing the database docker integration
