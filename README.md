# OrderSystem_WeiZhang

## APP Default Listening Port is 3000
### It can be changed in 
> main.go   
```
http.ListenAndServe(":3000", context.ClearHandler(r))
```

## APP Default Using Googel API Key
### API Key can be changed in 
> config/config.go  
```
const GoogleMapApiKey = "API Key"
```
### App Support Google Client & signature   
> config/config.go  
```
const GoogleMapClientId = ""  
const GoogleMapClientSignature = ""  
```
