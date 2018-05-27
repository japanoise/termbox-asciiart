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

func (a *Ascii) DrawAsciiNoClobber(x, y int) {
	for iy, tileList := range a.Data {
		for ix, tile := range tileList {
			if !(tile.C == " " && tile.Fg == termbox.ColorDefault && tile.Bg == tile.Fg) {
				termbox.SetCell(x+ix, y+iy, rune(tile.C[0]), tile.Fg, tile.Bg)
			}
		}
	}
}
