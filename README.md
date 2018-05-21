# HttpMon

Utilitário para monitorar disponibilidade de URLs http.

## Uso

``` shell
usage: httpmon --url=URL [<flags>]

Utilitário para monitorar disponibilidade de URLs http.
Flags:
  -h, --help         Show context-sensitive help (also try --help-long and --help-man).
  -t, --timeout=10s  Especifica o timeout da requisição.
  -v, --verbose      Imprime mais informações.
  -j, --json         Saida no formato json.
  -u, --url=URL      URL a monitorar.
      --version      Show application version.
```

## Sobre

Este utilitário avalia a disponibilidade de uma URL http e imprime se a mesma se encontra disponível ou não.

Além de imprimir de uma forma simplificada a disponibilidade da URL verificada, o utilitário ainda permite formatar a saída de forma mais verbosa ou em formato JSON. Estes dois últimos formatos são mais úteis para serem utilizados com ferramentas de agregação de log ou monitoramento por exemplo, pois exibe mais detalhes da tentativa de conexão.

Quando especificado um tipo de formato a saída imprime o código de retorno http e uma descrição do sucesso/erro. Este tipo de saída pode ser monitorado por ferramentas de APM ao longo do tempo de vida da aplicação e possibilita lançar alertas de acordo com a alteração de estado dos status do monitoramento.

Utilizar uma linguagem voltada para web como golang permite imputar muitos benefícios um utilitário como este, como por exemplo utilizar seu suporte a HTTP tracing (provido pelo pacote [net/http/httptrace](https://golang.org/pkg/net/http/httptrace/)) que pode ser utilizado para coletar informações durante um http request como por exemplo latência e DNS lookup ([https://blog.golang.org/http-tracing](https://blog.golang.org/http-tracing)).

# Roadmap

Incluir nas próximas versões

Suporte para:
1. monitorar conexões TCP
2. monitorar várias URLs
3. monitoramento contínuo das URLs
4. tracert de conexões http