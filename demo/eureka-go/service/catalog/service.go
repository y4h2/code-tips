package catalog

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

//go:generate go run github.com/golang/mock/mockgen -package catalog -self_package "eureka-go/service/catalog" -destination mockgen_client.go . FulfillmentClient
type FulfillmentClient interface {
	getFulfillmentStatus(sku string) (status fulfillmentStatus, err error)
}

type Service struct {
	render            *render.Render
	fulfillmentClient FulfillmentClient
	router            *mux.Router
}

// getAllCatalogItemsHandler returns a fake list of catalog items
func (s *Service) getAllCatalogItemsHandler(w http.ResponseWriter, req *http.Request) {
	catalog := make([]catalogItem, 2)
	catalog[0] = fakeItem("ABC1234")
	catalog[1] = fakeItem("STAPLER99")
	s.render.JSON(w, http.StatusOK, catalog)
}

// getCatalogItemDetailsHandler returns a fake catalog item. The key takeaway here
// is that we're using a backing service to get fulfillment status for the individual
// item.
func (s *Service) getCatalogItemDetailsHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	sku := vars["sku"]
	status, err := s.fulfillmentClient.getFulfillmentStatus(sku)
	if err == nil {
		s.render.JSON(w, http.StatusOK, catalogItem{
			ProductID:       1,
			SKU:             sku,
			Description:     "This is a fake product",
			Price:           1599, // $15.99
			ShipsWithin:     status.ShipsWithin,
			QuantityInStock: status.QuantityInStock,
		})
	} else {
		s.render.JSON(w, http.StatusInternalServerError, fmt.Sprintf("Fulfillment Client error: %s", err.Error()))
	}
}

func (s *Service) rootHandler(w http.ResponseWriter, req *http.Request) {
	s.render.Text(w, http.StatusOK, "Catalog Service, see http://github.com/cloudnativego/backing-catalog for API.")
}

func (s *Service) InitRoutes() *mux.Router {
	s.router.HandleFunc("/", s.rootHandler).Methods("GET")
	s.router.HandleFunc("/catalog", s.getAllCatalogItemsHandler).Methods("GET")
	s.router.HandleFunc("/catalog/{sku}", s.getCatalogItemDetailsHandler).Methods("GET")
	return s.router
}

func NewService(client FulfillmentClient) *Service {
	return &Service{
		render: render.New(render.Options{
			IndentJSON: true,
		}),
		fulfillmentClient: client,
		router:            mux.NewRouter(),
	}
}

func fakeItem(sku string) (item catalogItem) {
	item.SKU = sku
	item.Description = "This is a fake product"
	item.Price = 1599
	item.QuantityInStock = 75
	item.ShipsWithin = 14
	return
}
