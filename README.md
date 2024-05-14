As the result of my lack of creativity, i've built this api such as the template of .net apps "weatherforecast". 

😄


# Running this app

Install Go 1.22 or latest

Open the terminal into the root folder of the app and run:
```
  go run .\main.go
```

This will start a web server in the 8080 port, be sure that this port is available or change it into the main.go file at "http.ListenAndServe(":8080", nil)".


# Using the app

Using CURL in Powershell:
```
  curl -Method Get -Uri http://localhost:8080/weatherforecast
```
