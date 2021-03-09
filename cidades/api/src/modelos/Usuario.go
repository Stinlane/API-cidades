package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// cidade representa um usuário utilizando a rede social
type Cidades struct {
	ID          uint64    `json:"id,omitempty"`
	Nome        string    `json:"nome,omitempty"`
	Cpf         string    `json:"cpf,omitempty"`
	Telefone    string    `json:"telefone,omitempty"`
	Email       string    `json:",omitempty"`
	Rua         string    `json:"rua,omitempty"`
	Complemento string    `json:"complemento,omitempty"`
	Codigo      string    `json:"codigo,omitempty"`
	Senha       string    `json:"senha,omitempty"`
	NascidoEm   time.Time `json:"NascidoEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (cidade *Cidades) Preparar(etapa string) error {
	if erro := cidade.validar(etapa); erro != nil {
		return erro
	}

	if erro := cidade.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}

func (cidade *Cidades) validar(etapa string) error {
	if cidade.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if cidade.Cpf == "" {
		return errors.New("O cpf é obrigatório e não pode estar em branco")
	}

	if cidade.Telefone == "" {
		return errors.New("O telefone é obrigatório e não pode estar em branco")
	}

	if cidade.Rua == "" {
		return errors.New("A rua é obrigatório e não pode estar em branco")
	}

	if cidade.Codigo == "" {
		return errors.New("O codigo é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(cidade.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && cidade.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (cidade *Cidades) formatar(etapa string) error {
	cidade.Nome = strings.TrimSpace(cidade.Nome)
	cidade.Cpf = strings.TrimSpace(cidade.Cpf)
	cidade.Email = strings.TrimSpace(cidade.Email)
	cidade.Rua = strings.TrimSpace(cidade.Rua)
	cidade.Codigo = strings.TrimSpace(cidade.Codigo)
	cidade.Telefone = strings.TrimSpace(cidade.Telefone)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(cidade.Senha)
		if erro != nil {
			return erro
		}

		cidade.Senha = string(senhaComHash)
	}

	return nil
}
