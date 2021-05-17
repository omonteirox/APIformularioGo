package models

import "time"

type Usuario struct {
	ID             uint64    `json:"id,omitempty"`
	Nome           string    `json:"nome,omitempty"`
	CPF            string    `json:"cpf,omitempty"`
	DataNascimento string    `json:"dataNascimento,omitempty"`
	CriadoEm       time.Time `json:"CriadoEm,omitempty"`
}
