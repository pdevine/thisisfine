package main

import (
	"time"
	"math/rand"

	sprite "github.com/pdevine/go-asciisprite"
	tm "github.com/pdevine/go-asciisprite/termbox"
)

var allSprites sprite.SpriteGroup
var Width int
var Height int

const (
	LEFT = iota
	RIGHT
)

const DogCostume1 = `
      dd
     ddBBd
     BBddd
    dddddttt
    dddtttttt
     ttttttttt
     Btttttttw
    BBttwwttwBB
   BBBtwBBwtwBB
  BBBBtwBBwttwttBBB
  BBBBttwwtttttttBB
  BBBBttttttBtttttt
   BBttttttttBtttt
     tTtTttttt
     tTtTttttt
    dtTtTttttt
    dtTttTtttT
    dtTTTTTTtTT
    ddtttttTdtT
     dddddtTdtTT
      d  dtTT
      d  d d
      d  d`

const TableCostume1 = `
   wwww
   wBBwww
   wwww w
dddwwwwwwdddddddddddd
ddddwwddddddddddddddd
 ddddddddddddddddddddd
  d                d
  d                d
  d                d`

const SmokeCostume1 = `gggggggggggggggggggggggggggggggggGGGgggggg
ggggggggggggggggggggggggggggggGGGGGGGGgggg
gggggggggggggggggGGgggGgggGGGGGGGGGGGGGGGg
gggggggggggggggGGGGGGGGGGGGGGGGGGGGGGGGGGG
GGgggggggggggGGGGGGGGGGGGGGGGGGGGGGGGGGGGG
GGGggggggggGGGGGGGGGGGGGGGGGGGGGGGGG GGGGG
GGGGGGGGGGGGGGGGGGGGGGGGGG             GGG
GGGGGGGGGGGGGGGGG GGGGGGG                G
GGGGGGGGGGGGGGGG   GGG
 GGGGGGGGGGG
    GGGG`

type Dog struct {
	sprite.BaseSprite
}

type Table struct {
	sprite.BaseSprite
}

type Fire struct {
	sprite.BaseSprite
	Points int
	buf [][]int
}

type Smoke struct {
	sprite.BaseSprite
	Timer   int
	TimeOut int
}

type Text struct {
	sprite.BaseSprite
	font *sprite.Font
}

func NewDog() *Dog {
	d := &Dog{BaseSprite: sprite.BaseSprite{
		Visible: true},
	}
	d.Init()

	d.RegisterEvent("resizeScreen", func() {
		d.X = Width/2 - 22
		if Width/2 % 2 != 0 {
			d.X++
		}
		d.Y = Height/2 + 8
		if Height/2 % 2 != 0 {
			d.Y++
		}
	})

	surf := sprite.NewSurfaceFromString(DogCostume1, false)
	d.BlockCostumes = append(d.BlockCostumes, &surf)
	return d
}

func NewTable() *Table {
	t := &Table{BaseSprite: sprite.BaseSprite{
		Visible: true},
	}
	t.Init()

	t.RegisterEvent("resizeScreen", func() {
		t.X = Width/2 + 10
		t.Y = Height/2 + 24
	})

	surf := sprite.NewSurfaceFromString(TableCostume1, false)
	t.BlockCostumes = append(t.BlockCostumes, &surf)
	return t
}

func NewSmoke() *Smoke {
	s := &Smoke{BaseSprite: sprite.BaseSprite{
		Visible: true},
		TimeOut: 1,
	}
	s.Init()

	s.RegisterEvent("resizeScreen", func() {
		surf := sprite.NewSurfaceFromString(SmokeCostume1, true)

		yOff := 10
		bigSurf := sprite.NewSurface(Width*2, surf.Height+yOff, true)

		// fill top part w/ silver
		for y := 0; y < yOff; y++ {
			for x := 0; x < bigSurf.Width; x++ {
				bigSurf.Blocks[y][x] = 'g'
			}
		}

		for cnt := 0; cnt < Width*2 / surf.Width; cnt++ {
			bigSurf.Blit(surf, cnt*surf.Width, yOff)
		}

		s.BlockCostumes = []*sprite.Surface{&bigSurf}
	})

	return s
}

func (s *Smoke) Update() {
	s.Timer++
	if s.Timer > s.TimeOut {
		s.X--
		s.Timer = 0
	}
	if s.X <= -Width/2 {
		s.X = 0
	}
}

