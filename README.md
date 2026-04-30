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
* 
  
