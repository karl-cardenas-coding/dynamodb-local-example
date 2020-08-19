# DynamoDB-local-example
This is an example project to showcase how to use DynamoDB locally while leveraging Terraform, localstack, and noSQL Workbench for DynamoDB. In-depth explanation provided in this the following Medium article, "How to use DynamoDB locally..".



## Customizing the upload script
Custommizing the upload script is pretty easy. The only file that that needs to be modified is the `/internal/upload.go`. Simply replace the content of the `Model` struct. Keep the same name please, otherwise the program will fail.

Example ( - = *Removal*, + = *Addition*   )
```go
type Model struct {
-	OrderID    string `json:"orderId"`
-	CustomerID string `json:"customerId"`
-	Shipped    string `json:"shipped,omitempty"`
-	Email      string `json:"email"`
-	Address    string `json:"address"`
-	Cost       string `json:"cost"`
}


// New example struct
type Model struct {
+	EmployeeId    string `json:"employeeId"`
+	Division string `json:"Division"`
+	HiringDate    string `json:"HiringDate"`
+	Email      string `json:"email"`
+	PayGrade    string `json:"payGrade"`
+   HRCode       string `json:"hrCode"`   
}

```

## Compiling Binary
Once you have the required change to `Model` strcut all that is required for you to compile the binary for the operating system you need it for. If you have Go installed in your environment, simply execute the `go build` command.

 I have included build commands for the most common operating systems. Should you need to compile for other opearting systems please visit the Golang documentation for the proper `GOOS` and `GOARCH` [values](https://golang.org/doc/install/source). 


### Windows 64 bit
```bash
 env GOOS=windows GOARCH=amd64 go build -o=upload-logic.exe
```

### Linux 64 bit
```bash
 env GOOS=linux GOARCH=amd64 go build -o=upload-logic
```

### Darwin (Mac) 64 bit
```bash
 env GOOS=darwin GOARCH=amd64 go build -o=upload-logic
```