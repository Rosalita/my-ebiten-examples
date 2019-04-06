package charmenu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharListToCharGrid(t *testing.T) {

	oneRow := [][]CharBox{
		[]CharBox{
			CharBox{Char: "a"}, CharBox{Char: "b"}, CharBox{Char: "c"},
		},
	}

	twoRows := [][]CharBox{
		[]CharBox{
			CharBox{Char: "a"}, CharBox{Char: "b"}, CharBox{Char: "c"},
		},
		[]CharBox{
			CharBox{Char: "d"}, CharBox{Char: "e"},
		},
	}

	tests := []struct {
		charList   string
		lineLength int
		result     [][]CharBox
	}{
		{"abc", 3, oneRow},
		{"abcde", 3, twoRows},
	}

	for _, test := range tests {

		result := charListToCharGrid(test.charList, test.lineLength)
		assert.Equal(t, test.result, result)

	}
}

func TestSplitIntoLines(t *testing.T) {

	tests := []struct {
		charList   string
		lineLength int
		result     []string
	}{
		{"abc", 3, []string{"abc"}},
		{"abcdef", 3, []string{"abc", "def"}},
		{"12345", 2, []string{"12", "34", "5"}},
	}

	for _, test := range tests {
		result := splitIntoLines(test.charList, test.lineLength)
		assert.Equal(t, test.result, result)
	}

}

func TestInitGrid(t *testing.T) {

	grid := [][]CharBox{
		[]CharBox{
			CharBox{Char: "a"}, CharBox{Char: "b"}, CharBox{Char: "c"},
		},
		[]CharBox{
			CharBox{Char: "d"}, CharBox{Char: "e"},
		},
	}

	expected := [][]CharBox{
		[]CharBox{
			CharBox{Char: "a"}, CharBox{Char: "b"}, CharBox{Char: "c"},
		},
		[]CharBox{
			CharBox{Char: "d"}, CharBox{Char: "e"},
		},
	}

	tests := []struct {
		grid   [][]CharBox
		width  int
		height int
		result [][]CharBox
	}{
		{grid, 10, 20, expected},
	}

	for _, test := range tests {
		result := initGrid(test.grid, test.width, test.height)
		for y, row := range result {
			for x := range row {
				assert.NotNil(t, result[y][x].image)
				assert.NotNil(t, result[y][x].TxtX)
				assert.NotNil(t, result[y][x].TxtY)
			}
		}
	}
}
