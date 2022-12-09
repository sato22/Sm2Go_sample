package timer

import "fmt"

const (
	countSecond = 18
)

type DebugLogger interface {
	Println(string)
}

var logger DebugLogger

func ConfigureLog(p DebugLogger) {
	logger = p
}

type Switch interface {
	Get() bool
}

var buttonL Switch
var buttonM Switch
var buttonR Switch

func ConfigureLeftButton(s Switch) {
	buttonL = s
}

func ConfigureMiddleButton(s Switch) {
	buttonM = s
}

func ConfigureRightButton(s Switch) {
	buttonR = s
}

// interface Monitor
// ディスプレイに文字が表示され、表示が更新されることを確認するためのインターフェース
type Monitor interface {
	PrintVal(string)
}

var display Monitor

func ConfigureMonitor(m Monitor) {
	display = m
}

// interface Alarm
// アラームがなることを確認するためのインターフェース
type Alarm interface {
	Beep()
	Mute()
}

var alarm Alarm

func ConfigureAlarm(a Alarm) {
	alarm = a
}

var mm int = 0
var ss int = 0

var inputEnabled bool = true

var count int = 0

func setMinute() {
	// Called in [State: timer, Action: Do]
	mm++
	if mm == 60 {
		mm = 0
	}
	display.PrintVal(fmt.Sprintf("%02d : %02d", mm, ss))
}

func setSecond() {
	// Called in [State: timer, Action: Do]
	ss++
	if ss == 60 {
		ss = 0
	}
	display.PrintVal(fmt.Sprintf("%02d : %02d", mm, ss))
}

func countDown() {
	ss--
	if ss == -1 {
		if mm == 0 {
			ss = 0
		} else {
			ss = 59
			mm--
		}
	}

	if ss%2 == 1 {
		display.PrintVal(fmt.Sprintf("%02d : %02d", mm, ss))
	} else {
		display.PrintVal(fmt.Sprintf("%02d   %02d", mm, ss))
	}
}

func initialize() {
	display.PrintVal("00 : 00")
}

func alarmonEntry() {
	alarm.Beep()
}

func alarmonDo() {
	// nothing to do
}

func alarmonExit() {
	alarm.Mute()
}

func countdownEntry() {
	// nothing to do
}

func countdownDo() {
	count++

	if count == countSecond {
		countDown()
		count = 0
	}
}

func countdownExit() {
	// nothing to do
}

func timersetEntry() {
	// nothing to do
}

func timersetDo() {
	if !buttonL.Get() {
		setMinute()
	} else if !buttonM.Get() {
		setSecond()
	}
}

func timersetExit() {
	// nothing to do
}

func idleEntry() {
	initialize()
}

func idleDo() {
	// nothing to do
}

func idleExit() {
	// nothing to do
}

func AlarmOffCond() bool {
	return !buttonR.Get()
}

func EndCond() bool {
	return (ss == 0 && mm == 0)
}

func StopCond() bool {
	return !buttonR.Get()
}

func StartCond() bool {
	return !buttonR.Get()
}

func SetTimeCond() bool {
	return (!buttonL.Get() || !buttonM.Get())
}
