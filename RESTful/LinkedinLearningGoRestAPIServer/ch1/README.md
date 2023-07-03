To start a module or a application you can use the following command:

`go mod init [packageName/moduleName]`

ex
`go mod init gabuh.com/gabuh/module1`

Can specify the name and address where the module or application can be downloaded.

In the first example we created two modules in the src directory, `backend` and `practiceit` where each one has its own go file with functions and methods to be used. 

Since the `backend` module must be used by the `practiceit` module, in the `practiceit.go` file we make an import as following: 

```go
import (
       "github.com/gabuh/backend"
	)
```

and then we can use the code from the backend module.

ex: 

```go
   backend.Run(":8080")
```

Use the following commands in the `practiceit` module to use the `backend` offline:
`go mod edit -replace github.com/gabuh/backend=[path_Of_The_Module]`
ex:
`go mod edit -replace github.com/gabuh/backend=../backend`

and then hit the following one:

`go mod tidy`

then you ready to run.

OBS: All the mod interactions can be monitored looking at the file `go.mod` 

