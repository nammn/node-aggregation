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
- Build Dockerimages
    - Redis DB
        - https://hub.docker.com/_/redis
        - standard local port `6379`
        - `docker run -p 6379:6379 --name some-redis redis`
    - Webservice
- Set minikube using local Docker/deploy on global docker
- Deploy deployment File on Kubernetes/Minikube + Deployment
- Test using the provided ClientAPI 
## Versions

### Packages
- Dep for dependency management
- Mux for router management
### Overall Architecture
[Picture](https://app.mural.co/t/icdretro7302/m/icdretro7302/1560497471665/f9a5f4324db66f10d3ff2991f8d60dbd8201e69d)
[Testing](https://app.mural.co/t/icdretro7302/m/icdretro7302/1560498264807/d2413ef2ad16a8a05fe019c4ade929077f9c26da)


## API 

## Tests
- Travis for unit/integrationstest
- Validate that the service is actually running
## Running it

