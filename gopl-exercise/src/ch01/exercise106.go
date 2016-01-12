// Exercise 1.6: Modify the Lissajous program to produce images in multiple colors by adding
// more values to [palette] and then displaying them by changing the third argument of
// [SetColorIndex] in some interesting way.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"

	"os"
)

import (
	"log"
	"net/http"
	"time"
	"math/big"
)

import xr "crypto/rand"

// 调色板，修改线条颜色
var palette = []color.Color{
	color.Black,
	color.RGBA{0x44, 0x33, 0x88, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xaa, 0xbb, 0xcc, 0xff},
	color.RGBA{0x11, 0x88, 0xaa, 0xff},
}

// 获取颜色的随机数
func getRandomIndex() (uint8) {

	i, _ := xr.Int(xr.Reader, big.NewInt(int64(len(palette))))

	// big.Int不能直接转换为int,需要作一次字符串转换
	// 方法一
	//	s := fmt.Sprintf("%d", i)
	//	r, _ := strconv.Atoi(s)

	// 方法二
	return uint8(i.Uint64())
}

func main() {
	// 除非使用当前时间作为伪随机数生成器的种子,否则图形是固定的,可注释进行测试
	rand.Seed(time.Now().UTC().UnixNano())

	// 传一个参数web,在浏览器中打开查看图形
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous6(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous6(os.Stdout)
}

func lissajous6(out io.Writer) {
	const (
		cycles = 5     // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5),
				size + int(y * size + 0.5),
				getRandomIndex())
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
