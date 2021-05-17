package rotas

import (
	"api/src/router/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		Uri:    "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
	},
	{
		Uri:    "/usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuarios,
	},
	{
		Uri:    "/usuarios/{usuarioID}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarUsuario,
	},

	{
		Uri:    "/usuarios/{usuarioID}",
		Metodo: http.MethodPut,
		Funcao: controllers.AlterarUsuario,
	},
	{
		Uri:    "/usuarios/{usuarioID}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarUsuario,
	},
}
