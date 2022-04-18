package main

import (
	"knapset-generator/generator"
	"os"
	"strconv"
	"testing"
)

func Test_handleArgs(t *testing.T) {
	type args struct {
		bagSize   int
		nbItem    int
		printBool string
		path      string
	}
	tests := []struct {
		name       string
		args       args
		bagSizeGot int
		nbItemGot  int
	}{
		{
			name: "0 bag size & 10 item",
			args: args{
				bagSize: 0,
				nbItem:  10,
			},
		},
		{
			name: "no bag size & no item", // default no args
			args: args{
				bagSize: -1,
				nbItem:  -1,
			},
		},
		{
			name: "0 bag size & 0 item",
			args: args{
				bagSize: 0,
				nbItem:  0,
			},
		},
		{
			name: "0 bag size & 0 item + print true",
			args: args{
				bagSize:   0,
				nbItem:    0,
				printBool: "1",
			},
		},
		{
			name: "0 bag size & 0 item + print true & path = test.txt",
			args: args{
				bagSize:   0,
				nbItem:    0,
				printBool: "true",
				path:      "test.txt",
			},
		},
		{
			name: "0 bag size & 0 item + print false",
			args: args{
				bagSize:   0,
				nbItem:    0,
				printBool: "0",
			},
		},
		{
			name: "0 bag size & 0 item + print false & path = test.txt",
			args: args{
				bagSize:   0,
				nbItem:    0,
				printBool: "false",
				path:      "test.txt",
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
			name: "11 bag size & 2 item + print true",
			args: args{
				bagSize:   11,
				nbItem:    2,
				printBool: "1",
			},
		},
		{
			name: "0 bag size & 30 item + print true & path = test.txt",
			args: args{
				bagSize:   0,
				nbItem:    30,
				printBool: "true",
				path:      "test.txt",
			},
		},
		{
			name: "8 bag size & 0 item + print false",
			args: args{
				bagSize:   8,
				nbItem:    0,
				printBool: "0",
			},
		},
		{
			name: "15 bag size & 7 item + print false & path = test.txt",
			args: args{
				bagSize:   15,
				nbItem:    7,
				printBool: "false",
				path:      "test.txt",
			},
		},
		{
			name: "2*MAX_DEFAULT_BAG_SIZE bag size & 0 item",
			args: args{
				bagSize: 2 * generator.MAX_DEFAULT_BAG_SIZE,
				nbItem:  0,
			},
		},
		{
			name: "2*MAX_DEFAULT_BAG_SIZE bag size & 5 item",
			args: args{
				bagSize: 5,
				nbItem:  2 * generator.MAX_DEFAULT_BAG_SIZE,
			},
		},
		{
			name: "0 bag size & 2*MAX_DEFAULT_NB_ITEM",
			args: args{
				bagSize: 0,
				nbItem:  2 * generator.MAX_DEFAULT_NB_ITEM,
			},
		},
		{
			name: "5 bag size & 2*MAX_DEFAULT_NB_ITEM",
			args: args{
				bagSize: 5,
				nbItem:  2 * generator.MAX_DEFAULT_NB_ITEM,
			},
		},
		{
			name: "2*MAX_DEFAULT_BAG_SIZE bag size & 2*MAX_DEFAULT_BAG_SIZE item",
			args: args{
				bagSize: 2 * generator.MAX_DEFAULT_BAG_SIZE,
				nbItem:  2 * generator.MAX_DEFAULT_NB_ITEM,
			},
		},

		// TODO error case (-1) ?
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup args
			simulateArgs(&tt.args.bagSize, &tt.args.nbItem, tt.args.printBool, tt.args.path)

			bagSizeGot, nbItemGot, printBool, JSONpath := handleArgs()

			// Correct bag size
			if bagSizeGot != tt.args.bagSize {
				t.Errorf("handleArgs() bagSize = %v, expected %v", bagSizeGot, tt.args.bagSize)
			}

			// Correct nb item
			if nbItemGot != tt.args.nbItem {
				t.Errorf("handleArgs() nbItem = %v, expected %v", nbItemGot, tt.args.nbItem)
			}

			// Correct print bool
			if tt.args.printBool != "" {
				if printArg, _ := strconv.ParseBool(tt.args.printBool); printArg != printBool {
					t.Errorf("handleArgs() print bool = %v, expected %v (%v)", printBool, printArg, tt.args.printBool)
				}
			}

			// Correct path
			if tt.args.path == "" {
				if JSONpath != defaultJSONPath {
					t.Errorf("handleArgs() path = %v, expected defaultJSONPath (%v)", JSONpath, defaultJSONPath)
				}
			} else {
				if JSONpath != tt.args.path {
					t.Errorf("handleArgs() path = %v, expected %v", JSONpath, tt.args.path)
				}
			}
		})
	}
}

func simulateArgs(bagSize *int, nbItem *int, printBool string, path string) {
	os.Args = make([]string, 3)

	if *bagSize != -1 && *nbItem != -1 {
		os.Args[BAG_SIZE_ARG] = strconv.Itoa(*bagSize)
		os.Args[NB_ITEM_ARG] = strconv.Itoa(*nbItem)
	} else {
		os.Args = []string{"program"}
		*bagSize = 0
		*nbItem = 0
	}

	if printBool != "" {
		tempArr := make([]string, len(os.Args)+1)
		copy(tempArr, os.Args)
		os.Args = tempArr
		os.Args[CONSOLE_PRINT_ARG] = printBool
	}
	if path != "" {
		tempArr := make([]string, len(os.Args)+1)
		copy(tempArr, os.Args)
		os.Args = tempArr
		os.Args[JSON_PATH_ARG] = path
	}
}
