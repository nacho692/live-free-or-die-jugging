# Live Free Or Die Jugging: A Water Jug Riddle Solver

## Description
This application solves the [Die Hard Water Jug Problem](https://www.youtube.com/watch?v=2vdF6NASMiE)
but goes one step beyond and solves the more general problem:

Given an X gallons jug, Y gallons jug, and Z measurement goal (all integers),
measure either on either jug Z gallons of water.

## Running wjug

wjug is an interactive binary, it expects the user to provide X, Y and Z and 
generates a "friendly" output indicating the result of the water jug simulation

### Sample Run

```
Welcome to Water Jugs Riddle Solver!

The problem statement is:
There is have an X-gallon and a Y-gallon jug that you can fill from a lake. 
Measure Z gallons of water using only an X-gallon and Y-gallon jug.

Insert the value for the "x" jug, remember it must be positive: 5
Insert the value for the "y" jug, remember it must be positive: 4
Insert the value for the "z" goal, it must be smaller than either "x" or "y": 3
Fill Y 
(0/5, 4/4) 
Transfer to X 
(4/5, 0/4) 
Fill Y 
(4/5, 4/4) 
Transfer to X 
(5/5, 3/4) 
```

### Parameters
```
Usage of ./wjug:
  -s    silences most output so only the solution is printed
```

## Build

The built should be compatible with Mac, Linux and Windows architectures.
If you can't find the binary in [releases](https://github.com/nacho692/live-free-or-die-jugging/releases)
you can build your own.

### Dependencies
* [Go 1.20](https://go.dev/dl/)

```
go build ./cmd/wjug/...
```

## Dockerfile

If for any reason you'd rather run the program within a docker container, 
there are no images uploaded but you can build your own.

```
docker build -t wjug . 
docker run -it wjug
```
