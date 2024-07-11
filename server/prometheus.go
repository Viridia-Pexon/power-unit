package server

import (
	"github.com/ansrivas/fiberprometheus/v2"
)

/*
exposes prometheus endpoint for request metrics
---
aktiviert Endpunkt für Prometheis
*/
func (gp *Powerunit) InitMetrics(path string) {
	prometheus := fiberprometheus.New("gis-auth-proxy")
	prometheus.RegisterAt(gp.APP, path)
	gp.APP.Use(prometheus.Middleware)
}
