# April

April proposes improve resilience in microservices architectures. It
does chaos testing by randomly shutting down nodes, taking into account 
their importance.  
April is a CLI tool, being possible to either run chaos testing or others tools,
such as 'server' which hosts an API for remote access to the chaos test.

## Tools
Chaos testing.  
```bash 
april -f conf.yml -n 10
```  

*Bare* runs only the selection algorithm, returning a set of nodes.  
```bash 
april bare -f conf.yml -n 10  
```  

*Server* will listen on port 8080. Default port is 7070.  
Need an 'destroy server' running to destroy instances.   
```bash 
april server -p 8080  
```  

