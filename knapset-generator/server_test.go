package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initHttpReq(t *testing.T, bagSize int, numberOfItems int) *http.Request {
	path := fmt.Sprintf("/generate?numberOfItems=%d&bagSize=%d", numberOfItems, bagSize)
	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func Test_generateNewBag(t *testing.T) {
	type args struct {
		bagSize       int
		numberOfItems int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "bagSize = 100 & numberOfItems 12",
			args: args{
				bagSize:       100,
				numberOfItems: 12,
			},
		},
		{
			name: "bagSize = 0 & numberOfItems 0",
			args: args{
				bagSize:       0,
				numberOfItems: 0,
			},
		},
		{
			name: "bagSize = 10 & numberOfItems 100",
			args: args{
				bagSize:       10,
				numberOfItems: 100,
			},
		},
		{
			name: "bagSize = 0 & numberOfItems 5",
			args: args{
				bagSize:       0,
				numberOfItems: 5,
			},
		},
		{
			name: "bagSize = 5 & numberOfItems 0",
			args: args{
				bagSize:       5,
				numberOfItems: 0,
			},
		},
		{
			name: "bagSize = 100 & numberOfItems 0",
			args: args{
				bagSize:       100,
				numberOfItems: 0,
			},
		},
		{
			name: "bagSize = -1 & numberOfItems -1",
			args: args{
				bagSize:       -1,
				numberOfItems: -1,
			},
		},
		{
			name: "bagSize = 1 & numberOfItems 1",
			args: args{
				bagSize:       1,
				numberOfItems: 1,
			},
		},
		{
			name: "bagSize = -1 & numberOfItems 1",
			args: args{
				bagSize:       -1,
				numberOfItems: 1,
			},
		},
		{
			name: "bagSize = 1 & numberOfItems -1",
			args: args{
				bagSize:       1,
				numberOfItems: -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(generateNewBag)
			handler.ServeHTTP(rr, initHttpReq(t, tt.args.bagSize, tt.args.numberOfItems))

			//assertCorrectNumberOfItemAndBagSize(t, rr, tt.args.bagSize, tt.args.numberOfItems)
		})
	}
}

func assertCorrectNumberOfItemAndBagSize(t *testing.T, w *httptest.ResponseRecorder, bagSize int, numberOfItems int) {
	//knapset := entities.KnapSet{}
	//json.NewDecoder(w.Body).Decode(&knapset)
	//
	//// Assert bagSize
	//if bagSize < 0 { // 'Error' bagSize case
	//	// If negative bagSize then Knapset should not have any item
	//	if knapset.Items != nil {
	//		t.Errorf("incorrect bagSize (%d), expeceted 0 items, got %d", bagSize, len(knapset.Items))
	//	}
	//} else if bagSize == 0 { // Random bagSize case
	//	if knapset.BagSize == bagSize {
	//		t.Errorf("incorrect bagSize (%d), random int between 1 and %d, got %d", bagSize, generator.MAX_DEFAULT_BAG_SIZE, knapset.BagSize)
	//	}
	//} else { // Normal case
	//	if knapset.BagSize != bagSize {
	//		t.Errorf("incorrect bagSize, expeceted %d, got %d", bagSize, knapset.BagSize)
	//	}
	//}
	//
	//// Assert numberOfItems
	//if numberOfItems < 0 { // 'Error' numberOfItems case
	//	if knapset.Items != nil {
	//		t.Errorf("incorrect numberOfItems, expeceted %d (error), got %d", numberOfItems, len(knapset.Items))
	//	}
	//} else if numberOfItems == 0 { // Random numberOfItems case
	//	if knapset.Items == nil || len(knapset.Items) == numberOfItems {
	//		t.Errorf("incorrect numberOfItems (%d), expeceted random int between 1 and %d, got %d", numberOfItems, generator.MAX_DEFAULT_NB_ITEM, len(knapset.Items))
	//	}
	//} else { // Normal case
	//	if len(knapset.Items) != numberOfItems {
	//		t.Errorf("incorrect numberOfItems, expeceted %d, got %d", numberOfItems, len(knapset.Items))
	//	}
	//}
}
