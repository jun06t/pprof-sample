# How to use
## Save profile
```
$ curl -s http://localhost:6060/debug/pprof/profile > out
```
It takes 30 seconds.

## Analyze
```
$ go tool pprof main out 
File: main
Type: cpu
Time: Apr 7, 2018 at 10:47am (JST)
Duration: 30.11s, Total samples = 26.25s (87.18%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 26.17s, 99.70% of 26.25s total
Dropped 16 nodes (cum <= 0.13s)
      flat  flat%   sum%        cum   cum%
    15.58s 59.35% 59.35%     15.58s 59.35%  main.fib.func1
     4.66s 17.75% 77.10%     24.31s 92.61%  main.main
     4.07s 15.50% 92.61%      4.07s 15.50%  main.main.func1
     1.86s  7.09% 99.70%      1.86s  7.09%  runtime.usleep
         0     0% 99.70%     24.31s 92.61%  runtime.main
         0     0% 99.70%      1.86s  7.09%  runtime.mstart
         0     0% 99.70%      1.86s  7.09%  runtime.mstart1
         0     0% 99.70%      1.86s  7.09%  runtime.sysmon
```

## Visualize
```
$ go tool pprof -png main out > out.png
```