[![Go version](https://img.shields.io/github/go-mod/go-version/karl-cardenas-coding/dynamodb-local-example?filename=upload-tool%2Fgo.mod)](https://golang.org/dl/)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/karl-cardenas-coding/dynamodb-local-example.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/karl-cardenas-coding/dynamodb-local-example/alerts/)

# DynamoDB-local-example
This is an example project to showcase how to use DynamoDB locally while leveraging Terraform, localstack, and noSQL Workbench for DynamoDB. In-depth explanation provided in this the following Medium article, "How to use DynamoDB locally..".



## Changing JSON Content
You can either replace the JSON content inside the file `static/raw-data.json` OR provide your own file. Just make sure you change the variable value `json-file-path` in the `terraform.tfvars` file. If you want to leverage the AWS CLI approach then follow the same steps, either replace `static/formatted-data.json` OR change the content inside. Just ensure that the content is in the correct format that DynamoDB expects (DynamoDB JSON).



## Customizing the upload script
Customizing the upload script is pretty easy. The only file that that needs to be modified is the `upload-tool/internal/upload.go`. Simply replace the content of the `Model` struct. Keep the same name `Model` please, otherwise the program will fail.

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
+	Division      string `json:"Division"`
+	HiringDate    string `json:"HiringDate"`
+	Email         string `json:"email"`
+	PayGrade      string `json:"payGrade"`
+   HRCode        string `json:"hrCode"`   
}

```

## Compiling Binary
Once you have the required change to the `Model` struct all that is required of you is to compile the binary for the operating system you need it for. If you have Go installed on your environment, simply execute the `go build` command.

 I have included build commands for the most common operating systems. Should you need to compile for other opearting systems please visit the Golang documentation for the proper `GOOS` and `GOARCH` [values](https://golang.org/doc/install/source). Place the compiled binary in the root of the project structure.


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