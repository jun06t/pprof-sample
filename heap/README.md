# How to use
## Save profile
```
$ curl http://localhost:6060/debug/pprof/heap > prof.out
```

## Analyze
```
$ go tool pprof main prof.out 
File: main
Type: inuse_space
Time: Apr 7, 2018 at 1:51am (JST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 21.19MB, 100% of 21.19MB total
      flat  flat%   sum%        cum   cum%
   21.19MB   100%   100%    21.19MB   100%  main.leakyFunction
```

## Visualize
```
$ go tool pprof -png main prof.out > out.png
```