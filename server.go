package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/natajonasdacoliveira/credisis-unidades-federativas/db"
	"github.com/natajonasdacoliveira/credisis-unidades-federativas/model"
	"github.com/natajonasdacoliveira/credisis-unidades-federativas/tools"
)

func loginHandler(c echo.Context) error {
	jsonMap := make(map[string]interface{})

	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return echo.ErrUnauthorized
	} else {
		email := jsonMap["email"].(string)
		password := jsonMap["password"].(string)

		if password == "" || email == "" {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		user := model.User{Email: email, Password: password}

		ID, err := db.Login(user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return err
		}

		token, err := tools.CreateToken(ID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

func getMunicipiosHandler(c echo.Context) error {
	municipios, err := db.GetAllMunicipios()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, municipios)
}

func postMunicipiosHandler(c echo.Context) error {
	jsonMap := make(map[string]interface{})
	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return err
	} else {
		nome := jsonMap["nome"].(string)
		prefeito := jsonMap["prefeito"].(string)
		populacao := jsonMap["populacao"].(string)
		estado := jsonMap["id_estado_fk"].(string)

		var populacaoInt uint64
		populacaoInt = 0

		if populacao != "" {
			populacaoInt, err = strconv.ParseUint(populacao, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		estadoInt, err := strconv.ParseUint(estado, 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		if nome == "" || estadoInt < 1 {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		municipio := model.Municipio{
			Nome:      nome,
			Prefeito:  prefeito,
			Populacao: populacaoInt,
			IDEstado:  estadoInt,
		}

		err = db.CreateMunicipio(municipio)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}

		return c.String(http.StatusOK, "Município criado com sucesso")
	}
}

func updateMunicipiosHandler(c echo.Context) error {
	jsonMap := make(map[string]interface{})

	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return err
	} else {
		id := jsonMap["id"].(string)
		nome := jsonMap["nome"].(string)
		prefeito := jsonMap["prefeito"].(string)
		populacao := jsonMap["populacao"].(string)
		estado := jsonMap["id_estado_fk"].(string)

		var populacaoInt uint64
		populacaoInt = 0

		if populacao != "" {
			populacaoInt, err = strconv.ParseUint(populacao, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		var idInt uint64
		idInt = 0

		if id != "" {
			idInt, err = strconv.ParseUint(id, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		estadoInt, err := strconv.ParseUint(estado, 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		if idInt < 1 || nome == "" || estadoInt < 1 {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		municipio := model.Municipio{
			ID:        idInt,
			Nome:      nome,
			Prefeito:  prefeito,
			Populacao: populacaoInt,
			IDEstado:  estadoInt,
		}

		err = db.UpdateMunicipio(municipio)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}

		return c.String(http.StatusOK, "Município atualizado com sucesso")
	}
}

func deleteMunicipiosHandler(c echo.Context) error {
	jsonMap := make(map[string]interface{})

	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return err
	} else {
		id := jsonMap["id"].(string)
		var idInt uint64
		idInt = 0

		if id != "" {
			idInt, err = strconv.ParseUint(id, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		if idInt < 1 {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		municipio := model.Municipio{
			ID: idInt,
		}

		err = db.DeleteMunicipio(municipio)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}

		return c.String(http.StatusOK, "Município deletado com sucesso")
	}
}

func getEstadosHandler(c echo.Context) error {
	estados, err := db.GetAllEstados()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, estados)
}

func postEstadosHandler(c echo.Context) error {
	jsonMap := make(map[string]interface{})
	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return err
	} else {
		nome := jsonMap["nome"].(string)
		sigla := jsonMap["sigla"].(string)

		if nome == "" || sigla == "" {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		estado := model.Estado{
			Nome:  nome,
			Sigla: sigla,
		}

		err = db.CreateEstado(estado)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}

		return c.String(http.StatusOK, "Estado criado com sucesso")
	}
}

func updateEstadosHandler(c echo.Context) error {
	jsonMap := make(map[string]interface{})

	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return err
	} else {
		id := jsonMap["id"].(string)
		nome := jsonMap["nome"].(string)
		sigla := jsonMap["sigla"].(string)

		var idInt uint64
		idInt = 0

		if id != "" {
			idInt, err = strconv.ParseUint(id, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		if idInt < 1 || nome == "" || sigla == "" {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		estado := model.Estado{
			ID:    idInt,
			Nome:  nome,
			Sigla: sigla,
		}

		err = db.UpdateEstado(estado)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}

		return c.String(http.StatusOK, "Estado atualizado com sucesso")
	}
}

func deleteEstadosHandler(c echo.Context) error {
	jsonMap := make(map[string]interface{})

	if err := json.NewDecoder(c.Request().Body).Decode(&jsonMap); err != nil {
		return err
	} else {
		id := jsonMap["id"].(string)
		var idInt uint64
		idInt = 0

		if id != "" {
			idInt, err = strconv.ParseUint(id, 10, 64)
			if err != nil {
				return c.String(http.StatusBadRequest, "Dados inválidos")
			}
		}

		if idInt < 1 {
			return c.String(http.StatusBadRequest, "Dados inválidos")
		}

		estado := model.Estado{
			ID: idInt,
		}

		err = db.DeleteEstado(estado)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return err
		}

		return c.String(http.StatusOK, "Estado deletado com sucesso")
	}
}

func main() {
	os.Setenv("ACCESS_SECRET", "Zq3t6w9z$C&F)J@NcRfUjXn2r5u7x!A%")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	jwtGroup := e.Group("/jwt")

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("ACCESS_SECRET")),
	}))

	e.POST("/login", loginHandler)

	e.GET("/municipios", getMunicipiosHandler)
	e.GET("/estados", getEstadosHandler)

	jwtGroup.POST("/municipios", postMunicipiosHandler)
	jwtGroup.PUT("/municipios", updateMunicipiosHandler)
	jwtGroup.DELETE("/municipios", deleteMunicipiosHandler)

	jwtGroup.POST("/estados", postEstadosHandler)
	jwtGroup.PUT("/estados", updateEstadosHandler)
	jwtGroup.DELETE("/estados", deleteEstadosHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
