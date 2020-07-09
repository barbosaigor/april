# April
**April is under construction**  

April proposes improve resilience in microservices architectures. It
does chaos testing by randomly shutting down services, taking into account 
their importance.  
April is a CLI tool, being possible to either run chaos testing or others tools,
such as 'server' which hosts an API for remote access to the chaos test.
The selection algorithm firstly picks K services, then it picks N dependencies on those services. Heavy services are more likely to be picked.  
It's important that Chaos Server is running to terminate instances.  

## Installation  
```bash 
go get -u github.com/barbosaigor/april/...
```   

## Tools
Chaos test. 
Need a running 'chaos server' to terminate instances.  
-f configuraion file path (Default is conf.yml)  
-n maximum number of services to choose  
-c chaos server endpoint (Default is localhost:7071)  
-u username for chaos server auth  
-s password for chaos server auth  
```bash 
april -f conf.yml -n 10 -u bob -s mysecret
```  

*Bare* runs only the selection algorithm, returning a set of services.  
-f configuraion file path  
-n maximum number of services to choose  
```bash 
april bare -f conf.yml -n 10  
```  

*Server* hosts an API (HTTP) which apply chaos testing and bare algorithm.
Need a running 'chaos server' to terminate instances.  
-p port number (Default is 7070)  
-c chaos server endpoint (Default is localhost:7071)  
```bash 
april server -p 8080  // will listen on port 8080
``` 
## Configuration file
*Fields*  
_version_: could be ommited, default is 1   
_services_: describes a list of services  
_servicename_: is the name of a service that April is going to work  
_weight_: represents the service importance for the April pick algorithm  
_depedencies_: describe a list of services which the service depends  
_selector_: describe how chaos server must search for the service name. 
E.g if you are using docker containers and a framework such as docker compose,
compose will define the container name as a concatenation between your service name and a hash somewhere, in this case, it is better to look for the _infix_ corresponding to the service.  
```yaml
# template
version: 1
services:
    servicename:
        weight: [0-9]+ (any natural number)
        dependencies:
            - [a-zA-Z\_\-]* (dependency name)
        selector: prefix|infix|postfix|all (how should match the service name instance)
```  

*Example conf.yaml*  
```yaml
version: 1
services:
  payment:
    weight: 10
    dependencies:
      - profile
      - fees
    selector: postfix  

  fees:
    weight: 5
    selector: infix  

  profile:
    weight: 20
    selector: infix  

  inventory:
    weight: 15
    selector: infix  

  shipping:
    weight: 5
    dependencies:
      - inventory
      - profile
    selector: infix  

  storefront:
    weight: 20
    dependencies:
      - shipping
      - inventory
      - profile
      - payment
      - fees
    selector: infix
```

## Design approach 
The Aprils design is divided into two parts: CLI and Chaos server. CLI runs the algorithm and request to Chaos server for terminate instances. 
Therefore, we gain flexibility about technologies that manage instances, a Chaos server could terminate Docker containers, Kubernetes instances etc.  

![Aprils design](./res/aprils-diagram-1.png)  

## What is a Chaos server ?
Chaos server hosts an API that terminantes instances. Aprils CLI runs its algorithm and asks the Chaos server to finish 
the selected instances. The API implementation lives in april/destroyer package, so chaos servers must include that package and
implement the Destroyer interface, which contain the business logic for terminate instances. 

## Chaos Servers
Docker chaos server stop containers [dockercs](https://github.com/barbosaigor/dockercs).  
(Under development) Kubenetes chaos server terminate pods [kubernetescs](https://github.com/barbosaigor/kubernetescs), in future it may terminate deployments and services.  
