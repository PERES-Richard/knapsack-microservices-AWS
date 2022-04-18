package generator

import (
	. "knapset-generator/entities"
	"reflect"
	"testing"
)

func Test_generateItem(t *testing.T) {
	type args struct {
		itemID  int
		maxSize int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "size = 1",
			args: args{
				maxSize: 1,
				itemID:  0,
			},
		},

		// default Item Size
		{
			name: "size = default",
			args: args{
				maxSize: MAX_DEFAULT_BAG_SIZE,
				itemID:  1,
			},
		},
		// TODO error cases ?
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateItem(tt.args.itemID, tt.args.maxSize)

			isItemCorrect(t, got, tt.args.itemID, tt.args.maxSize)
		})
	}
}

func isItemCorrect(t *testing.T, item Item, itemID int, maxSize int) {
	// Same ID from args
	if !reflect.DeepEqual(item.Id, itemID) {
		t.Errorf("item ID = %v, expected %v", item.Id, itemID)
	}

	// Value between 1 and MAX_ITEM_VALUE const
	if item.Value <= 0 {
		t.Errorf("value is <= 0")
	}
	if item.Value > MAX_ITEM_VALUE {
		t.Errorf("value = %v, should be between 1 and MAX_ITEM_VALUE (%v)", item.Value, MAX_ITEM_VALUE)
	}

	// Size should be between 1 and either args (or MAX_DEFAULT_BAG_SIZE if not defined)
	if item.Size > maxSize || item.Size == 0 {
		t.Errorf("size = %v, should be between 1 and %v", item.Size, maxSize)
	}
}

func Test_generateItems(t *testing.T) {
	type args struct {
		nbItem  int
		bagSize int
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
			name: "10 bag size & 10 item",
			args: args{
				bagSize: 10,
				nbItem:  10,
			},
		},
		{
			name: "10 bag size & MAX_DEFAULT_NB_ITEM item",
			args: args{
				bagSize: 10,
				nbItem:  MAX_DEFAULT_NB_ITEM,
			},
		},
		{
			name: "MAX_DEFAULT_BAG_SIZE bag size & 5 item",
			args: args{
				bagSize: MAX_DEFAULT_BAG_SIZE,
				nbItem:  5,
			},
		},
		{
			name: "MAX_DEFAULT_BAG_SIZE bag size & MAX_DEFAULT_NB_ITEM item",
			args: args{
				bagSize: MAX_DEFAULT_BAG_SIZE,
				nbItem:  MAX_DEFAULT_NB_ITEM,
			},
		},
		{
			name: "4 bag size & 5 item",
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

		// TODO error cases ?
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateItems(tt.args.nbItem, tt.args.bagSize)

			// Right number of items
			if len(got) != tt.args.nbItem {
				t.Errorf("generateItems() : expected %v items, got %v", tt.args.nbItem, len(got))
			}

			// Correct items
			for id, item := range got {
				isItemCorrect(t, item, id, tt.args.bagSize)
			}
		})
	}
}

func TestGenerateNewKnapSet(t *testing.T) {
	type args struct {
		bagSize int
		nbItem  int
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

		// TODO error cases ?
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateNewKnapSet(tt.args.bagSize, tt.args.nbItem)

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
				var bagMaxSize int
				if tt.args.bagSize == 0 {
					bagMaxSize = MAX_DEFAULT_BAG_SIZE
				} else {
					bagMaxSize = tt.args.bagSize
				}

				isItemCorrect(t, item, id, bagMaxSize)
			}
		})
	}
}
