
# Servidor HTTP Básico para compartilhar file server com Autenticação

Este é um simples servidor HTTP escrito em Go que serve arquivos de um diretório específico e utiliza autenticação básica.

## Pré-requisitos

- Go instalado (versão 1.16 ou superior)

## Instalação

Clone este repositório e navegue até o diretório do projeto:

```sh
git clone <URL_DO_REPOSITORIO>
cd <NOME_DO_DIRETORIO>
```

## Uso

Compile e execute o servidor com o seguinte comando:

```sh
go run main.go <diretorio> <porta>
```

- `<diretorio>`: O diretório que você deseja servir.
- `<porta>`: A porta na qual o servidor irá rodar.

### Exemplo

Para servir o diretório `public` na porta `8080`, execute:

```sh
go run main.go public 8080
```

## Autenticação

Este servidor usa autenticação básica para proteger o acesso aos arquivos servidos. As credenciais padrão são:

- Usuário: `admin`
- Senha: `teste`

### Como alterar as credenciais

Para alterar as credenciais, você precisará modificar a função `Secret` no código fonte (`main.go`). A função `Secret` é responsável por retornar o hash da senha para um dado usuário. O hash da senha pode ser gerado utilizando [esta ferramenta](https://unix4lyfe.org/crypt/).

### Função Secret

A função `Secret` verifica o usuário e retorna o hash da senha correspondente. No exemplo fornecido:

```go
func Secret(user, realm string) string {
    if user == "admin" {
        return "$1$/q9tTBog$KtdXn0eQTpiUP9g//xLaS1" //https://unix4lyfe.org/crypt/ MD5 Crypt: md5 salt
    }
    return ""
}
```

Se você deseja adicionar mais usuários ou alterar a senha, gere um novo hash e substitua o valor na função `Secret`.

## Estrutura do Código

- `main.go`: Arquivo principal que contém a implementação do servidor HTTP com autenticação básica.

## Dependências

Este projeto utiliza o pacote `go-http-auth` para gerenciamento de autenticação. Certifique-se de instalar as dependências com o comando:

```sh
go get -u github.com/abbot/go-http-auth
```

## Log

O servidor logará a porta em que está rodando e qualquer erro que ocorra durante a execução.

## Contributing
Contributions are welcome! If you have any improvements or new features you'd like to add to Laima, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or improvement.
3. Make the necessary changes and commit them.
4. Push your branch to your forked repository.
5. Submit a pull request to the main repository.

* Please ensure that your code follows the established coding conventions and includes appropriate tests for any new functionality.

## Licença
Laima is licensed under the MIT license. Please refer to the LICENSE file for more information.

## Doações
If you enjoyed using Laima, please consider making a donation to support the continuous development of the project. You can make a donation using one of the following options:
* Pix: rodrigo@hangell.org

* Cryptocurrencies or nft MetaMask: 0xEd4d1be72F807Faa358C966a8eF63367c200130F

![Created By](https://media.licdn.com/dms/image/D4D03AQF0vBM0rLZMKg/profile-displayphoto-shrink_200_200/0/1704050191664?e=1726099200&v=beta&t=JiPipqyppQaj1f6tR6tI2cMojmCAgJFQXkJgZdAZKqk)



<div>
  <a href="https://hangell.org" target="_blank"><img src="https://img.shields.io/badge/website-000000?style=for-the-badge&logo=About.me&logoColor=white" target="_blank"></a>
  <a href="https://play.google.com/store/apps/dev?id=5606456325281613718" target="_blank"><img src="https://img.shields.io/badge/Google_Play-414141?style=for-the-badge&logo=google-play&logoColor=white" target="_blank"></a>
  <a href="https://www.youtube.com/channel/UC8_zG7RFM2aMhI-p-6zmixw" target="_blank"><img src="https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white" target="_blank"></a>
  <a href="https://www.facebook.com/hangell.org" target="_blank"><img src="	https://img.shields.io/badge/Facebook-1877F2?style=for-the-badge&logo=facebook&logoColor=white" target="_blank"></a>
  <a href="https://www.linkedin.com/in/rodrigo-rangel-a80810170" target="_blank"><img src="https://img.shields.io/badge/-LinkedIn-%230077B5?style=for-the-badge&logo=linkedin&logoColor=white" target="_blank"></a>

</div>