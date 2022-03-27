# brankas api
 brankas forms api implementation 

### Upload Form API
- Below parameters read from the environment:
   - Port - The port to connect
   - Auth - The authentication

- Database part is not implemented due to the time limit.
- Code flow :
    - User do the GET request and get the html page.
    - Upload the file ( only images)
    - server recrives the fie
    - do size check
    - do type Check
    - save file to disk

- Execution:
   - go build -o brankasapi main.go
   - ./brankasapi

- Open the browder and type
   - http://localhost:PORT Number/form
