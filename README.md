# go-redis-example running Redis docker locally

### Description
We took a look at the basics of working with Redis in Go. 
We looked at how we can connect to Redis databases using the github.com/go-redis/redis package and we looked at how you can interact with that database using some of the methods that the package provides for us. (ping and pong as an example)

### Why Redis?
Redis is a fantastic open-source in-memory data structure store which can be used for various purposes such a database for your app, or a caching service or even a message broker.

It supports a wide variety of different data structures and is incredibly versatile and fast. If you are concerned with things like resiliency then you’ll be pleased to hear that it has built-in replication and can be run in a cluster setup to ensure that your applications are not reliant on a single instance.

For the purpose of this tutorial however, we are going to be keeping it nice and simple with a single, locally running instance of redis which we’ll be running with Docker.

