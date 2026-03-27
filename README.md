# Daily Quest

Daily Quest é um projeto onde o Usuário irá criar várias tarefas e terá que fazer, tendo sistema de streak, xp, levels, criação de usuários, recompensas, algo bem gameficado

Esse projeto estou fazendo com o Intuito de aprender a mecher em algo real

## Arquitetura:

irei utilizar a arquitetura MVC

```
src/
|-- controler       - Entrada e Validação de dados
|-- model           - Regra de negócio e objeto principal
|-- view            - Gerenciamento de dados (público/privado) e converters
|-- test            - Teste de integração da aplicação
|-- configuation    - arquivos de configuração, etc...
main.go 
.env
.gitignore
```

### Fluxos:

**Como funciona o fluxo**
quando chega uma requisição HTTP:

```
HTTP Request -> Controller -> Service -> Domain (Model)

1. O controller recebe o JSON
2. Converte em struct (UserRequest)
3. Cria um domain object
4. Chama o service
5. O service aplica regra de negócio
6. (futuramente) salva no banco
```


## Packages:

- **godotenv**
- **gin-gonic**
- **Package validator**
- **zap logger**

