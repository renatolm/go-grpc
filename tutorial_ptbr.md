Tutorial projeto gRPC com Go:
1. cria uma pasta para o projeto
2. executa o comando `go mod init <url>` (essa url preferencialmente é uma url do github onde o projeto pode ser baixado)
3. cria a pasta internal
4. a pasta database dentro de internal foi copiada do projeto de graphql
5. executa o comando `go mod tidy` para baixar as dependencias

Criando o protofile:
1. Cria uma pasta `proto` na raiz
2. Cria um arquivo `car_manufacturer.proto` (o nome é indiferente)
3. Colocar `syntax = "proto3";` nesse arquivo
4. Precisamos escolher um package para esse arquivo, é muito comum utilizar `package pb` (de protocol buffers)
5. Indicar onde o pacote vai ser instalado com `option go_package = "internal/pb";`
6. Cria uma message para o Manufacturer com os campos do struct (os números são apenas a ordem)
7. Cria uma message ManufacturerResponse para descrever a resposta de um pedido de Manufacturer
8. Cria uma message de request de criação de Manufacturer CreateManufacturerRequest, aqui não vai o id porque ainda não sabemos antes de criar
9. Cria um serviço ManufacturerService

Geração de código com protoc:
1. Rodar o comando `protoc --go_out=. --go-grpc_out=. proto/car_manufacturer.proto`
2. Se olhar dentro da pasta internal, ele vai ter criado uma pasta pb com os arquivos
3. Depois executar o comando `go mod tidy` para baixar as dependencias novas

Implementando CreateManufacturer:
1. Dentro de internal criar uma pasta `service` e nela criar o arquivo `manufacturer.go`
2. Nesse arquivo, colocar o package como `service` e criar uma struct ManufacturerService
3. Nesse struct precisamos colocar o `pb.UnimplementedManufacturerServiceServer` que está no arquivo `internal>pb>car_manufacturer_grpc.pb.go`
4. Também vamos adicionar o DB do Manufacturer no struct
5. E precisamos fazer os imports dos packages `pb` e `database`
6. Também precisamos implementar um método de criação `NewManufacturerService`
7. Agora precisamos implementar um método para criar um Manufacturer, basicamente podemos copiar o esqueleto do método de `car_manufacturer_grpc.pb.go`

Criando o servidor gRPC:
1. Na raiz do projeto, criar a pasta `cmd`, dentro dela `grpcServer` e dentro o arquivo `main.go`
2. No arquivo main, primeiramente fazemos a conexão com o banco de dados
3. Depois disso criamos o manufacturerDb e o serviço de Manufacturer
4. E aí criamos o server gRPC e registramos o serviço nele
5. Agora precisamos criar uma conexão TCP, para isso usamos um listener

Interagindo com o servidor gRPC:
1. Para interagir com o servidor, vamos usar um client gRPC chamado evans (https://github.com/ktr0731/evans)
2. Primeiro instalamos o client evans
3. Depois colocamos o server para rodar `go run cmd/grpcServer/main.go` e iniciamos o evans com o comando `evans -r repl`
4. Se não tiver com o package selecionado, podemos usar `show package` para listar e `package pb` para usar o nosso package
5. De maneira similar, `show service` para listar os services e `service ManufacturerService` para usar o nosso service
6. Então usar o comando `call CreateManufacturer` e preencher os campos
7. Provavelmente vai dar erro porque a tabela do sqlite não foi criada previamente
8. Podemos criar ela na mão em outro terminal com `sqlite3 db.sqlite` e `create table manufacturer (id string, name string, origin_country string);`
9. Se tentar de novo no evans vai dar certo