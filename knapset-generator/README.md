# Generator

## Webserver

> TL;DR
> ```GO
> go build ./...
> ```
> ```GO
> go run server.go
> ```

## CLI

> TL;DR
> ```GO
> go build ./...
> ```
> ```GO
> go run cli/cli.go
> ```

## How to use
> Item value is always a random int between `1` and `100`
> Item size is always a random int between `1` and `100`

### Webserver

Once the server is started, simply request the ``/generate`` route with ``bagSize`` & ``numberOfItems`` args, example :

```Shell
curl -X POST "localhost:8080/generate?bagSize=100&numberOfItems=12" -o data.json
```

### CLI
Requires Go installed (tested with ``Go 1.17``)

1. Build generator with `go build ./...`
2. Run generator with 0 to 4 [args](#args) (to customize the following behavior)
3. This will generate a new Knapsack problem (bag size + list of item with a value & size)
4. Print the result on the console
5. And on a JSON file

### Args
```GO
go run cli/cli.go A B C D 
```
- **A**: size of the bag => positive int or 0 for default value (current default value is random between `1` and `100`)
- **B**: number of items that will be generated => positive int or 0 for default value (current default value is `1` and `100`)
- **C**: print to console the problem generated => boolean (`true ; t ; 1 ; 0 ; f ; false ; ...`), default is `false`
- **D**: path to save of the JSON file where the generated problem will be written => string, default is `data.json`, target directory must already exist

### Examples
#### Default
```GO
go run generator
```
Will generate a problem :
- with a random bag size between `1` and `100`
- with a random number of items between `1` and `100`
- that will not be printed on the console
- but saved on a `data.json` file to the current directory

#### Default but printed
```GO
go run cli/cli.go 0 0 true
```
Will generate a problem :
- with a random bag size between `1` and `100`
- with a random number of items between `1` and `100`
- that **will be** printed on the console
- but saved on a `data.json` file to the current directory

#### Default but saved in a different file
```GO
go run cli/cli.go 0 0 true ../custom.json
```
Will generate a problem :
- with a random bag size between `1` and `100`
- with a random number of items between `1` and `100`
- that will **not** be printed on the console
- but saved on a `custome.json` file to the **`..` (parent)** directory

#### Custom Bag size of 150, with 23 different items, printed on the console and saved in a custom file
```GO
go run cli/cli.go 150 23 true 150-23.json
```
Will generate a problem :
- with a bag size `150`
- with `23` random items (value is always random between `1` and `100`)
- that **will be** printed on the console
- but saved on a `150-23.json` file to the current directory
