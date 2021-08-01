# Password Service

Go-based web microservice for generating passwords.

## Usage

Make sure to have Go installed on your system. Assuming you're running Linux or macOS (haven't tested it on Windows, nor have a reason to), compile the program producing an executable and run it

```bash
go build *.go
./main
```

The webservice now runs on port `3333`. It exposes a single GET method at `/generate` with the following available query parameters:

- `length` password length as an integer, if it's not provided it'll assume the default value of 8 characters
- `upper` boolean value to set whether you wish the password to include uppercase letters
- `numbers` boolean value to set whether you wish the password to include numbers (digits)
- `specials` boolean value to set whether you wish the password to include special characters

Alternatively, you can run the application inside a Docker container

```bash
docker build --tag <tag> --file Dockerfile .
docker run --name <name> --publish 3333:3333 <tag>
```

## License

Project is open-sourced under MIT license.