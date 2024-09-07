package main

import (
	"fmt"
	"image/color"
	"os"
	"time"

	//	"github.com/hjson/hjson-go/v4"
	"encoding/json"

	"math"
	"math/rand"

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
      { "Ibits": 12,
	"blue_pos":    [  4,   5,   6,  7,  -1,  -1,  -1,  -1,  -1,  -1,  -1,  -1],
	"green_pos":   [ -1,  -1,  -1,   4,  5,   6,   7,  -1,  -1,  -1,  -1,  -1],
	"red_pos":     [ -1,  -1,  -1,  -1,  -1,  -1, -1,  -1,   4,   5,   6,   7],
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

type tappableRaster struct {
	fyne.CanvasObject
	OnTapped func()
}

func NewTappableRaster(raster fyne.CanvasObject, onTapped func()) *tappableRaster {
	return &tappableRaster{CanvasObject: raster, OnTapped: onTapped}
}

func (t *tappableRaster) Tapped(ev *fyne.PointEvent) {
	fmt.Println("x,y:", ev.Position.X, ev.Position.Y)
	t.OnTapped()
}

// func DoRasterTap(ev *fyne.PointEvent) {
func DoRasterTap() {
	fmt.Println("Tapped")
}

//func (t *tappableRaster) pixelColor(x,y,w,h int) color.Color {
//	fmt.Println( "x,y",x,y,w,h)
//	return(color.Black)
//}

// Field Names MUST start with a capital letter
type MandelColor struct {
	Ibits         int
	Blue_pos      []int
	Red_pos       []int
	Green_pos     []int
	Default_color []uint8
}

type Mandel struct {
	up_to_date      bool
	size            int
	cur_x           int
	cur_y           int
	cur_granularity int
	tiles           [][]Color
	// Math
	//iterations int // Defined by Color
	threshold    float64
	span         float64
	span_one_dot float64
	min_x, min_y float64
	// Window
	cur_w, cur_h int
	// Colors
	all_colors      []MandelColor
	all_color_names []string
	cur_color_num   int
	new_color_num   int
	// Zoom Tap
	cur_zoom float64
	new_zoom float64
	// Roam
	roam_state           RoamState
	cur_roam_speed       int // 1 to 100 (fast)
	new_roam_speed       int
	cur_draw_speed       int
	new_draw_speed       int
	cur_pan_total_steps  int
	cur_zoom_total_steps int
	roam_tgt_x           float64
	roam_tgt_y           float64
	roam_tgt_span_adj    float64 // 0.1-0.99
	roam_step            int
}

type Color struct {
	red   uint8
	green uint8
	blue  uint8
}

type RoamState int

const (
	PanTo RoamState = iota + 1
	ZoomIn
	ZoomOut
	PanFrom
)

type Point struct {
	x float64
	y float64
}

func NewColor(r, g, b uint8) Color {
	c := Color{
		red:   r,
		green: g,
		blue:  b,
	}
	return c
}

func NewPoint(set_x, set_y float64) Point {
	p := Point{
		x: set_x,
		y: set_y,
	}
	return p
}

func (m *Mandel) CalcIterationsOneXY(c, di float64) int {
	newA := 0.0
	newBi := 0.0
	a := 0.0
	bi := 0.0
	iterations := 1 << (m.all_colors[m.cur_color_num].Ibits)
	for i := 0; i < iterations; i++ {
		if i == 0 {
			a = c
			bi = di
		} else {
			newA = a*a - bi*bi - c
			newBi = 2.0*a*bi - di
			a = newA
			bi = newBi
			if a > m.threshold {
				return i
			}
		}
	}
	return 0
}

func (m *Mandel) CalcIterationsOnePoint(p Point) int {
	iters := m.CalcIterationsOneXY(p.x, p.y)
	return iters
}

func (m *Mandel) CalcOnePointRGB(p Point) (red_color uint8, green_color uint8, blue_color uint8) {

	iters := m.CalcIterationsOneXY(p.x, p.y)

	red_color = 0
	green_color = 0
	blue_color = 0
	red_adj := 0
	green_adj := 0
	blue_adj := 0
	if iters == 0 {
		red_color = uint8(m.all_colors[m.cur_color_num].Default_color[0])
		green_color = uint8(m.all_colors[m.cur_color_num].Default_color[1])
		blue_color = uint8(m.all_colors[m.cur_color_num].Default_color[2])
	} else {
		for i := 0; i < m.all_colors[m.cur_color_num].Ibits; i++ {
			if (iters & (1 << i)) != 0 {
				red_adj = m.all_colors[m.cur_color_num].Red_pos[i]
				green_adj = m.all_colors[m.cur_color_num].Green_pos[i]
				blue_adj = m.all_colors[m.cur_color_num].Blue_pos[i]
				if red_adj > 0 {
					red_color |= 1 << (red_adj)
				}
				if green_adj > 0 {
					green_color |= 1 << (green_adj)
				}
				if blue_adj > 0 {
					blue_color |= 1 << (blue_adj)
				}
			}
		}
	}
	return
}

func (m *Mandel) CalcOnePointColor(p Point) (c Color) {
	red, green, blue := m.CalcOnePointRGB(p)
	c = NewColor(red, green, blue)
	return
}

func (m *Mandel) CalcOneDot() {
	var p Point

	realx := m.min_x + float64(m.cur_x)*m.span_one_dot
	realy := m.min_y + m.span - float64(m.cur_y)*m.span_one_dot

	p = NewPoint(realx, realy)

	color := m.CalcOnePointColor(p)

	m.tiles[m.cur_x][m.cur_y].red = color.red
	m.tiles[m.cur_x][m.cur_y].green = color.green
	m.tiles[m.cur_x][m.cur_y].blue = color.blue
}

func (m *Mandel) AdvanceToNextDot() {
	if !m.up_to_date {
		m.cur_x = (m.cur_x + m.cur_granularity) % m.size
		if m.cur_x == 0 {
			m.cur_y = (m.cur_y + m.cur_granularity) % m.size
			if m.cur_y == 0 {
				if m.cur_granularity == 1 {
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
	m.cur_granularity = 64
	m.cur_x = 0
	m.cur_y = 0
}

func (m *Mandel) ResetSpan() {
	m.span = 3.0
	m.min_x = -1.0
	//	m.max_x= 2.0
	m.min_y = -1.5
	//	m.max_y= 1.5
	m.span_one_dot = m.span / float64(m.size)
}

func (m *Mandel) ResetWindow(w, h int) {
	// Check
	if (w > max_size) || (h > max_size) {
		fmt.Println("Monitor is too big")
		panic(1)
	}
	// New Window Size
	m.cur_w = w
	m.cur_h = h
	// New Mandelbrot Size
	max_val := 0
	if w > h {
		max_val = w
	} else {
		max_val = h
	}
	max_mult64 := (max_val / 64) * 64
	// scale
	m.size = max_mult64
	m.span_one_dot = m.span / float64(m.size)
	// Reset Drawing
	m.ResetBasic()
}

func (m *Mandel) DrawOneDot(px, py, w, h int) color.Color {
	use_px := 0
	use_py := 0
	if (w != m.cur_w) || (h != m.cur_h) {
		m.ResetWindow(w, h)
		use_px = px % m.size
		use_py = py % m.size
	} else {
		use_px = px
		use_py = py
	}
	//fmt.Println("px:",px,"py:",py,"w:",w,"h:",h)
	idx_x := 0
	idx_y := 0
	gran := 64
	if m.up_to_date {
		idx_x = use_px
		idx_y = use_py
	} else {
		if m.cur_granularity == 64 {
			gran = 64
		} else if use_py < m.cur_y {
			gran = m.cur_granularity
		} else {
			gran = m.cur_granularity * 2
		}
		if gran == 0 {
			panic(1)
		}
		idx_x = (use_px / gran) * gran
		idx_y = (use_py / gran) * gran
	}
	ret_red := uint8(m.tiles[idx_x][idx_y].red)
	ret_green := uint8(m.tiles[idx_x][idx_y].green)
	ret_blue := uint8(m.tiles[idx_x][idx_y].blue)
	ret_color := color.RGBA{ret_red, ret_green, ret_blue, 0xff}
	return (ret_color)
}

func (m *Mandel) Status() {
	fmt.Println(m.up_to_date, m.cur_granularity, m.cur_x, m.cur_y)
}

func (m *Mandel) CalcBundleSize() int {
	bsize := 0
	if m.cur_granularity == 64 {
		bsize = 4
	} else if m.cur_granularity == 32 {
		bsize = 16
	} else if m.cur_granularity == 16 {
		bsize = 64
	} else if m.cur_granularity == 8 {
		bsize = 256
	} else if m.cur_granularity == 4 {
		bsize = 1024
	} else if m.cur_granularity == 2 {
		bsize = 4096
	} else if m.cur_granularity == 1 {
		bsize = 4096 * 4
	}
	return bsize
}

func (m *Mandel) UpdateSome() {

	// Update one Dot and advance
	bsize := m.CalcBundleSize()
	for b := 0; b < bsize; b++ {
		m.CalcOneDot()
		m.AdvanceToNextDot()
	}
	// Stall longer for courser granularities
	for d := 0; d < (101 - m.cur_draw_speed); d++ {
		time.Sleep(time.Nanosecond * 100000)
	}
}

func (m *Mandel) RoamTgtScreenTwo(x, y float64) bool {

	new_span := 3.0
	for i := 0; i < m.cur_zoom_total_steps; i++ {
		new_span = new_span * m.roam_tgt_span_adj
	}
	half_new_span := new_span / 2.0

	upper_left_pnt := NewPoint((x - half_new_span), (y + half_new_span))
	upper_right_pnt := NewPoint((x + half_new_span), (y + half_new_span))
	lower_left_pnt := NewPoint((x - half_new_span), (y - half_new_span))
	lower_right_pnt := NewPoint((x + half_new_span), (y - half_new_span))

	upper_left_iters := m.CalcIterationsOnePoint(upper_left_pnt)
	upper_right_iters := m.CalcIterationsOnePoint(upper_right_pnt)
	lower_left_iters := m.CalcIterationsOnePoint(lower_left_pnt)
	lower_right_iters := m.CalcIterationsOnePoint(lower_right_pnt)

	f64_upper_left_iters := float64(upper_left_iters)
	f64_upper_right_iters := float64(upper_right_iters)
	f64_lower_left_iters := float64(lower_left_iters)
	f64_lower_right_iters := float64(lower_right_iters)

	/*
		fmt.Println(upper_left_pnt)
		fmt.Println(upper_left_iters)
		fmt.Println(upper_right_pnt)
		fmt.Println(upper_right_iters)
		fmt.Println(lower_left_pnt)
		fmt.Println(lower_left_iters)
		fmt.Println(lower_right_pnt)
		fmt.Println(lower_right_iters)
	*/

	same_cnt := 0
	if upper_left_iters == upper_right_iters {
		same_cnt++
	}
	if upper_left_iters == lower_left_iters {
		same_cnt++
	}
	if upper_left_iters == lower_right_iters {
		same_cnt++
	}
	if upper_right_iters == lower_right_iters {
		same_cnt++
	}
	if upper_right_iters == lower_left_iters {
		same_cnt++
	}
	if lower_left_iters == lower_right_iters {
		same_cnt++
	}
	//fmt.Println("Screen 2: ", same_cnt)

	iterbits := m.all_colors[m.cur_color_num].Ibits
	max_iters := (1 << iterbits)
	f64_max_iters := float64(max_iters)
	good_pnt := 0.7
	if (f64_upper_left_iters / f64_max_iters) > good_pnt {
		return true
	} else if (f64_upper_right_iters / f64_max_iters) > good_pnt {
		return true
	} else if (f64_lower_left_iters / f64_max_iters) > good_pnt {
		return true
	} else if (f64_lower_right_iters / f64_max_iters) > good_pnt {
		return true
	} else if same_cnt > 3 {
		return false
	} else {
		return true
	}
}

func RoamCalcDistance(x, y float64) float64 {
	distance := math.Sqrt(x*x + y*y)
	return (distance)
}

// Must not be in center
func (m *Mandel) RoamTgtScreenOne(x, y float64) bool {
	distance := RoamCalcDistance(x, y)
	if distance < 1.5 {
		return false
	} else {
		if distance > 2.5 {
			return false
		} else {
			return true
		}
	}
}

func (m *Mandel) RoamGenNewTgt() {
	new_x := 0.0
	new_y := 0.0
	found_good_tgt := false
	for found_good_tgt == false {
		new_x = float64(rand.Intn(100))/100.0*3 - 1
		new_y = float64(rand.Intn(100))/100.0*3 - 1.5
		if m.RoamTgtScreenOne(new_x, new_y) {
			if m.RoamTgtScreenTwo(new_x, new_y) {
				found_good_tgt = true
			}
		}
	}
	m.roam_tgt_x = new_x
	m.roam_tgt_y = new_y
}

func (m *Mandel) RoamDelay() {
	for delays := 0; delays < (101 - m.cur_roam_speed); delays++ {
		time.Sleep(time.Nanosecond * 100000000)
	}
}

func (m *Mandel) RoamAdjustSetMinXMinY(imageCenter Point) {
	m.min_x = imageCenter.x - (m.span / 2.0)
	m.min_y = imageCenter.y - (m.span / 2.0)
}

func (m *Mandel) RoamAdjustPanTo() {
	percentPanned := float64(m.roam_step) / float64(m.cur_pan_total_steps)
	new_x := 0.5 + (m.roam_tgt_x-0.5)*percentPanned
	new_y := m.roam_tgt_y * percentPanned
	imageCenter := NewPoint(new_x, new_y)
	m.RoamAdjustSetMinXMinY(imageCenter)
}

func (m *Mandel) RoamAdjustPanFrom() {
	percentPanned := float64(m.roam_step) / float64(m.cur_pan_total_steps)
	new_x := m.roam_tgt_x - (m.roam_tgt_x-0.5)*percentPanned
	new_y := m.roam_tgt_y - m.roam_tgt_y*percentPanned
	imageCenter := NewPoint(new_x, new_y)
	m.RoamAdjustSetMinXMinY(imageCenter)
}

func (m *Mandel) RoamAdjustZoomIn() {
	// Reduce span
	m.span = m.span * m.roam_tgt_span_adj
	m.span_one_dot = m.span / float64(m.size)
	// Set upper left point
	imageCenter := NewPoint(m.roam_tgt_x, m.roam_tgt_y)
	m.RoamAdjustSetMinXMinY(imageCenter)
}

func (m *Mandel) RoamAdjustZoomOut() {
	// Increase span
	new_span := 3.0
	for i := 0; i < (m.cur_zoom_total_steps - m.roam_step); i++ {
		new_span = new_span * m.roam_tgt_span_adj
	}
	m.span = new_span
	m.span_one_dot = m.span / float64(m.size)
	// Set upper left point
	imageCenter := NewPoint(m.roam_tgt_x, m.roam_tgt_y)
	m.RoamAdjustSetMinXMinY(imageCenter)
}

func (m *Mandel) FrcRedraw() {
	m.up_to_date = false
	m.cur_granularity = 64
}

func (m *Mandel) RoamAdjust() {
	m.FrcRedraw()

	next_phase_pan := m.cur_pan_total_steps + 1
	next_phase_zoom := m.cur_zoom_total_steps + 1
	if m.roam_state == PanTo {
		if m.roam_step == next_phase_pan {
			m.roam_state = ZoomIn
			m.roam_step = 0
		} else {
			m.RoamAdjustPanTo()
		}
	} else if m.roam_state == PanFrom {
		if m.roam_step == next_phase_pan {
			m.roam_step = 0
			m.ResetSpan()
			m.RoamGenNewTgt()
			m.roam_state = PanTo
		} else {
			m.RoamAdjustPanFrom()
		}
	} else if m.roam_state == ZoomIn {
		if m.roam_step == next_phase_zoom {
			m.roam_state = ZoomOut
			m.roam_step = 0
		} else {
			m.RoamAdjustZoomIn()
		}
	} else if m.roam_state == ZoomOut {
		if m.roam_step == next_phase_zoom {
			m.roam_state = PanFrom
			m.roam_step = 0
		} else {
			m.RoamAdjustZoomOut()
		}
	}
	// advance step
	m.roam_step++
	// Report Roam
	//fmt.Println(m.roam_state, " : Step : ", m.roam_step)
	//fmt.Println("mx:", m.min_x, " : My: ", m.min_y, " : Span:", m.span, " : tx:", m.roam_tgt_x, ": ty:", m.roam_tgt_y)
}

func NewMandel() Mandel {
	//	var lcl_all_colors []MandelColor
	m := Mandel{
		size:            256,
		cur_x:           0,
		cur_y:           0,
		cur_granularity: 64,
		up_to_date:      false,
		// Math
		span:      3.0,
		threshold: 1000.0,
		min_x:     -1.0,
		//		max_x: 2.0,
		min_y: -1.5,
		//	max_y: 1.5,
		//Window
		cur_w: 256,
		cur_h: 256,
		// Color
		cur_color_num: 0,
		new_color_num: 0,
		// Zoom
		cur_zoom: 2.0,
		new_zoom: 2.0,
		// Roam
		cur_roam_speed:       50,
		new_roam_speed:       50,
		cur_draw_speed:       70, // 100 is fast
		new_draw_speed:       70, // 100
		cur_pan_total_steps:  20,
		cur_zoom_total_steps: 150,
		roam_tgt_x:           1.5,
		roam_tgt_y:           0.0,
		roam_tgt_span_adj:    0.95,
		roam_state:           PanTo,
		roam_step:            0,
	}
	m.span_one_dot = m.span / float64(m.size)
	m.tiles = make([][]Color, max_size)
	for i := 0; i < max_size; i++ {
		m.tiles[i] = make([]Color, max_size)
	}
	err := json.Unmarshal([]byte(all_colors_str), &m.all_colors)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	//fmt.Println(m.all_colors)
	//fmt.Println("extra",lcl_all_colors[0].Ibits)
	//fmt.Println("extra",lcl_all_colors[0].Blue_pos)
	err = json.Unmarshal([]byte(all_color_names_str), &m.all_color_names)
	if err != nil {
		fmt.Printf("Unable to marshal JSON due to %s", err)
		panic(1)
	}
	//fmt.Println(m.all_color_names)
	return m
}

func (m *Mandel) DoColorChange() {
	m.cur_color_num = m.new_color_num
	m.ResetBasic()
}

func (m *Mandel) DoSpeedChange() {
	m.cur_roam_speed = m.new_roam_speed
	m.cur_draw_speed = m.new_draw_speed
}

func (m *Mandel) ColorSettingsCallback(s string) {
	fmt.Println("Color Settings Callback:", s)
	for i := 0; i < len(m.all_color_names); i++ {
		if m.all_color_names[i] == s {
			m.new_color_num = i
			break
		}
	}
}

func (m *Mandel) DoZoomChange(reset bool) {
	m.cur_zoom = m.new_zoom
	if reset {
		m.ResetSpan()
		m.ResetBasic()
	}
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Mandelbrot Toy")
	myWindow.SetPadded(false)
	//myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.Resize(fyne.NewSize(256, (256 + 27)))
	//	myCanvas := myWindow.Canvas()
	myMandel := NewMandel()
	myMandel.RoamGenNewTgt()

	// Control Menu Set up
	menuItemColor := fyne.NewMenuItem("Color Settings", func() {
		fmt.Println("In Color Settings")
		var popup *widget.PopUp

		color_hello := widget.NewLabel("Coloring Scheme")
		colorSelect := widget.NewSelect(myMandel.all_color_names, myMandel.ColorSettingsCallback)
		//okPopUpButton := widget.NewButton("Ok",ColorSettingsCallbackOk)
		//cancelPopUpButton := widget.NewButton("Ok",ColorSettingsCallbackCancel)
		popUpContent := container.NewVBox(
			color_hello,
			colorSelect,
			container.NewHBox(
				widget.NewButton("Ok", func() {
					myMandel.DoColorChange()
					popup.Hide()

				}),
				widget.NewButton("Cancel", func() {
					popup.Hide()
				}),
			),
		)
		popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
		popup.Show()
	})
	menuItemSpeed := fyne.NewMenuItem("Speed Settings", func() {
		var popup *widget.PopUp
		roaming_speed_hello := widget.NewLabel("Roaming Speed")
		roaming_slider := widget.NewSlider(1.0, 100.0)
		roaming_slider.SetValue(float64(myMandel.cur_roam_speed))
		roaming_slider.OnChanged = func(v float64) {
			myMandel.new_roam_speed = int(v)
		}
		draw_speed_hello := widget.NewLabel("Drawing Speed")
		draw_slider := widget.NewSlider(1.0, 100.0)
		draw_slider.SetValue(float64(float64(myMandel.cur_draw_speed)))
		draw_slider.OnChanged = func(v float64) {
			myMandel.new_draw_speed = int(v)
		}
		popUpContent := container.NewVBox(
			roaming_speed_hello,
			roaming_slider,
			draw_speed_hello,
			draw_slider,
			container.NewHBox(
				widget.NewButton("Ok", func() {
					myMandel.DoSpeedChange()
					popup.Hide()
				}),
				widget.NewButton("Cancel", func() {
					popup.Hide()
				}),
			),
		)
		popup = widget.NewModalPopUp(popUpContent, myWindow.Canvas())
		popup.Show()
	})
	/*
			menuItemZoom:= fyne.NewMenuItem("Zoom Settings", func() {
				fmt.Println("In Zoome Settings")
				var popup *widget.PopUp

				zoom_hello := widget.NewLabel("Zoom Out / Zoom In")
				zoom_slider:= widget.NewSlider(-10.0,10.0)
				zoom_slider.SetValue(myMandel.cur_zoom)
				zoom_slider.OnChanged = func(v float64) {
		            fmt.Println("Slider changed:", v)
					myMandel.new_zoom = v
		    	}
				popUpContent := container.NewVBox(
					zoom_hello,
					zoom_slider,
					container.NewHBox(
						widget.NewButton("Ok",func() {
							myMandel.DoZoomChange(false)
							popup.Hide()
						},),
						widget.NewButton("Cancel",func() {
							popup.Hide()
						},),
						widget.NewButton("Reset",func() {
							myMandel.new_zoom = 2.0
							myMandel.DoZoomChange(true)
							popup.Hide()
						},),
					),
				)
				popup = widget.NewModalPopUp(popUpContent,myWindow.Canvas())
				popup.Show()
			})
	*/
	menuItemQuit := fyne.NewMenuItem("Quit", func() {
		//fmt.Println("In DoQuit:")
		os.Exit(0)
	})
	//	menuControl:= fyne.NewMenu("Control", menuItemColor, menuItemZoom, menuItemQuit);
	menuControl := fyne.NewMenu("Control", menuItemColor, menuItemSpeed, menuItemQuit)
	// About Menu Set up
	menuItemAbout := fyne.NewMenuItem("About...", func() {
		dialog.ShowInformation("About Mandlbrot Toy v1.0.0", "Author: Craig Warner \n\ngithub.com/craig-warner/mandelbrot-toy", myWindow)
	})
	menuHelp := fyne.NewMenu("Help ", menuItemAbout)
	mainMenu := fyne.NewMainMenu(menuControl, menuHelp)
	myWindow.SetMainMenu(mainMenu)

	myRaster := canvas.NewRasterWithPixels(myMandel.DrawOneDot)
	//	myTappableRaster := NewTappableRaster(myRaster,DoRasterTap)
	//	myCanvas.SetContent(container.NewWithoutLayout(myRaster))
	//	myWindow.SetContent(myTappableRaster)
	myWindow.SetContent(myRaster)
	//	myCanvas.Refresh(myRaster)
	//	myCanvas.Refresh(myRaster)
	myRaster.Refresh()

	go func() {
		for {
			//fmt.Println(myMandel)
			if !myMandel.up_to_date {
				myMandel.UpdateSome()
				myRaster.Refresh()
			} else {
				//fmt.Println("Loop:", myMandel.roam_tgt_x, myMandel.roam_tgt_y)
				myMandel.RoamDelay()
				myMandel.RoamAdjust()
			}
		}
	}()
	myWindow.ShowAndRun()

}
