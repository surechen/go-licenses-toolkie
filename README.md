# go-licenses-toolkie
open source licenses toolkie for go.
### Function:
- 1. Collect licenses and notices file from gopath pkg directory.


### Example:
For project servicecomb-mesher(https://github.com/apache/servicecomb-mesher) 
```bash
cd /usr/local/src/go/src/github.com/surechen/servicecomb-mesher
go build
cd /usr/local/src/go/src/github.com/surechen/go-licenses-toolkie
go build
 ./go-licenses-toolkie /usr/local/src/go/src/github.com/surechen/servicecomb-mesher
 ```
