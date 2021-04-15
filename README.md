# credisis-unidades-federativas
Atividade 2 

# Instalando
git clone https://github.com/natajonasdacoliveira/credisis-unidades-federativas

cd /credisis-unidades-federativas

# Troubleshooting
Caso necessário, instalar os pacotes necessários com o comando go get.

Exemplo : go get github.com/labstack/echo/v4

# Banco de dados
Acesse o arquivo db/db.go e modifique de acordo com o banco utilizado.

"root:root@tcp(127.0.0.1:3306)/UF"

Em ordem: nome, senha, IP, nome do banco.

# API
Podem ser encontradas no arquivo server.go, na função main.

Exemplo de rota não autenticada:localhost:1323/municipios (GET)

Exemplo de rota autenticada:localhost:1323/jwt/municipios (POST, PUT, DELETE)

As rotas autenticadas utilizam o bearer token que é recebido na /login (POST)

Usuários:

{
	"email":"teste@gmail.com",
	"password":"123123Teste"
}

{
	"email":"teste2@gmail.com",
	"password":"123123Teste"
}

Utilize o body para enviar esses dados
