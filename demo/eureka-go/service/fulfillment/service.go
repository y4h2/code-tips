package fulfillment

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Service struct {
	render *render.Render
	router *mux.Router
}

// getFullfillmentStatusHandler simulates actual fulfillment by supplying bogus values for QuantityInStock
// and ShipsWithin for any given SKU. Used to demonstrate a backing service supporting
// a primary service.
func (s *Service) getFullfillmentStatusHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	sku := vars["sku"]
	s.render.JSON(w, http.StatusOK, fulfillmentStatus{
		SKU:             sku,
		ShipsWithin:     14,
		QuantityInStock: 100,
	})
}

func (s *Service) rootHandler(w http.ResponseWriter, req *http.Request) {
	s.render.Text(w, http.StatusOK, "Fulfillment Service, see http://github.com/cloudnativego/backing-fulfillment for API.")
}

func (s *Service) InitRoutes() *mux.Router {
	s.router.HandleFunc("/", s.rootHandler).Methods("GET")
	s.router.HandleFunc("/skus/{sku}", s.getFullfillmentStatusHandler).Methods("GET")
	return s.router
}

func NewService() *Service {
	return &Service{
		render: render.New(render.Options{
			IndentJSON: true,
		}),
		router: mux.NewRouter(),
	}
}
