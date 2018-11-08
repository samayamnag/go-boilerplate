# I Change My City


### How To Run This Project
```bash
#move to directory
cd $GOPATH/src/github.com/<username>

# Clone into YOUR $GOPATH/src
git clone https://github.com/samayamnag/go-icmyc-complaints.git

#move to project
cd go-icmyc-complaints

# Install Dependencies
dep ensure

# Run Project
go run main.go

```

Or

```bash
# GET WITH GO GET
go get github.com/samayamnag/go-icmyc-complaints

# Go to directory

cd $GOPATH/src/github.com/go-icmyc-complaints

# Install Dependencies
dep ensure


# Run Project
go run main.go
```

### Avialable routes
```bash
1. localhost:<port> - GET
2. localhost:<port>users - GET - All Users
3. localhost:<port>/users/ - POST - Create User
4. localhost:<port>/users/{id} - GET - Get User
5. localhost:<port>/v1/users - GET - All Users
6. localhost:<port>/v1/users/ - POST - Create User
7. localhost:<port>/v1/users/{id} - GET - Get User
```