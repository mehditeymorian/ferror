<h1 align="center">
Ferror</h1>
<p align="center">GoFiber Error Response Handler</p>

<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/ferror?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>&nbsp;
<img src="https://img.shields.io/badge/license-mit-red?style=for-the-badge&logo=none" alt="license" />

<img src="https://img.shields.io/badge/Version-1.0.2-informational?style=for-the-badge&logo=none" alt="version" />
</p>

# Motivation
When you are developing software, it usually runs in different environments, staging and production being the most common environments. In the staging environment, you want to know as much as possible about the failure that happened. You want the status code, root cause, and other useful information. In contrast, the production environment is where you want to minimize the error output to information that is necessary. We found out that the status code, root cause, and message related to the situation are the essential information that you need to know when your request goes wrong.

**Ferror** is an error response handler, which outputs error responses with three different pieces of information described above, excluding root cause when in the production environment. It is important to exclude the root cause error in the production because it might include sensitive information such as table names, component names, etc.

## Usage

### Initialization
```go
errorHandler := ferror.NewErrorHandler(
    ferror.DevelopmentMode(true),
    ferror.OnErrorHandling(func(ctx *fiber.Ctx, err ferror.Error) {
        log.Infof("endpoint: %s msg: %s cause: %v\n", ctx.Path(), err.Message, err.Cause)
}))
```
Use in handlers like:
```go
func PeopleList(ctx *fiber.Ctx) error {
	peoples, err := db.PeopleList()
	if errors.Is(err, ErrNotFound) {
		return errorHandler.NotFound(ctx, err, "no people found")
	} else if err != nil {
		return errorHandler.InternalServerError(ctx, err, "unknown error happened")
	}

	return ctx.JSON(peoples)
}
```

## Responses

DevMode = true
```json
{
  "status": "response status code description",
  "message": "an abstract message of the error situation",
  "error": "root cause error"
}
```

DevMode = false
```json
{
  "status": "response status code description",
  "message": "an abstract message of the error situation",
}
```




# Installation
```bash
go install github.com/mehditeymorian/ferror
```

# Contribution
Any contribution in any form is welcomed. Open an issue to discuss it.

# Contact
- [Email](mailto:mehditeymorian322@gmail.com)