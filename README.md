# Go Server for our Blog System
[![Build Status](https://travis-ci.com/dsjlfjasdlkfjaklsf/go-server.svg?branch=main)](https://travis-ci.com/dsjlfjasdlkfjaklsf/go-server)
write, manage and share your blog 

## Overview

```
- App
  - controller   router and checker
  - model        model
  - service      database operator
  - util         util librart
  server.go      server
- config        
  - config.yaml  configurations of server and database
main.go          
```
## Running the server
To run the server, follow these simple steps:

```
go get ...
go run main.go
```

其中，数据库的Host、User、Password、Name，以及Server的Host、Port均按照``config.yaml``中的格式给出。