package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"PROYECTO/controllers"

	"github.com/gorilla/mux"
)

type Handler struct {
	controller *controllers.Controller
}

func NewHandler(controller *controllers.Controller) (*Handler, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler se necesita un controlador no nulo")
	}
	return &Handler{
		controller: controller,
	}, nil
}

func (h *Handler) UpdateF1Team(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un equipo de F1, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un equipo de F1, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.controller.UpdateF1Team(body, id)
	if err != nil {
		log.Printf("fallo al actualizar un equipo de F1, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un equipo de F1, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteF1Team(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	err := h.controller.DeleteF1Team(id)
	if err != nil {
		log.Printf("fallo al eliminar un equipo de F1, con error: %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("fallo al eliminar un equipo de F1 con id %s", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (h *Handler) ReadF1Team(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	equipo, err := h.controller.ReadF1Team(id)
	if err != nil {
		log.Printf("fallo al leer un equipo de F1, con error: %s", err.Error())
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("el equipo de F1 con id %s no se pudo encontrar", id)))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(equipo)
}

func (h *Handler) ListF1Teams(writer http.ResponseWriter, req *http.Request) {
	equipos, err := h.controller.ListF1Teams(100, 0)
	if err != nil {
		log.Printf("fallo al leer equipos de F1, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los equipos de F1", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(equipos)
}

func (h *Handler) CreateF1Team(writer http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo equipo de F1, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo equipo de F1", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	newID, err := h.controller.CreateF1Team(body)
	if err != nil {
		log.Println("fallo al crear un nuevo equipo de F1, con error:", err.Error())
		http.Error(writer, "fallo al crear un nuevo equipo de F1", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("ID del nuevo equipo de F1: %d", newID)))
}