func NewFire(side int) *Fire {
	f := &Fire{BaseSprite: sprite.BaseSprite{
		X: 0,
		Y: 0,
		Visible: true},
		Points: 5,
	}
	f.Init()

	f.RegisterEvent("flamesHigher", func() {
		f.Points -= 1
		if f.Points == 0 {
			f.Points = 1
		}
	})

	f.RegisterEvent("flamesLower", func() {
		f.Points += 1
		if f.Points > 20 {
			f.Points = 20
		}
	})

	f.RegisterEvent("resizeScreen", func() {
		if side == RIGHT {
			f.X = Width/2 + 25
		}

		f.buf = make([][]int, Height+1)

		for cnt := range f.buf {
			f.buf[cnt] = make([]int, Width/2+1)
		}

		surf := sprite.NewSurface(Width/2, Height, true)
		f.BlockCostumes = []*sprite.Surface{&surf}
	})

	return f
}

func (f *Fire) Update() {
	for cnt := 0; cnt < int(Width/2 / f.Points); cnt++ {
		f.buf[Height-1][rand.Intn(Width/2)] = 65
	}
	surf := sprite.NewSurface(Width/2, Height, true)
	for h := 0; h < Height; h++ {
		for w := 0; w < Width/2; w++ {
			f.buf[h][w] += f.buf[h][w+1]
			f.buf[h][w] += f.buf[h+1][w]
			f.buf[h][w] += f.buf[h+1][w+1]
			f.buf[h][w] /= 4

			if f.buf[h][w] > 15 {
				surf.Blocks[h][w] = 'r'
			} else if f.buf[h][w] > 9 {
				surf.Blocks[h][w] = 'o'
			} else if f.buf[h][w] > 4 {
				surf.Blocks[h][w] = 'y'
			}
		}
	}
	f.BlockCostumes = []*sprite.Surface{&surf}
}

func NewText() *Text {
	t := &Text{BaseSprite: sprite.BaseSprite{
		X: 20,
		Y: 30,
		Visible: true},
		font: sprite.NewPakuFont(),
	}
	t.Init()

	surf := sprite.NewSurfaceFromString(t.font.BuildString("This is fine."), false)
        t.BlockCostumes = append(t.BlockCostumes, &surf)

	return t
}

func setPalette() {
	sprite.ColorMap['o'] = tm.Color214
	sprite.ColorMap['y'] = tm.Color226
	sprite.ColorMap['r'] = tm.Color197
	sprite.ColorMap['d'] = tm.Color52
	sprite.ColorMap['B'] = tm.ColorBlack
	sprite.ColorMap['w'] = tm.ColorWhite
	sprite.ColorMap['t'] = tm.Color173
	sprite.ColorMap['T'] = tm.Color130
	sprite.ColorMap['g'] = tm.ColorSilver
	sprite.ColorMap['G'] = tm.ColorGray
	sprite.ColorMap['X'] = tm.ColorBlack
}

func main() {
	err := tm.Init()
	if err != nil {
		panic(err)
	}
	defer tm.Close()

	w, h := tm.Size()
	Width = w*2
	Height = h*2

	setPalette()

	allSprites.Init(Width, Height, true)
	allSprites.Background = tm.Color187

	eventQueue := make(chan tm.Event)
	go func() {
		for {
			eventQueue <- tm.PollEvent()
		}
	}()

	d := NewDog()
	t := NewTable()
	s := NewSmoke()
	text := NewText()
	f1 := NewFire(LEFT)
	f2 := NewFire(RIGHT)
	allSprites.Sprites = append(allSprites.Sprites, d)
	allSprites.Sprites = append(allSprites.Sprites, text)
	allSprites.Sprites = append(allSprites.Sprites, t)
	allSprites.Sprites = append(allSprites.Sprites, s)
	allSprites.Sprites = append(allSprites.Sprites, f1)
	allSprites.Sprites = append(allSprites.Sprites, f2)

mainloop:
	for {
		tm.Clear(tm.Color187, tm.Color187)

		select {
		case ev := <-eventQueue:
			if ev.Type == tm.EventKey {
				if ev.Key == tm.KeyEsc {
					break mainloop
				}
				if ev.Key == tm.KeyArrowUp {
					allSprites.TriggerEvent("flamesHigher")
				} else if ev.Key == tm.KeyArrowDown {
					allSprites.TriggerEvent("flamesLower")
				}
			} else if ev.Type == tm.EventResize {
				Width = ev.Width*2
				Height = ev.Height*2
				allSprites.Init(Width, Height, true)
				allSprites.Background = tm.Color187
				allSprites.TriggerEvent("resizeScreen")
			}
		default:
			allSprites.Update()
			allSprites.Render()
			time.Sleep(60 * time.Millisecond)
		}
	}

}

