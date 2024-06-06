<center>
<img src="https://imgur.com/tkdEL4z.png">
</center>

### Apresentação

![Obsoleto](https://img.shields.io/badge/AVISO-Documentação%20obsoleta-red)

A documentação atual está obsoleta, irei fazer sua atualização no dia 06/06/24.

No dia 15/04/2024, fiz a instalação de uma impressora térmica no comércio do meu pai. Percebi que o processo de gerar um código de barra era lento,cansativo e tedioso. Ao pesquisar, descobri o ZPL Viewer e sua API. Decidi então criar um pequeno programa no qual você configura as informações referente o modelo do código de barras que deseja criar, e ele faz o download automaticamente para você. Meu pai ficou muito feliz, pois agora, com apenas um comando, ele gera os seus tão amados códigos de barra.

### Recursos do ZPL-DOWN

- [x] [0.3.1] `PARAMETROS ACEITOS`
  - [x] DPMM
  - [x] WIDTH
  - [x] HEIGHT
  - [x] PREFIX
  - [x] ZPL-CONFIG
  - [x] INDEX
  - [x] AMOUNT
  - [x] PADWIDTH

### 🖥️ Pré-requisitos
- Certifique-se de ter o GO instalado em seu computador.
- Certifique-se de possuir o Git instalado em seu sistema para clonar este repositório.

### 🖥️ Instalação e Execução

```bash
$ git clone https://github.com/guilherme-dell/zpl-down.git
$ cd zpl-down/cmd
$ go run main.go
```
<p style="color:red;">Antes de executar o programa, configure o código de barras que deseja gerar.</p>

### Configurando as informações do código de barra:


|                  | Descrição                                                               |
|----------------  |-------------------------------------------------------------------------|
|`Dpmm`            | Densidade de impressão desejada, em pontos por milímetro.               |
|`Width`           | Largura da etiqueta, em polegadas.                                      |
|`Height`          | Altura da etiqueta, em polegadas.                                       |
|`Prefix`          | Prefixo que deseja usar no código de barras.                            |
|`Zpl-Config`      | Configuração do código ZPL referente o código de barras.                |
|`Index`           | Número do index inicial utilizado para gerar o código de barras.        |
|`Amount`          | Quantidade de códigos de barras que deseja gerar.                       |
|`PadWidth`        | Quantidade de digitos zeros à esquerda.                                 |


### Exemplo:

```json 
{
    "BarcodeDir": "./barcodes",
    "BarcodeConfig": {
        "Dpmm": "8dpmm",
        "Width": "0.98",
        "Height": "0.39",
        "ZplConfig": "^xa^fo4,15^be1^bcn,40,,,,a^fd",
        "Prefix": "CJ",
        "Index": 1,
        "Amount": 2,
        "PadWidth": 4
    }
}

```
Certifique-se de inserir todas as informações exigidas. Após isso, realize a execução do programa.

### Qualquer dúvida referente ao uso do programa é só me chamar no Discord 👾
### "Se encontrar algum erro ou quiser me dar dicas de programação, me chame no Discord!"
