# Execução e testes

## Ambiente sem uso do Docker e CloudRun
A partir da pasta raiz do projeto rodar o comando:

    go run main.go 

No navegador acessar a seguinte URL

    http://localhost:8080/clima?cep={CEP DESEJADO}

onde {CEP DESEJADO} deve ser substituído pelo CEP cujo clima será consultado

Exemplo

    http://localhost:8080/clima?cep=04112080


## Ambiente DEV com uso de docker compose
Fazer o build da imagem:

    docker compose build --no-cache

Para entrar na imagem para desenvolvimento, executar o comando abaixo:

    docker compose run --rm labcloudrun


## Ambiente de produção com uso do Docker
Fazer o build da imagem:

    docker compose -f docker-compose.prod.yaml build --no-cache

Para subir o ambiente, executar o comando abaixo:

    docker compose -f docker-compose.prod.yaml up -d

No navegador acessar a seguinte URL

    http://localhost:8080/clima?cep={CEP DESEJADO}

onde {CEP DESEJADO} deve ser substituído pelo CEP cujo clima será consultado

Exemplo

    http://localhost:8080/clima?cep=04112080

Para parar o ambiente, execute o comando abaixo:

    docker compose -f docker-compose.prod.yaml down

## Com uso do CloudRun
Este projeto está disponível para uso no cloud run através da URL

    https://lab-cloud-run-rrg43hhm3q-uc.a.run.app/clima?cep={CEP DESEJADO}

onde {CEP DESEJADO} deve ser substituído pelo CEP cujo clima será consultado

Exemplo

    https://lab-cloud-run-rrg43hhm3q-uc.a.run.app/clima?cep=04112-080

