# todo-list-cli
A CLI to connect with my backend todo list app


# Design considerations:

### Directory structure: 

* Cmd directory --> executables
* Internal directory --> business logic
* Sub-directories inside cmd and interal are based on features/domain

# Packaging and deploying: 

* This project uses sqllite3 package which is a cgo package --> requires gcc compiler and also environment variable `CGO_ENABLED=1`


# Things learnt about go that were helpful:

* go doc `method name` gives docs for that method. Eg:
  > `go doc database/sql.DB.Query` or `go doc context.Background`

  NOTE: `go doc -src` shows source code for the same method
* To interact with sqlite db - `sqlite3 <db name>` and to quit cli --> `.quit`

  `.schema <table name>` prints schema of the table

* Go get vs go install --> go get calls go install after downloading library. Go install will compile the downloaded library. To just download the library and not install --> `go get -d <library name>` 

* `go env` gives all the golang env variables. GOROOT is where the go binary is stored. GOPATH is where the go libraries and packages downloaded using the get or install command are stored. Ensure the `$GOPATH/bin` variable is added to the \$PATH variable if I want to use any of the executables that are downloaded using golang. Eg: cobra-cli, or if I build my own application and store it there can call it from the cli. 
The value of GOPATH is usually `$HOME/go`. The downloaded source code is stored in `$GOPATH/pkg/mod` and the compiled binaries are stored in `$GOPATH/bin`.
NOTE: The `~` acts as a shortcut for $HOME in macos and unix systems. So `~/go/bin` will be `/Users/nibinmatteo/go/bin` in my computer. 
  
