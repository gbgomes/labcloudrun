# Execução e testes
## Ambiente DEV, sem uso do Docker e CloudRun
A partir da pasta raiz do projeto rodar o comando:

    go run main.go 

No navegador acessar a seguinte URL

    http://localhost:8080/clima?cep={CEP DESEJADO}

onde {CEP DESEJADO} deve ser substituído pelo CEP cujo clima será consultado

Exemplo

    http://localhost:8080/clima?cep=04112080

## Com uso do Docker
Fazer o build da imagem:

    docker build -t labcloudrun .

Executar com o comando:

    labCloudRun docker run -p 8080:8080 labcloudrun

No navegador acessar a seguinte URL

    http://localhost:8080/clima?cep={CEP DESEJADO}

onde {CEP DESEJADO} deve ser substituído pelo CEP cujo clima será consultado

Exemplo

    http://localhost:8080/clima?cep=04112080

## Com uso do CloudRun
Este projeto está disponível para uso no cloud run através da URL

    https://lab-cloud-run-rrg43hhm3q-uc.a.run.app/clima?cep={CEP DESEJADO}

onde {CEP DESEJADO} deve ser substituído pelo CEP cujo clima será consultado

Exemplo

    http://localhost:8080/clima?cep=04112080

