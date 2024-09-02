package main

import (
	"fmt"
	"image/color"
	"time"
	"os"
//	"github.com/hjson/hjson-go/v4"
    "encoding/json"

	//	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

const (
	max_size = 10000 
)
/*
const (
	all_colors_str = `[
      { "Ibits": 9,
	"blue_pos":    [  5,   6,   7,  -1,  -1,  -1,  -1,  -1,  -1],
	"green_pos":   [ -1,  -1,  -1,   5,   6,   7,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   5,   6,   7],
	"default_color": [0,0,0] 
      }
	]`
)
*/
const (
	all_colors_str = `[
      { "Ibits": 9,
	"blue_pos":    [  5,   6,   7,  -1,  -1,  -1,  -1,  -1,  -1],
	"green_pos":   [ -1,  -1,  -1,   5,   6,   7,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   5,   6,   7],
	"default_color": [0,0,0] 
      },
      { "ibits": 9,
	"green_pos":    [  5,   6,   7,  -1,  -1,  -1,  -1,  -1,  -1],
	"blue_pos":   [ -1,  -1,  -1,   5,   6,   7,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   5,   6,   7],
	"default_color": [0,0,0]
      },
      { "ibits": 9,
	"red_pos":    [  5,   6,   7,  -1,  -1,  -1,  -1,  -1,  -1],
	"blue_pos":   [ -1,  -1,  -1,   5,   6,   7,  -1,  -1,  -1],
	"green_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   5,   6,   7],
	"default_color":[0,0,0]
      },
      { "ibits": 9,
	"blue_pos":    [  3,   4,   5,  -1,  -1,  -1,  -1,  -1,  -1],
	"green_pos":   [ -1,  -1,  -1,   3,   4,   5,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   3,   4,   5],
	"default_color": [0,0,0]
      },
      { "ibits": 9,
	"green_pos":    [  3,   4,   5,  -1,  -1,  -1,  -1,  -1,  -1],
	"blue_pos":   [ -1,  -1,  -1,   3,   4,   5,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   3,   4,   5],
	"default_color": [0,0,0]
      },
      { "ibits": 9,
	"red_pos":    [  3,   4,   5,  -1,  -1,  -1,  -1,  -1,  -1],
	"blue_pos":   [ -1,  -1,  -1,   3,   4,   5,  -1,  -1,  -1],
	"green_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   3,   4,   5],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":    [  2, 3,  4,   5,   6,   7],
	"green_pos":   [ 1,-1,-1, -1,-1,-1,-1],
	"red_pos":     [ -1,-1, -1,-1,-1,-1],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":    [   2,   3,   4,  5,  6,  7],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":    [ -1,  -1,  -1, -1, -1, -1],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":    [   2,   3,   4,  5,  6,  7],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":    [ -1,  -1,  -1, -1, -1, -1],
	"default_color": [113,1,147]
      },
      { "ibits": 6,
	"blue_pos":    [  0,  1,  2, 3, 4, 5],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":    [  0,  1,  2, 3, 4, 5],
	"default_color": [0,0,0]
	  },
      { "ibits": 6,
	"blue_pos":    [   -1,   -1,  -1,  -1,  -1, -1],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":    [  0,  1,  2, 3, 4, 5],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"red_pos":    [  0,  1,  2, 3, 4, 5],
	"blue_pos":    [  0,  1,  2, 3, 4, 5],
	"green_pos":    [   2,   3,   4,  5,  6,  7],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":     [   0,   1,   2,  3,  4,  5],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":     [   1,   2,   3,  4,  5,  6],
	"default_color": [0,0,0]
      },
      { "ibits": 6,
	"blue_pos":     [   2,   3,   4,  5,  6,  7],
	"red_pos":     [   2,   3,   4,  5,  6,  7],
	"green_pos":     [   2,   3,   4,  5,  6,  7],
	"default_color": [0,0,0]
      },
      { "ibits": 12,
	"blue_pos":     [   2,   3,   4,  5,  6,  7, -1 ,-1 ,-1,-1, -1, -1 ],
	"green_pos":     [  -1 ,-1 ,-1,-1, -1, -1,  2,   3,   4,  5,  6,  7],
	"red_pos":     [  -1,-1, -1, -1, -1, -1, -1 ,-1 ,-1,-1, -1, -1 ],
	"default_color": [0,0,0]
	  }
    ]`
	all_color_names_str = `[
	    "bold: blue,green,red",
	    "bold: green,blue,red",
	    "bold: red,blue,green",
	    "dim: blue,green,red",
	    "dim: green,blue,red",
	    "dim: red,blue,green",
	    "all blue",
	    "all purple",
	    "all purple - purple center",
	    "all maroon",
	    "all orange",
	    "all lime green",
	    "all gold",
	    "all white",
	    "high resolution: blue, green"
	]`

)
// Field Names MUST start with a capital letter
type MandelColor struct {
  	Ibits int 
    Blue_pos []int
    Red_pos []int
    Green_pos []int
    Default_color []int
}

type Mandel struct {
	up_to_date bool
	size int
	cur_x int
	cur_y int
	cur_granularity int
	tiles [][]Color
	// Math
	//iterations int // Defined by Color
	threshold float64 
	span float64 
	span_one_dot float64 
	min_x,max_x,min_y,max_y float64 
	// Window
	cur_w, cur_h int
	// Colors
	all_colors []MandelColor 
	all_color_names []string
	cur_color_num int
	new_color_num int
}

type Color struct {
    red int 
    green int
    blue int
}

func (m *Mandel) CalcIterationsOneDot(c,di float64) int{
	newA :=0.0
	newBi :=0.0
	a := 0.0
	bi := 0.0
	iterations := 1 <<(m.all_colors[m.cur_color_num].Ibits)
	for i:=0; i<iterations;i++ {
	    if i == 0 {
            a=c
            bi=di
		} else {
            newA=a*a-bi*bi-c
            newBi=2.0*a*bi-di
            a=newA
            bi=newBi
            if (a>m.threshold) {
                return i
			}
		}
	}
	return 0
}

func (m *Mandel) CalcOneDot() {

	realx := m.min_x + float64(m.cur_x) * m.span_one_dot
	realy := m.min_y + float64(m.cur_y) * m.span_one_dot
	iters := m.CalcIterationsOneDot(realx,realy)

	red_color := 0 
	green_color := 0 
	blue_color := 0 
	red_adj:= 0 
	green_adj:= 0 
	blue_adj:= 0 
	for i:=0;i<m.all_colors[m.cur_color_num].Ibits;i++{
		if((iters & ( 1<<i)) != 0) {
		    red_adj = m.all_colors[m.cur_color_num].Red_pos[i]
		    green_adj = m.all_colors[m.cur_color_num].Green_pos[i]
		    blue_adj = m.all_colors[m.cur_color_num].Blue_pos[i]
			if(red_adj > 0) {
				red_color |= 1 << (red_adj)
			}
			if(green_adj > 0) {
				green_color |= 1 << (green_adj)
			}
			if(blue_adj > 0) {
				blue_color |= 1 << (blue_adj)
			}
		}
	}
	m.tiles[m.cur_x][m.cur_y].red = red_color 
	m.tiles[m.cur_x][m.cur_y].green = green_color 
	m.tiles[m.cur_x][m.cur_y].blue = blue_color 
	//m.tiles[m.cur_x][m.cur_y].red = (((iters >> 6) & 0x7) << 5) 
	//m.tiles[m.cur_x][m.cur_y].green = (((iters >> 3) & 0x7) << 5) 
	//m.tiles[m.cur_x][m.cur_y].blue = (((iters >> 0) & 0x7) << 5) 
}

func (m *Mandel) AdvanceToNextDot() {
	if(!m.up_to_date) {
		m.cur_x = (m.cur_x + m.cur_granularity) % m.size
		if (m.cur_x == 0)  {
			m.cur_y = (m.cur_y + m.cur_granularity) % m.size
			if(m.cur_y == 0) {
				if(m.cur_granularity == 1) {
					m.up_to_date = true
				} else {
					m.cur_granularity = m.cur_granularity >> 1
				}
			}
		}
	}
}

func (m *Mandel) ResetBasic() {
	// Reset Drawing 
	m.up_to_date = false
	m.cur_granularity = 16
	m.cur_x = 0
	m.cur_y = 0
}

func (m *Mandel) ResetWindow(w, h int) {
	// Check
	if((w > max_size) || (h > max_size)) {
		fmt.Println("Monitor is too big")
		panic(1)
	}
	// New Window Size
	m.cur_w = w
	m.cur_h = h
	// New Mandelbrot Size
	max_val:=0
	if(w > h) {
		max_val = w
    } else {
		max_val = h
    }
	max_mult16 := (max_val / 16) * 16
		// scale
	m.size = max_mult16
	m.span_one_dot = m.span / float64(m.size)
	// Reset Drawing 
	m.ResetBasic()
}

func (m *Mandel) DrawOneDot (px, py, w, h int) color.Color {
	use_px := 0
	use_py := 0
	if(w != m.cur_w) || (h != m.cur_h) {
		m.ResetWindow(w,h)
		use_px = px % m.size
		use_py = py % m.size
	} else {
		use_px = px
		use_py = py
	}
	//fmt.Println("px:",px,"py:",py,"w:",w,"h:",h)
	idx_x := 0
	idx_y := 0
	gran := 16
	if m.up_to_date {
	    idx_x = use_px
	    idx_y = use_py
	} else {
		if( m.cur_granularity == 16) {
			gran = 16
		} else if( use_py < m.cur_y) {
			gran = m.cur_granularity
		} else {
			gran = m.cur_granularity * 2
		}
		if(gran == 0){
			panic(1)
		}
    	idx_x = (use_px / gran) * gran
	    idx_y = (use_py / gran) * gran
	}
	ret_red := uint8(m.tiles[idx_x][idx_y].red)
	ret_green := uint8(m.tiles[idx_x][idx_y].green)
	ret_blue := uint8(m.tiles[idx_x][idx_y].blue)
	ret_color := color.RGBA{ret_red,ret_green,ret_blue, 0xff}
	return(ret_color)
}

func (m *Mandel) Status() {
	fmt.Println(m.up_to_date,m.cur_granularity,m.cur_x,m.cur_y)
}
func (m *Mandel) UpdateSome() {
	for bundle:=0;bundle<256;bundle++ {
		m.CalcOneDot()
		m.AdvanceToNextDot()
	}
//	m.Status()
}

func NewMandel() *Mandel {
//	var lcl_all_colors []MandelColor
	m := &Mandel {
		size:256,
		cur_x:0,
		cur_y:0,
		cur_granularity: 16,
		up_to_date: false,
		// Math
		span: 3.0,
		threshold:100,
		min_x: -1.0,
		max_x: 2.0,
		min_y: -1.5,
		max_y: 1.5,
		//Window
		cur_w: 256,
		cur_h: 256,
		// Color
		cur_color_num: 0,
		new_color_num: 0,
	}
	m.span_one_dot = m.span / float64(m.size)
	m.tiles = make([][]Color, max_size)       
    for i:=0;i<max_size;i++ {
    	m.tiles[i] = make([]Color, max_size)  
	}
//	lcl_all_colors_str := `[{"Ibits":9},{"Ibits":12},{"Ibits":5}]`
/*
	lcl_all_colors_str := `[
      { "ibits": 9,
	"blue_pos":    [  5,   6,   7,  -1,  -1,  -1,  -1,  -1,  -1],
	"green_pos":   [ -1,  -1,  -1,   5,   6,   7,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1,   5,   6,   7],
	"default_color": [0,0,0] 
      }
	]`
*/
	//err := json.Unmarshal([]byte(lcl_all_colors_str), &lcl_all_colors)
	err := json.Unmarshal([]byte(all_colors_str), &m.all_colors)
    if err != nil {
        fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
    }
	fmt.Println(m.all_colors)
	//fmt.Println("extra",lcl_all_colors[0].Ibits)
	//fmt.Println("extra",lcl_all_colors[0].Blue_pos)
	err = json.Unmarshal([]byte(all_color_names_str), &m.all_color_names)
    if err != nil {
        fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
    }
	fmt.Println(m.all_color_names)
	return m
}


func (m *Mandel) DoColorChange() {
	m.cur_color_num = m.new_color_num
	m.ResetBasic()
}

func (m *Mandel) ColorSettingsCallback(s string) {
	fmt.Println("Color Settings Callback:",s)
	for i:=0;i< len(m.all_color_names); i++ { 
		if(m.all_color_names[i] == s) {
			m.new_color_num = i
			break
		}
	}
}

func main() {
	
	myApp := app.New()
	myWindow := myApp.NewWindow("Mandelbrot Toy")
	myWindow.SetPadded(false)
	//myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.Resize(fyne.NewSize(256, (256+27)))
//	myCanvas := myWindow.Canvas()
	myMandel := NewMandel()

	// Control Menu Set up
	menuItemColor:= fyne.NewMenuItem("Color Settings", func() {
		fmt.Println("In Color Settings")
		var popup *widget.PopUp

		hello := widget.NewLabel("Color Hello")
		colorSelect := widget.NewSelect(myMandel.all_color_names,myMandel.ColorSettingsCallback)
		//okPopUpButton := widget.NewButton("Ok",ColorSettingsCallbackOk)
		//cancelPopUpButton := widget.NewButton("Ok",ColorSettingsCallbackCancel)
		popUpContent := container.NewVBox(
			hello,
			colorSelect,
			container.NewHBox(
				widget.NewButton("Ok",func() {
					myMandel.DoColorChange()
					popup.Hide()

				},),
				widget.NewButton("Cancel",func() {
					popup.Hide()
				},),
			),
		)
		popup = widget.NewModalPopUp(popUpContent,myWindow.Canvas())
		popup.Show()
	})
	menuItemQuit:= fyne.NewMenuItem("Quit", func() {
		//fmt.Println("In DoQuit:")
		os.Exit(0)
	})
	menuControl:= fyne.NewMenu("Control", menuItemColor, menuItemQuit);
	// About Menu Set up
	menuItemAbout := fyne.NewMenuItem("About...", func() {
		dialog.ShowInformation("About Mandlbrot Toy v0.1.0", "Author: Craig Warner \n\ngithub.com/craig-warner/mandelbrot-toy", myWindow)
	})
	menuHelp := fyne.NewMenu("Help ", menuItemAbout)
	mainMenu := fyne.NewMainMenu(menuControl, menuHelp)
	myWindow.SetMainMenu(mainMenu)

	myRaster := canvas.NewRasterWithPixels(myMandel.DrawOneDot)
//	myCanvas.SetContent(container.NewWithoutLayout(myRaster))
	myWindow.SetContent(myRaster)
//	myCanvas.Refresh(myRaster)
//	myCanvas.Refresh(myRaster)
	myRaster.Refresh()

	go func() {
		for { 
			time.Sleep(time.Nanosecond * 10000)
			if !myMandel.up_to_date {
				myMandel.UpdateSome()
			}
//			fmt.Println(doDraw,doQuit)
		    myRaster.Refresh()
//		    myCanvas.Refresh(myRaster)
		}
	}()
	myWindow.ShowAndRun()

}