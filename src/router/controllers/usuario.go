package controllers

import "net/http"

func CriarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Criando usuário"))
}
func BuscarUsuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Buscando todos os usuários"))
}
func BuscarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Buscando usuário específico"))
}
func AlterarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Modificando usuário"))
}
func DeletarUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Deletando usuário"))
}
