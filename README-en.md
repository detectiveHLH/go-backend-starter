<h1 align="center">go-backend-starter</h1>

[中文](./README.md) | English

## About

Use Go Module as a dependency management, build Go-based web server based on Gin, use Endless to make the server restart smoothly, and use Swagger to automatically generate Api documents.

## Install

```bash
git clone https://github.com/detectiveHLH/go-backend-starter.git
cd go-backend-starter
go get
```

## Note

- go的版本需要高于1.11
- 使用goland时，需要确保Go module是Enable状态

## Usage

- View [Swagger API documentation](http://localhost:8080/swagger/index.html)
- View [Log in](http://localhost:8080/login?username=test&password=123) API, can get the access token for the system

## License

[MIT](./LICENSE)

