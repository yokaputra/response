# Go Rest API Response

This project is an implementation of flexible and easy-to-use HTTP response management in the Go language. With this project, you can easily generate JSON HTTP responses that correspond to the success or failure of your operations.

## Key Features

- **Success Response**: Create JSON responses that correspond to the success of your operations with a single function call.
- **Error Response**: Handle errors in your project and create responses that match the type of error.
- **Standard Error Codes**: There are standard error codes that can be used to identify the type of error.
- **Automatic HTTP Status Codes**: HTTP status codes will be automatically set based on the type of error that occurred.
- **Custom Messages**: You can provide custom messages for each generated response.
- **Additional Data**: Include additional data in your JSON responses.

## Getting Started

### Download

```shell
go get -u github.com/yokaputra/response
```

## Usage

Here's an example of how to use this project:

```go
// Generating a success response with default status code 200
successResponse := NewSuccessResponse("Operation successful", nil)

// Generating a success response with custom status code 201
successResponse := NewSuccessResponse("Operation successful", nil, http.StatusCreated)

// Generating an error response
err := errors.NewError(errors.ErrorCodeNotFound, "Data not found")
errorResponse := NewErrorResponse(err)
```

## Contribution

---

To contrib to this project, you can open a PR or an issue.

## License

This project is licensed under the Apache License 2.0. Please refer to the LICENSE file for more information.
