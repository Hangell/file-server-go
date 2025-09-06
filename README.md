
# Basic HTTP Server to Share File Server with Authentication

This is a simple HTTP server written in Go that serves files from a specific directory and uses basic authentication.

## Prerequisites

- Go installed (version 1.16 or higher)

## Installation

Clone this repository and navigate to the project directory:

```sh
git clone <REPOSITORY_URL>
cd <DIRECTORY_NAME>
```

## Code Structure

```cmd
/shared-files-server
│   auth.go
│   go.mod
│   go.sum
│   main.go
│   README.md
│   server.go
```
- `auth.go`: Contains the authentication logic. The `Secret` function validates user credentials, and `NewAuthenticator` creates a new BasicAuthenticator.
- `main.go`: The entry point of the program. It processes command-line arguments, creates the authenticator, and starts the server.
- `server.go`: Contains the server logic. The `StartServer` function configures and starts the HTTP server.


## Usage

Compile and run the server with the following command:

```sh
go run main.go <directory> <port>
```

- `<directory>`: The directory you want to serve.
- `<port>`: The port on which the server will run.

### Example

To serve the `public` directory on port `9000`, run:

```sh
go run main.go public 9000
```

## Authentication

This server uses basic authentication to protect access to the served files. The default credentials are:

- User: `admin` | `dev` | `user`
- Password: `teste`

### How to Change the Credentials

To change the credentials, you will need to modify the `Secret` function in the source code (`auth.go`). The `Secret` function is responsible for returning the password hash for a given user. The password hash can be generated using [this tool](https://unix4lyfe.org/crypt/).

### Secret Function

The `Secret` function checks the user and returns the corresponding password hash. In the provided example:

```go
func Secret(user, realm string) string {
    if user == "admin" || user == "dev" || user == "user" {
        return "$1$/q9tTBog$KtdXn0eQTpiUP9g//xLaS1" //https://unix4lyfe.org/crypt/ MD5 Crypt: md5 salt
    }
    return ""
}
```

If you want to add more users or change the password, generate a new hash and replace the value in the `Secret` function.

## Dependencies

This project uses the `go-http-auth` package for authentication management. Make sure to install the dependencies with the command:

```sh
go get -u github.com/abbot/go-http-auth
```

## Log

The server will log the port it is running on and any errors that occur during execution.

## Contributing
Contributions are welcome! If you have any improvements or new features you'd like to add to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or improvement.
3. Make the necessary changes and commit them.
4. Push your branch to your forked repository.
5. Submit a pull request to the main repository.

* Please ensure that your code follows the established coding conventions and includes appropriate tests for any new functionality.

## License
*File Server* is licensed under the MIT license. Please refer to the LICENSE file for more information.

## Donations
If you enjoyed using this project, please consider making a donation to support the continuous development of the project. You can make a donation using one of the following options:
* Pix: rodrigo@hangell.org
* Cryptocurrencies or NFT MetaMask: 0xEd4d1be72F807Faa358C966a8eF63367c200130F

## 👨‍💻 Author

**Rodrigo Rangel**

<div align="left">
  <a href="https://hangell.org" target="_blank">
    <img src="https://img.shields.io/badge/website-000000?style=for-the-badge&logo=About.me&logoColor=white" alt="Website" />
  </a>
  <a href="https://play.google.com/store/apps/dev?id=5606456325281613718" target="_blank">
    <img src="https://img.shields.io/badge/Google_Play-414141?style=for-the-badge&logo=google-play&logoColor=white" alt="Google Play" />
  </a>
  <a href="https://www.youtube.com/channel/UC8_zG7RFM2aMhI-p-6zmixw" target="_blank">
    <img src="https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white" alt="YouTube" />
  </a>
  <a href="https://www.facebook.com/hangell.org" target="_blank">
    <img src="https://img.shields.io/badge/Facebook-1877F2?style=for-the-badge&logo=facebook&logoColor=white" alt="Facebook" />
  </a>
  <a href="https://www.linkedin.com/in/rodrigo-rangel-a80810170" target="_blank">
    <img src="https://img.shields.io/badge/-LinkedIn-%230077B5?style=for-the-badge&logo=linkedin&logoColor=white" alt="LinkedIn" />
  </a>
</div>

---

<p align="center">
  Made with ❤️ for the developer community
</p>
