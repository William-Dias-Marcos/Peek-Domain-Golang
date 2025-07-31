
<h1 align="center">PeekDomain</h1>

PeekDomain é uma ferramenta de linha de comando simples e eficiente escrita em Go, que permite consultar os IPs e os servidores DNS (nameservers) de qualquer domínio diretamente pelo terminal.

-----

### ⚙️ Funcionalidades

- [x] Buscar os **IPs (IPv4 e IPv6)** associados a um domínio
- [x] Consultar os **nameservers (servidores DNS)** de um domínio
- [x] Interface amigável via terminal com comandos separados

-----

### 🛠 Tecnologias

- [Go](https://golang.org) — Linguagem principal
- [urfave/cli](https://github.com/urfave/cli) — package para CLI
- `net` package — Resolução DNS usando a standard library

-----

### 🚀 Como Rodar

Antes de rodar o PeekDomain, certifique-se de ter o Go instalado. Feito isso, você poderá executá-lo facilmente com os seguintes comandos:

1.  **Clone o repositório:**

    ```bash
    git clone https://github.com/seu-usuario/PeekDomain.git
    cd Peek-Domain-Golang
    ```

2.  **Compile o executável:**

    ```bash
    go build -o peekdomain
    ```

3.  **Execute o PeekDomain:**

      * Para obter os endereços IP:
        ```bash
        ./peekdomain ip --host google.com
        ```
      * Para obter os nameservers:
        ```bash
        ./peekdomain server --host google.com
        ```

-----

## 👨🏼‍💻 Autor

William Dias Marcos

 <a href = "mailto:william.diasmarcos@gmail.com"><img src="https://img.shields.io/badge/-Gmail-%23333?style=for-the-badge&logo=gmail&logoColor=white"        target="_blank"></a>
 <a href="https://www.linkedin.com/in/william-dias-marcos" target="_blank"><img src="https://img.shields.io/badge/-LinkedIn-%230077B5?style=for-the-badge&logo=linkedin&logoColor=white" target="_blank"></a>
