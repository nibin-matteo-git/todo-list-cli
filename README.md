# todo-list-cli
A CLI to connect with my backend todo list app


# Design considerations:

### Directory structure: 

* Cmd directory --> executables
* Internal directory --> business logic
* Sub-directories inside cmd and interal are based on features/domain

# Packaging and deploying

* This project uses sqllite3 package which is a cgo package --> requires gcc compiler and also environment variable `CGO_ENABLED=1`
* 
