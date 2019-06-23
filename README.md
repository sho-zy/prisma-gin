# prisma-gin
## Set up
- Prisma Client
  > `git clone https://github.com/sho-zy/prisma-gin.git`  
  > `dep ensure`  
  > `docker-compose up -d`  
  > `prisma deploy`  

- Prisma App
  > `dep ensure -update`  
  > `go run index.go`  

## URLs
- Prisma Admin  
  http://localhost:4466

- RestAPI Server  
  http://localhost:8080

## Commands
- Update and deploy the Prisma Client.  
  ( execute when you changed `./prisma/datamodel.prisma` )  
  > `prisma deploy`

  after, you can use the Prisma Admin.

- Start the RestAPI server. 
  > `go run index.go`  

  after, you can use the RestAPI server.

## Cases
- Case when you add tables.
  1. Add table definition to the following modules.  
   `./prisma/datamodel.prisma`  
  2. execute the following commands.  
        > `prisma deploy`  
  3. edit the resolver ( `./index.go` ).  
  4. execute the following commands.  
        > `dep ensure -update`  
        > `go run ./server`  
