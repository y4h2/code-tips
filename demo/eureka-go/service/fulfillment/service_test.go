package fulfillment

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetFulfillmentStatusReturns200ForExistingSKU(t *testing.T) {
	assert := assert.New(t)

	service := NewService()

	targetSKU := "THINGAMAJIG12"

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/skus/"+targetSKU, nil)

	request = mux.SetURLVars(request, map[string]string{"sku": targetSKU})
	service.getFullfillmentStatusHandler(recorder, request)

	var detail fulfillmentStatus
	assert.Equal(http.StatusOK, recorder.Code)

	payload, err := ioutil.ReadAll(recorder.Body)
	assert.NoError(err)
	err = json.Unmarshal(payload, &detail)
	assert.NoError(err)

	assert.Equal(detail, fulfillmentStatus{
		SKU:             "THINGAMAJIG12",
		ShipsWithin:     14,
		QuantityInStock: 100,
	})
}
