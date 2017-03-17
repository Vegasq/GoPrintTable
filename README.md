# GoPrintTable

## Install

For stable release:

`go get gopkg.in/vegasq/GoPrintTable.v1`

`import "gopkg.in/vegasq/GoPrintTable.v1"`

For devel release:

`go get github.com/vegasq/GoPrintTable`

`import "github.com/vegasq/GoPrintTable"`



## Usage
Let's assume we have such structure:
```go
header := []string{"Node Name", "IP", "Status"}
node1 := []string{"controller", "192.168.0.100", "Online"}
node2 := []string{"minion1", "192.168.0.101", "Offline"}
node3 := []string{"minion2", "192.168.0.102"}
t := [][]string{header, node1, node2, node3}
```

Now let's print it to console with:

```go
GoPrintTable.PrintTableWithHeader(t)
```

Output will be:

```
----------------------------------------
| Node Name  | IP            | Status  |
----------------------------------------
| controller | 192.168.0.100 | Online  |
| minion1    | 192.168.0.101 | Offline |
| minion2    | 192.168.0.102 | -       |
----------------------------------------
```

Or with:

```go
GoPrintTable.PrintTable(t)
```

Output will be:
```
----------------------------------------
| Node Name  | IP            | Status  |
| controller | 192.168.0.100 | Online  |
| minion1    | 192.168.0.101 | Offline |
| minion2    | 192.168.0.102 | -       |
----------------------------------------
```

