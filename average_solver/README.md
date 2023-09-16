# Average Solver
# Run Solver
## Run with Docker

Build the image with : ``docker build -t average-solver .``  
Run the image with : ``docker run --name average-solver -p 3000:3000 -d average-solver``

## Run locally

Requires Julia 1.7 or above :

```Shell
./bin/server
```

# Use Server

Send a formatted JSON 'knapset' (see an example below) at ``/average-solver/solve`` route, example :

```Shell
curl -X POST -H "Content-Type: application/json" "localhost:3000/average-solver/solve" -d @data.json
```

## Knapset template example
```Json
{
  "bagSize": 50,
  "items":
  [{
      "id": 0,
      "size": 64,
      "value": 78
    },
    {
      "id": 1,
      "size": 35,
      "value": 87
    },
    {
      "id": 2,
      "size": 76,
      "value": 53
    }]
}
```

## Response template example
```Json
{
  "max_value": 87,
  "items": [
    {
      "id": 1,
      "size": 34,
      "value": 87
    }
  ],
  "duration": 0.006
}
```

# Algorithm description

The algorithm get the average value per size for each items and gets the highest average items that fits into the total bag size. 

# Multi processing implementation

Not needed since it's not a very complex algorithm and speed is not the bottleneck. 

# Benchmark
// TODO

```
