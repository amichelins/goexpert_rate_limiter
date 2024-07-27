# goexpert_rate_limiter

O Projeto se baseia no PKG github.com/amichelins/amsrtl

O PKG amsrtl pode ser inicializado via variaveis de ambiente com a rotina
amsrtl.NewEnvLimiter({storage}), onde {storage} é um PKG que implmente a interface
Storage.
Neste caso utilizamos o PKG github.com/amichelins/amsrtl/storage/redis
que utiliza o REDIS como storage

As variaveis de ambiente utilizados são:
LIMITER_MAX -> Limite maxímo de conexões por IP
LIMITER_BLOCK_DURATION -> Duração em segundos do bloqueio, se não informado,  300 segundos
LIMITER_TOKENS -> Lista de tokens e limites enviados via header API_KEY.
                  É um JSON com este formato:
                [{"token":"{TOKEN}","limit": {LIMITE}},{"token":"{TOKEN}","limit": {LIMITE}}]

                EX: [{"token":"{TOKENA}","limit": 10},{"token":"TOKENB","limit": 10}]

Build:
       docker  compose up --build -d -V

A porta de conexão http é a 8080

Teste:
    Foi adicionado o main_test.go contendo:

    1 - Ter rodado o docker  compose up --build -d -V para iniciar os serviços

    2 - Rodar: go test -timeout 300s github.com/amichelins/goexpert_rate_limiter