package generator

import (
	. "generator/entities"
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
		// Item Size 1
		{"size = 1", args{0, 1}},

		// default Item Size
		{"size = MAX_DEFAULT_SIZE", args{1, MAX_DEFAULT_SIZE}},

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
		t.Errorf("item ID = %v, expected %v", item, itemID)
	}

	// Value between 1 and MAX_ITEM_VALUE const
	if item.Value <= 0 {
		t.Errorf("value is <= 0")
	}
	if item.Value > MAX_ITEM_VALUE {
		t.Errorf("value = %v, should be between 1 and %v", item, MAX_ITEM_VALUE)
	}

	// Size should be between 1 and either args (or MAX_DEFAULT_SIZE if not defined)
	if item.Size > maxSize || item.Size == 0 {
		t.Errorf("size = %v, should be between 1 and %v", item, maxSize)
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
		// 1 item & 1 bag size
		{
			name: "1 item & 1 bag size",
			args: args{
				nbItem:  1,
				bagSize: 1,
			},
		},

		// 10 item & 10 bag size
		{
			name: "10 item & 0 bag size",
			args: args{
				nbItem:  10,
				bagSize: 10,
			},
		},

		// 5 item & 4 bag size
		{
			name: "5 item & 4 bag size",
			args: args{
				nbItem:  5,
				bagSize: 4,
			},
		},

		// 1 item & 5 bag size
		{
			name: "1 item & 5 bag size",
			args: args{
				nbItem:  1,
				bagSize: 5,
			},
		},

		// TODO error cases ?
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateItems(tt.args.nbItem, tt.args.bagSize)

			// Right number of items
			if len(got) != tt.args.nbItem {
				t.Errorf("generateItems() = %v, expected %v items, got %v", got, tt.args.nbItem, len(got))
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
		knapSet KnapSet
		nbItem  int
	}
	tests := []struct {
		name string
		args args
		want KnapSet
	}{
		// TODO: Add test cases. How to use cmd args in tests ?
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateNewKnapSet(tt.args.knapSet, tt.args.nbItem); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateNewKnapSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
