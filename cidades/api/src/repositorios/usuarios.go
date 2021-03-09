package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)
NascidoEm
//Cidades representa um repositório de cidades
type Cidades struct {
	db *sql.DB
}

// NovoRepositorioDeCidades cria um repositório de usuários
func NovoRepositorioDeCidades(db *sql.DB) *Cidades {
	return &Cidades{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Cidades) Criar(cidade modelos.Cidades) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into cidades (nome, cpf, telefone, email, rua, complemento, codigo, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(cidade.Nome, cidade.Cpf, cidade.Email, cidade.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos os usuários que atendem um filtro de nome ou cpf
func (repositorio Cidades) Buscar(nomeOuCpf string) ([]modelos.Cidade, error) {
	nomeOuCpf = fmt.Sprintf("%%%s%%", nomeOuCpf) // %nomeOuCpf%

	linhas, erro := repositorio.db.Query(
		"select id, nome, cpf, telefone, email, rua, complemento, codigo, nascidoEm from cidades where nome LIKE ? or cpf LIKE ?",
		nomeOuCpf, nomeOuCpf,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var cidades []modelos.Cidade

	for linhas.Next() {
		var cidade modelos.Cidade

		if erro = linhas.Scan(
			&cidade.ID,
			&cidade.Nome,
			&cidade.Cpf,
			&cidade.Telefone,
			&cidade.Email,
			&cidade.Rua,
			&cidade.Complemento,
			&cidade.Codigo,
			&cidade.NascidoEm,
		); erro != nil {
			return nil, erro
		}

		cidades = append(cidades, cidade)
	}

	return cidades, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio Cidades) BuscarPorID(ID uint64) (modelos.Cidade, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, cpf, telefone, email, rua, complemento, codigo, nascidoEm from cidades where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.Cidade{}, erro
	}
	defer linhas.Close()

	var cidade modelos.Cidade

	if linhas.Next() {
		if erro = linhas.Scan(
			&cidade.ID,
			&cidade.Nome,
			&cidade.Cpf,
			&cidade.Email,
			&cidade.NascidoEm,
		); erro != nil {
			return modelos.Cidade{}, erro
		}
	}

	return cidade, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Cidades) Atualizar(ID uint64, cidade modelos.cidade) error {
	statement, erro := repositorio.db.Prepare(
		"update cidades set nome = ?, cpf = ?, telefone = ?, email = ?, rua = ?, complemento = ?, codigo = ?,  where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(cidade.Nome, cidade.pf, cidade.Email, ID); erro != nil {
		return erroC
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Cidades) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from cidades where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Cidades) BuscarPorEmail(email string) (modelos.Cidade, error) {
	linha, erro := repositorio.db.Query("select id, senha from cidades where email = ?", email)
	if erro != nil {
		return modelos.Cidade{}, erro
	}
	defer linha.Close()

	var cidade modelos.Cidade

	if linha.Next() {
		if erro = linha.Scan(&cidade.ID, &cidade.Senha); erro != nil {
			return modelos.Cidade{}, erro
		}
	}

	return cidade, nil

}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Cidades) BuscarSenha(cidadeID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from cidades where id = ?", cidadeID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var cidade modelos.Cidade

	if linha.Next() {
		if erro = linha.Scan(&cidade.Senha); erro != nil {
			return "", erro
		}
	}

	return cidade.Senha, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Cidades) AtualizarSenha(cidadeID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update cidades set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, cidadeID); erro != nil {
		return erro
	}

	return nil
}
