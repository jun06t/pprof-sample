# How to use
## Save profile
```
$ curl -s http://localhost:6060/debug/pprof/heap > out
```

## Analyze
```
$ go tool pprof main out 
File: main
Type: inuse_space
Time: Apr 7, 2018 at 11:04am (JST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 126.34MB, 99.61% of 126.84MB total
Dropped 6 nodes (cum <= 0.63MB)
      flat  flat%   sum%        cum   cum%
  126.34MB 99.61% 99.61%   126.34MB 99.61%  main.leakyFunction
```

## Visualize
```
$ go tool pprof -png main out > out.png
```
![out](https://user-images.githubusercontent.com/4252009/38450162-3c294ece-3a54-11e8-9a38-ac581cc3ebe2.png)
