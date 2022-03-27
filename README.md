# brankasapi
 brankas forms api implementation 

Upload Form API
Two parameters read from the environment:
1) Port - The port to connect
2) Auth - The authentication

Database part is not implemented due to the time limit.
Code flow:
    1)User do the GET request and get the form html page.
    Upload the file ( only images)
    server recrievs the fie
    do size check
    do type Check
    save file to disk

Execution:
go build -o brankasapi main.go
./brankasapi
