<center>
<img src="https://cdn.discordapp.com/attachments/976803430624526346/1230046001193156658/ZPLdown.gif?ex=6631e49b&is=661f6f9b&hm=831834b08e81486cba646958dc72616e630302534778f774b9203a704401c3f0&">
</center>

### Apresentação
No dia 15/04/2024, fiz a instalação de uma impressora térmica no comércio do meu pai. Percebi que o processo de gerar um código de barra era lento,cansativo e tedioso. Ao pesquisar, descobri o ZPL Viewer e sua API. Decidi então criar um pequeno programa no qual você configura as informações referente o modelo do código de barras que deseja criar, e ele faz o download automaticamente para você. Meu pai ficou muito feliz, pois agora, com apenas um comando, ele gera os seus tão amados códigos de barra.

### Recursos do ZPL-DOWN

- [x] [0.2] `PARAMETROS ACEITOS`
  - [x] DPMM
  - [x] WIDTH
  - [x] HEIGHT
  - [x] PREFIX
  - [ ] INDEX
  - [ ] ZPLCONFIG

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
|`GenerateAmount`  | Quantidade de códigos de barras que deseja gerar.                       |

### Exemplo:

```go
var barcode zpl.Barcode

barcode.Dpmm = "8dpmm"
barcode.Width = "1"
barcode.Height = "2"
barcode.Prefix = "XPTO"
barcode.GenerateAmount = 12

barcode.Download()
```
Certifique-se de inserir todas as informações exigidas. Após isso, realize a execução do programa.

### Qualquer dúvida referente ao uso do programa é só me chamar no Discord 👾
### "Se encontrar algum erro ou quiser me dar dicas de programação, me chame no Discord!"
