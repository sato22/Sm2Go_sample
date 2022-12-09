// main file

package main

import (
	"Sm2go_sample/kitchenTimer_m5stack/timer"

	"tinygo.org/x/drivers/buzzer"
	"tinygo.org/x/drivers/examples/ili9341/initdisplay"
	"tinygo.org/x/drivers/ili9341"

	"tinygo.org/x/tinyfont"
	// "tinygo.org/x/tinyfont/examples/initdisplay"
	"tinygo.org/x/tinyfont/freesans"

	"image/color"
	"machine"
	"time"
)

var (
	black = color.RGBA{0x00, 0x00, 0x00, 0xFF}
	white = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	// red   = color.RGBA{255, 0, 0, 255}
	// blue  = color.RGBA{0, 0, 255, 255}
	// green = color.RGBA{0, 255, 0, 255}
)

// ターミナルエミュレータ用にPrintln関数を置き換え
type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	println(debstr)
}

// ブザーを鳴らす際の音階の設定
type note struct {
	tone     float64
	duration float64
}

var notes []note = []note{
	{buzzer.C3, buzzer.Whole},
}

// main文用にBeep()関数, Mute()関数を定義
type DebugBuzzer struct{}

var bzr = DebugBuzzer{}

var (
	buzz buzzer.Device
)

func (b DebugBuzzer) Beep() {
	for _, n := range notes {
		buzz.Tone(n.tone, n.duration)
		time.Sleep(10 * time.Millisecond)
	}
}

func (b DebugBuzzer) Mute() {
	buzz.Toggle()
}

// main文用にPrintVal()関数(hoge_impl.goにて定義)を再定義
var (
	display *ili9341.Device
)

type DebugMonitor struct{}

var moni = DebugMonitor{}

func (m DebugMonitor) PrintVal(str string) {
	// 一度文字の上に白い画像を配置することで文字を消去してから再び表示するため、チラつく
	display.FillRectangle(0, 0, 320, 240, white)
	tinyfont.WriteLine(display, &freesans.Bold18pt7b, 0, 135, str, black)
}

func main() {
	// ディスプレイの初期設定
	display = initdisplay.InitDisplay()
	display.FillScreen(white)

	// アラームの初期設定
	bzrPin := machine.SPEAKER_PIN
	bzrPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	buzz = buzzer.New(bzrPin)

	// 左、真ん中、右のボタンの初期設定
	btnL := machine.BUTTON_A
	btnL.Configure(machine.PinConfig{Mode: machine.PinInput})

	btnM := machine.BUTTON_B
	btnM.Configure(machine.PinConfig{Mode: machine.PinInput})

	btnR := machine.BUTTON_C
	btnR.Configure(machine.PinConfig{Mode: machine.PinInput})

	// インターフェースで丸め込む
	timer.ConfigureLeftButton(btnL)
	timer.ConfigureMiddleButton(btnM)
	timer.ConfigureRightButton(btnR)

	timer.ConfigureMonitor(moni)

	timer.ConfigureAlarm(bzr)

	timer.ConfigureLog(logTest)

	for {
		timer.OneStep()
		time.Sleep(time.Millisecond * 50)
		// 5ms (Step()) + 50ms (time.Sleep()) = 一周当たり55msかかるとする
	}
}
