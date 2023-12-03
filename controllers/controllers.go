// controllers.go
package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"PROYECTO/models"
	repositorio "PROYECTO/repository"
)

var (
	updateQuery = "UPDATE f1_teams SET %s WHERE id=:id;"
	deleteQuery = "DELETE FROM f1_teams WHERE id=$1;"
	selectQuery = "SELECT id, Equipo, Driver_1, Driver_2, Carro, Puntos, CampeonatoConstructores, Clasificacion FROM f1_teams WHERE id=$1;"
	listQuery   = "SELECT id, Equipo, Driver_1, Driver_2, Carro, Puntos, CampeonatoConstructores, Clasificacion FROM f1_teams LIMIT $1 OFFSET $2;"
	createQuery = "INSERT INTO f1_teams (Equipo, Driver_1, Driver_2, Carro, Puntos, CampeonatoConstructores, Clasificacion) VALUES (:Equipo, :Driver_1, :Driver_2, :Carro, :Puntos, :CampeonatoConstructores, :Clasificacion) RETURNING id;"
)

type Controller struct {
	repo repositorio.Repository[models.F1Teams]
}

func NewController(repo repositorio.Repository[models.F1Teams]) (*Controller, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &Controller{
		repo: repo,
	}, nil
}

func (c *Controller) UpdateF1Team(reqBody []byte, id string) error {
	nuevosValores := make(map[string]interface{})
	err := json.Unmarshal(reqBody, &nuevosValores)
	if err != nil {
		log.Printf("fallo al actualizar un equipo de F1, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un equipo de F1, con error: %s", err.Error())
	}

	if len(nuevosValores) == 0 {
		log.Println("no se proporcionaron nuevos valores para actualizar el equipo de F1")
		return fmt.Errorf("no se proporcionaron nuevos valores para actualizar el equipo de F1")
	}

	query := buildUpdateQuery(nuevosValores)
	nuevosValores["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValores)
	if err != nil {
		log.Printf("fallo al actualizar un equipo de F1, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un equipo de F1, con error: %s", err.Error())
	}
	return nil
}

func buildUpdateQuery(nuevosValores map[string]interface{}) string {
	columns := []string{}
	for key := range nuevosValores {
		columns = append(columns, fmt.Sprintf("%s=:%s", key, key))
	}
	columnsString := strings.Join(columns, ",")
	return fmt.Sprintf(updateQuery, columnsString)
}

func (c *Controller) DeleteF1Team(id string) error {
	err := c.repo.Delete(context.TODO(), deleteQuery, id)
	if err != nil {
		log.Printf("fallo al eliminar un equipo de F1, con error: %s", err.Error())
		return fmt.Errorf("fallo al eliminar un equipo de F1, con error: %s", err.Error())
	}
	return nil
}

func (c *Controller) ReadF1Team(id string) ([]byte, error) {
	equipo, err := c.repo.Read(context.TODO(), selectQuery, id)
	if err != nil {
		log.Printf("fallo al leer un equipo de F1, con error:1 %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un equipo de F1, con error: %s", err.Error())
	}

	equipoJSON, err := json.Marshal(equipo)
	if err != nil {
		log.Printf("fallo al leer un equipo de F1, con error:2 %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un equipo de F1, con error: %s", err.Error())
	}
	return equipoJSON, nil
}

func (c *Controller) ListF1Teams(limit, offset int) ([]byte, error) {
	equipos, _, err := c.repo.List(context.TODO(), listQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer equipos de F1, con error1: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer equipos de F1, con error: %s", err.Error())
	}

	jsonEquipos, err := json.Marshal(equipos)
	if err != nil {
		log.Printf("fallo al leer equipos de F1, con error2: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer equipos de F1, con error: %s", err.Error())
	}
	return jsonEquipos, nil
}

func (c *Controller) CreateF1Team(reqBody []byte) (int64, error) {
	nuevoEquipo := &models.F1Teams{}
	err := json.Unmarshal(reqBody, nuevoEquipo)
	if err != nil {
		log.Printf("fallo al crear un nuevo equipo de F1, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo equipo de F1, con error: %s", err.Error())
	}

	valoresColumnasNuevoEquipo := map[string]interface{}{
		"Equipo":                  nuevoEquipo.Equipo,
		"Driver_1":                nuevoEquipo.Driver_1,
		"Driver_2":                nuevoEquipo.Driver_2,
		"Carro":                   nuevoEquipo.Carro,
		"Puntos":                  nuevoEquipo.Puntos,
		"CampeonatoConstructores": nuevoEquipo.CampeonatoConstructores,
		"Clasificacion":           nuevoEquipo.Clasificacion,
	}

	nuevoID, err := c.repo.Create(context.TODO(), createQuery, valoresColumnasNuevoEquipo)
	if err != nil {
		log.Printf("fallo al crear un nuevo equipo de F1, con errores: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear un nuevo equipo de F1, con error: %s", err.Error())
	}
	return nuevoID, nil
}
