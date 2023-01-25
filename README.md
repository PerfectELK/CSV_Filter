### Program which delete needless columns in csv file

#### Usage examples:
```
go run main.go  -file="test.csv" -include="Test,Test 1" -comma=";"
go run main.go  -file="test.csv" -exclude="Test,Test 1" -comma=";"
```
#### Command line arguments:
```
file - file path
comma - cell separator
include - cell, which will be save
exclude - cell, which will be delete (ignore include of use this arg)
```
