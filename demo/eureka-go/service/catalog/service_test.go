package catalog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetDetailsForCatalogItemReturnsProperData(t *testing.T) {
	assert := assert.New(t)

	targetSKU := "THINGAMAJIG12"

	mockCtl := gomock.NewController(t)
	mockClient := NewMockFulfillmentClient(mockCtl)
	mockClient.EXPECT().getFulfillmentStatus(targetSKU).Return(fulfillmentStatus{
		SKU:             targetSKU,
		ShipsWithin:     99,
		QuantityInStock: 1000,
	}, nil)

	service := NewService(mockClient)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/catalog/"+targetSKU, nil)

	request = mux.SetURLVars(request, map[string]string{"sku": targetSKU})
	service.getCatalogItemDetailsHandler(recorder, request)

	var detail catalogItem

	assert.Equal(recorder.Code, http.StatusOK)

	payload, err := ioutil.ReadAll(recorder.Body)
	assert.NoError(err)
	err = json.Unmarshal(payload, &detail)
	assert.NoError(err)

	assert.Equal(detail.QuantityInStock, 1000)
	assert.Equal(detail.ShipsWithin, 99)
	assert.Equal(detail.SKU, "THINGAMAJIG12")
	assert.Equal(detail.ProductID, 1)
}
