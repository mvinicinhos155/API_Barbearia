# 💈API_BARBEARIA

## 📌Descrição
Essa API foi desenvolvida em Go para o gerenciamento de uma barbearia, permitindo o controle de usuários, agendamentos e serviços. Este foi o meu primeiro projeto "grande" que eu desenvolvi utilizando Go. A aplicação utiliza Docker para o gerenciamento do banco de dados PostgreSQL.

## 🚀 Tecnologias
Go (Golang)<br>
PostgreSQL<br>
Docker

## 📦 Como rodar o projeto

Você vai precisar ter instalado:<br>
Go (versão 1.25.6 ou superior)<br>
Docker<br>
Git

Configurar variáveis de ambiente
Crie um arquivo .env na raiz do projeto com base no .env.example:

DB_HOST=<br>
DB_PORT=<br>
DB_USER=<br>
DB_PASSWORD=<br>
DB_NAME=<br>
JWT_SECRET=sua_chave_secreta

🐳 Subir o banco com Docker<br>
docker-compose up -d

▶️ Rodar a aplicação<br>
go run ./cmd

A API estará disponível em:

http://localhost:8080