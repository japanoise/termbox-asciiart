package asciiart

import termbox "github.com/nsf/termbox-go"

type Ascii struct {
	Data [][]AsciiTile
}

type AsciiTile struct {
	C  string
	Bg termbox.Attribute
	Fg termbox.Attribute
}

func (a *Ascii) DrawAscii(x, y int) {
	for iy, tileList := range a.Data {
		for ix, tile := range tileList {
			termbox.SetCell(x+ix, y+iy, rune(tile.C[0]), tile.Fg, tile.Bg)
		}
	}
}
