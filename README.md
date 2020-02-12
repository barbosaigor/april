# April

April proposes improve resilience in microservices architectures. It
does chaos testing by randomly shutting down nodes, taking into account 
their importance.  
April is a CLI tool, being possible to either run chaos testing or others tools,
such as 'server' which hosts an API for remote access to the chaos test.

## Tools
Chaos test.  
-f configuraion file path  
-n number of nodes to choose  
```bash 
april -f conf.yml -n 10
```  

*Bare* runs only the selection algorithm, returning a set of nodes.  
-f configuraion file path  
-n number of nodes to choose  
```bash 
april bare -f conf.yml -n 10  
```  

*Server* hosts an API (HTTP) which apply chaos testing and bare algorithm.  
Need an 'chaos server' running to destroy instances.  
-p port number (Default is 7070)  
```bash 
april server -p 8080  // will listen on port 8080
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
A [Docker](https://github.com/barbosaigor/aprilcsdocker) implementation  

