x-scrap
======

scrapping kreuzwerker blog posts for fun and finding missing untagged keywords in blog posts

usage
------

    ./xscrap -urls <url,url2> -tags <tag1,tag2, ... >

with 
 * _urls_ a comma seperated list of urls to scrap
 * _tags_ a comma seperated list of tags to find

for example

```
go run cmd/xscrap/main.go \
    -urls https://kreuzwerker.de/post/aws-re-inforce-2022,https://kreuzwerker.de/post/kreuzwerker-goes-aws-ambassador-in-seattle \
    -tags aws,security,cloud
``` 

Build
------

To build the binary invoke 
```
go build -o xscrap cmd/xscrap/main.go 
```

Tests
------

To run the tests invoke
```
go test ./... -cover
```
