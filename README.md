# go-lang-tutorial

## Troubleshootings

### Plugin do Vscode

A extension do vscode apenas funciona corretamente caso o workspace esteja aberto na pasta a qual se encontra o arquivo go.mod.

Caso o workspace esteja aberto em outra pasta, haverá problemas de auto complete e de import de dependências.

Uma forma de resolver isso é abrir um dos modulos na pasta no qual está o arquivo go.mod e adicionar as outras dependências ao projeto. Para isso, basta clicar com o botão direito no menu de navegação do vscode, clicar em "Add folder to workspace..." e adicionar as outras pastas do projeto.
