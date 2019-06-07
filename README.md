# KTBCS-Container-PoC-ServiceB2
#### Containerized Applications for CICD pipeline for DevOps team @G-Able<br />Developed by Phuwanai Thummavet <www.serial-coder.com>

## Services
1. ServiceB2 - Go HTTP Proxy Server connecting to another external service (e.g., http://httpbin.org)

## To compile ServiceB2 locally
```
cd KTBCS-Container-PoC-ServiceB2
go build
```

## To build Docker image and deploy container using Docker Compose
```
cd KTBCS-Container-PoC-ServiceB2
docker-compose up --build
```

## To access to ServiceB2 (w/ several types of example requests)
```
http://localhost:32000
http://localhost:32000/status/200
http://localhost:32000/status/404
http://localhost:32000/status/500
http://localhost:32000/status/503
http://localhost:32000/delay/2
http://localhost:32000/delay/10
```

## To tear down the deployed container using Docker Compose
```
cd KTBCS-Container-PoC-ServiceB2
docker-compose down
```
"# ktbserviceb2" 
