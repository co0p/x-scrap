x-scrap
======

scrapping kreuzwerker blog posts for fun and finding missing untagged keywords in blog posts

usage
------

    ./xscrap -url <url> -tags <tag1,tag2, ... >

with 
 * _url_ being the url to scrap
 * _tags_ a comma seperated list of tags to find

for example

```
./xscrap -url https://kreuzwerker.de/post/aws-re-inforce-2022 -tags aws,security,cloud
``` 

Build
------

To build the binary invoke 
```
go build -o xscrap cmd/xscap/main.go 
```

Tests
------

To run the tests invoke
```
go test ./...
```
