# Naive Solver
## Run

Requires python 3.8 or above :

```Shell
pip install sanic uuid6
python3 main.py
```

## Use

Send a formatted JSON knapset at ``/solve`` route, example :

```Shell
curl -X POST "localhost:8081/solve" -d @data.json
```
 
## Algorithm description

Just a naive brute force solver of Knap problem. 
The program will test every single combination in descending order.  

Example for 4 A-B-C-D items :

```
A + remaining (B C D)
    AB + (C D)
        ABC + (D)
            ABCD + (null)
        ABD + (null)
    AC + (D)
        ACD + (null)
    AD + (null)
        
B + remaining (B C D)
    BC + (D)
        BCD + (null)
    BD + (null)

C + remaining (B C D)
    CD + (null)

D + remaining (null)
```

## Multi processing implementation

// TODO n*4 threads split into n-1 threads for each first entry if remaining

## Benchmark
// TODO

Without multi threading : 

```shell
$ python3 -m benchmark.benchmark_bestPathFinder  
On average it took 3.0108033169999997 seconds  
...  
On average it took 2.987342144000004 seconds  
```
