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
Time: Apr 7, 2018 at 10:56am (JST)
Duration: 30.16s, Total samples = 25.97s (86.11%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 25.95s, 99.92% of 25.97s total
Dropped 11 nodes (cum <= 0.13s)
Showing top 10 nodes out of 20
      flat  flat%   sum%        cum   cum%
    14.93s 57.49% 57.49%     14.93s 57.49%  main.fib.func1
     5.33s 20.52% 78.01%     23.90s 92.03%  main.main
     3.64s 14.02% 92.03%      3.64s 14.02%  main.main.func1
     1.91s  7.35% 99.38%      1.91s  7.35%  runtime.usleep
     0.14s  0.54% 99.92%      0.14s  0.54%  runtime.mach_semaphore_signal
         0     0% 99.92%      0.14s  0.54%  runtime.goready
         0     0% 99.92%      0.14s  0.54%  runtime.goready.func1
         0     0% 99.92%      0.14s  0.54%  runtime.goroutineReady
         0     0% 99.92%      0.14s  0.54%  runtime.mach_semrelease
         0     0% 99.92%     23.90s 92.03%  runtime.main
```

## Visualize
```
$ go tool pprof -png main out > out.png
```

![out](https://user-images.githubusercontent.com/4252009/38450088-17f3db9c-3a53-11e8-88eb-76e73755da28.png)
