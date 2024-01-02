<!DOCTYPE html>
<html lang="pt-BR">

<head>
    <meta charset="UTF-8">
    <title>README - Projeto API Restful Go e React</title>
</head>

<body>
    <header>
        <h1>Projeto de Desenvolvimento de API Restful com Go e React</h1>
        <p>Este projeto consiste na construção de uma API Restful do zero utilizando a linguagem Go para o backend e
            React para o frontend, proporcionando uma estrutura completa de CRUD (Create, Read, Update, Delete) para um
            recurso específico, conectando-se a um banco de dados PostgreSQL via Docker. A integração entre backend e
            frontend foi realizada com a configuração do CORS.</p>
    </header>

<main>
    <section>
        <h2>Tecnologias Utilizadas</h2>
        <ul>
            <li><strong>Backend:</strong>
                <ul>
                    <li>Go</li>
                    <li>Roteador para criação de endpoints</li>
                    <li>GORM (Go Object Relational Mapper) para a interação com o banco de dados PostgreSQL</li>
                    <li>Docker para a execução do ambiente de banco de dados</li>
                </ul>
            </li>
            <li><strong>Frontend:</strong>
                <ul>
                    <li>React para a construção da interface do usuário</li>
                    <li>CORS configurado para a integração com o backend</li>
                <li><a href="https://github.com/jonathantx/frontend-personalities" target="_BLANK">Repositório Front-end</a></li>
                </ul>
            </li>
        </ul>
    </section>
    <section>
        <h2>Funcionalidades Implementadas</h2>
        <ul>
            <li>Criação de rotas e endpoints para a manipulação de recursos</li>
            <li>Utilização de JSON para formatação das requisições e respostas</li>
            <li>Conexão com banco de dados PostgreSQL via Docker para persistência de dados</li>
            <li>Implementação do CRUD completo para um recurso específico</li>
            <li>Utilização do GORM para facilitar a criação, deleção e edição de dados no banco de dados</li>
            <li>Uso de middlewares para controle de requisições</li>
            <li>Integração bem-sucedida entre o backend em Go e o frontend em React</li>
        </ul>
    </section>
    <section>
        <h2>Como Executar o Projeto</h2>
        <ol>
            <li><strong>Backend:</strong>
                <ol>
                    <li>Certifique-se de ter o Docker instalado em sua máquina.</li>
                    <li>Clone este repositório.</li>
                    <li>Navegue até o diretório do backend.</li>
                    <li>Execute o comando <code>docker-compose up</code> para iniciar o banco de dados PostgreSQL.
                    </li>
                    <li>Execute a aplicação Go.</li>
                </ol>
            </li>
            <li><strong>Frontend:</strong>
                <ol>
                    <li>Navegue até o diretório do frontend.</li>
                    <li>Instale as dependências usando <code>npm install</code> ou <code>yarn install</code>.</li>
                    <li>Inicie o servidor React com <code>npm start</code> ou <code>yarn start</code>.</li>
                </ol>
            </li>
        </ol>
    </section>
    <section>
        <p>Projeto desenvolvido juntamente com a formação de GO na Alura.</p>
    </section>
    </main>
</body>
</html>