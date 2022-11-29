package generator

import (
	. "knapset-generator/entities"
	"reflect"
	"testing"
)

func Test_generateItem(t *testing.T) {
	type args struct {
		itemID int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "id = 1",
			args: args{
				itemID: 0,
			},
		},
		{
			name: "id = 0",
			args: args{
				itemID: 1,
			},
		},
		{
			name: "id = -2",
			args: args{
				itemID: -2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateItem(tt.args.itemID)

			if err != nil {
				// If the error is wanted
				if tt.args.itemID >= 0 {
					t.Errorf("generateItem() : unexpected error %s, nbItem = %d", err, tt.args.itemID)
				}

				// or unexpected
				return
			}

			isItemCorrect(t, got, tt.args.itemID)
		})
	}
}

func isItemCorrect(t *testing.T, item Item, itemID int) {
	// Same ID from args
	if itemID < 0 || !reflect.DeepEqual(item.Id, itemID) {
		t.Errorf("item ID = %v, expected %v", item.Id, itemID)
	}

	// Value between 1 and MAX_ITEM_VALUE const
	if item.Value <= 0 {
		t.Errorf("error: value is <= 0")
	}
	if item.Value > MAX_ITEM_VALUE {
		t.Errorf("error: value = %v, should be between 1 and MAX_ITEM_VALUE (%v)", item.Value, MAX_ITEM_VALUE)
	}

	// Size should be between 1 and either args (or MAX_DEFAULT_BAG_SIZE if not defined)
	if item.Size > MAX_DEFAULT_BAG_SIZE || item.Size <= 0 {
		t.Errorf("error size = %v, should be between 1 and %v", item.Size, MAX_DEFAULT_BAG_SIZE)
	}
}

func Test_generateItems(t *testing.T) {
	type args struct {
		nbItem      int
		bagSize     int
		isErrorCase bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1 bag size & 1 item",
			args: args{
				bagSize: 1,
				nbItem:  1,
			},
		},
		{
			name: "10 bag size & 10 items",
			args: args{
				bagSize: 10,
				nbItem:  10,
			},
		},
		{
			name: "10 bag size & MAX_DEFAULT_NB_ITEM items",
			args: args{
				bagSize: 10,
				nbItem:  MAX_DEFAULT_NB_ITEM,
			},
		},
		{
			name: "MAX_DEFAULT_BAG_SIZE bag size & 5 items",
			args: args{
				bagSize: MAX_DEFAULT_BAG_SIZE,
				nbItem:  5,
			},
		},
		{
			name: "MAX_DEFAULT_BAG_SIZE bag size & MAX_DEFAULT_NB_ITEM items",
			args: args{
				bagSize: MAX_DEFAULT_BAG_SIZE,
				nbItem:  MAX_DEFAULT_NB_ITEM,
			},
		},
		{
			name: "4 bag size & 5 items",
			args: args{
				bagSize: 4,
				nbItem:  5,
			},
		},
		{
			name: "5 bag size & 1 item",
			args: args{
				bagSize: 5,
				nbItem:  1,
			},
		},
		{
			name: "5 bag size & -1 item",
			args: args{
				bagSize:     5,
				nbItem:      -1,
				isErrorCase: true,
			},
		},
		{
			name: "-50 bag size & 12 items",
			args: args{
				bagSize:     -50,
				nbItem:      12,
				isErrorCase: true,
			},
		},
		{
			name: "-12 bag size & -3 item",
			args: args{
				bagSize:     -12,
				nbItem:      -3,
				isErrorCase: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateItems(tt.args.nbItem, tt.args.bagSize)

			// Handling wanted errors
			if tt.args.isErrorCase {
				if err == nil {
					t.Errorf("generateItems() : expected an error but got nil. nbItem = %d, bagSize = %d", tt.args.nbItem, tt.args.bagSize)
				}
				return
			}

			// Handling unwanted errors
			if err != nil {
				t.Errorf("GenerateNewKnapSet() : expected no error but got %s. nbItem = %d, bagSize = %d", err, tt.args.nbItem, tt.args.bagSize)
			}

			// Right number of items
			if len(got) != tt.args.nbItem {
				t.Errorf("generateItems() : expected %v items, got %v", tt.args.nbItem, len(got))
			}

			// Correct items
			for id, item := range got {
				isItemCorrect(t, item, id)
			}
		})
	}
}

func TestGenerateNewKnapSet(t *testing.T) {
	type args struct {
		bagSize     int
		nbItem      int
		isErrorCase bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "0 bag size & 0 item", // default no args
			args: args{
				bagSize: 0,
				nbItem:  0,
			},
		},
		{
			name: "0 bag size & 10 item",
			args: args{
				bagSize: 0,
				nbItem:  10,
			},
		},
		{
			name: "10 bag size & 0 item",
			args: args{
				bagSize: 10,
				nbItem:  0,
			},
		},
		{
			name: "1 bag size & 1 item",
			args: args{
				bagSize: 1,
				nbItem:  1,
			},
		},
		{
			name: "10 bag size & 10 item",
			args: args{
				bagSize: 10,
				nbItem:  10,
			},
		},
		{
			name: "5 bag size & -1 item",
			args: args{
				bagSize:     5,
				nbItem:      -1,
				isErrorCase: true,
			},
		},
		{
			name: "-50 bag size & 12 items",
			args: args{
				bagSize:     -50,
				nbItem:      12,
				isErrorCase: true,
			},
		},
		{
			name: "-12 bag size & -3 item",
			args: args{
				bagSize:     -12,
				nbItem:      -3,
				isErrorCase: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateNewKnapSet(tt.args.bagSize, tt.args.nbItem)

			// Handling wanted errors
			if tt.args.isErrorCase {
				if err == nil {
					t.Errorf("GenerateNewKnapSet() : expected an error but got nil. nbItem = %d, bagSize = %d", tt.args.nbItem, tt.args.bagSize)
				}
				return
			}

			// Handling unwanted errors
			if err != nil {
				t.Errorf("GenerateNewKnapSet() : expected no error but got %s. nbItem = %d, bagSize = %d", err, tt.args.nbItem, tt.args.bagSize)
			}

			// Right bag size
			if tt.args.bagSize == 0 {
				if got.BagSize > MAX_DEFAULT_BAG_SIZE || got.BagSize == 0 {
					t.Errorf("GenerateNewKnapSet() : expected MAX_DEFAULT_BAG_SIZE (%v) bag size, got %v", MAX_DEFAULT_BAG_SIZE, got.BagSize)
				}
			} else {
				if got.BagSize != tt.args.bagSize || got.BagSize < 0 {
					t.Errorf("GenerateNewKnapSet() : expected %v bag size, got %v", tt.args.bagSize, got.BagSize)
				}
			}

			// Right number of items
			if tt.args.nbItem == 0 {
				if len(got.Items) > MAX_DEFAULT_NB_ITEM || len(got.Items) == 0 {
					t.Errorf("GenerateNewKnapSet() : expected MAX_DEFAULT_NB_ITEM (%v) items, got %v", MAX_DEFAULT_NB_ITEM, len(got.Items))
				}
			} else {
				if len(got.Items) != tt.args.nbItem {
					t.Errorf("GenerateNewKnapSet() : expected %v items, got %v", tt.args.nbItem, len(got.Items))
				}
			}

			// Correct items
			for id, item := range got.Items {
				isItemCorrect(t, item, id)
			}
		})
	}
}
