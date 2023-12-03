// models.go
package models

type F1Teams struct {
	ID                      int64  `db:"id" json:"id"`
	Equipo                  string `db:"equipo" json:"equipo"`
	Driver_1                string `db:"driver_1" json:"driver_1"`
	Driver_2                string `db:"driver_2" json:"driver_2"`
	Carro                   string `db:"carro" json:"carro"`
	Puntos                  int    `db:"puntos" json:"puntos"`
	CampeonatoConstructores int    `db:"campeonatoconstructores" json:"campeonatoconstructores"`
	Clasificacion           int    `db:"clasificacion" json:"clasificacion"`
}
