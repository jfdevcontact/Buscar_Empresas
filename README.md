# Buscar_Empresas
 Um software que faz consulta por meio de uma api para saber informaçoẽs de determinada empresa

 Função que vou fazer FUTURAMENTE

salva essas infomações em un banco de dados MYSQL ou em um arquivo TXT

SITE AONDE EU PEGUEI ESSA API (NÃO PRECISA DE KEY).(ELA É FREE)

SITE: https://developers.receitaws.com.br/#/operations/queryCNPJFree


Coloquei uma cor VERDE no terminal, quando os dados da empresa são retornado

SITE: A onde eu peguei os esquemas de cores: https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go


Informações técnicas

Essa api ela retorno um json. Estou usndo Unmarshal do golang para tratar esses dados

Esta API permite 3 consultas por minuto

Não estou tratando erros, por exemplo. Não estou tratando erro caso o CNPJ não existe

Não estou tratando o erro caso o limite de requests seja "estourado em 3 minutos"


Meu canalno YT: https://www.youtube.com/@jf.dev.contact