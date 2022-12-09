package timer

// This is a test file for testing state transitions

import (
	"Sm2go_sample/kitchenTimer_m5stack/sm2go"
	"fmt"
	"log"
	"testing"
	"time"
)

type DebugStruct struct{}

var logTest = DebugStruct{}

func (l DebugStruct) Println(debstr string) {
	log.Println(debstr)
}

// Button
type Button struct {
	name    string
	release bool // true → release状態，false → push状態
}

func (b *Button) Push() {
	b.release = false
	log.Println(b.name, "Push")
}

func (b *Button) Release() {
	b.release = true
	log.Println(b.name, "Release")
}

func (b Button) Get() bool {
	return b.release
}

// Screen
type Screen struct{}

func (s Screen) PrintVal(str string) {
	logger.Println(fmt.Sprintf("screen display: %s", str))
}

// Alarm
type Buzzer struct{}

func (b Buzzer) Beep() {
	logger.Println("alarm beep!")
}

func (b Buzzer) Mute() {
	logger.Println("alarm mute")
}

// define struct
var leftButton = &Button{"leftButton", true}
var middleButton = &Button{"middleButton", true}
var rightButton = &Button{"rightButton", true}

var screen = &Screen{}

var buzzer = &Buzzer{}

func TestDevice01(t *testing.T) {
	env := sm2go.NewTestEnv() // TestEnv構造体

	ConfigureLeftButton(leftButton)
	ConfigureMiddleButton(middleButton)
	ConfigureRightButton(rightButton)

	ConfigureMonitor(screen)

	ConfigureAlarm(buzzer)

	ConfigureLog(logTest)

	// goroutine(base.go Task())
	env.Add(sm2go.Continue, func() {
		for {
			time.Sleep(10 * time.Millisecond)
			OneStep()
		}
	},
	)

	// goroutine(user operation)
	env.Add(sm2go.Done, func() {
		logger.Println("----------------leftButtonPush-------------")
		env.Sleep(1 * time.Second)
		leftButton.Push()
		env.Sleep(50 * time.Millisecond)
		leftButton.Release()

		logger.Println("----------------middleButtonPush-------------")
		env.Sleep(1 * time.Second)
		middleButton.Push()
		env.Sleep(50 * time.Millisecond)
		middleButton.Release()

		logger.Println("----------------rightButtonPush-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(50 * time.Millisecond)
		rightButton.Release()
		env.Sleep(15 * time.Second)
	},
	)

	env.Set(1)
	env.Go()
}

func TestDevice02(t *testing.T) {
	env := sm2go.NewTestEnv() // TestEnv構造体

	ConfigureLeftButton(leftButton)
	ConfigureMiddleButton(middleButton)
	ConfigureRightButton(rightButton)

	ConfigureMonitor(screen)

	ConfigureAlarm(buzzer)

	ConfigureLog(logTest)

	// goroutine(base.go Task())
	env.Add(sm2go.Continue, func() {
		for {
			env.Sleep(10 * time.Millisecond)
			OneStep()
		}
	},
	)

	// goroutine(user operation)
	// alarm ON/OFF
	env.Add(sm2go.Done, func() {
		logger.Println("----------------middleButtonPush-------------")
		env.Sleep(1 * time.Second)
		middleButton.Push()
		env.Sleep(100 * time.Millisecond)
		middleButton.Release()

		logger.Println("----------------rightButtonPush-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(50 * time.Millisecond)
		rightButton.Release()
		env.Sleep(10 * time.Second)

		logger.Println("----------------rightButtonPush to mute-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(50 * time.Millisecond)
		rightButton.Release()

		logger.Println("----------------middleButtonPush-------------")
		env.Sleep(1 * time.Second)
		middleButton.Push()
		env.Sleep(100 * time.Millisecond)
		middleButton.Release()

		logger.Println("----------------rightButtonPush-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(50 * time.Millisecond)
		rightButton.Release()
		env.Sleep(10 * time.Second)
	},
	)

	env.Set(1)
	env.Go()
}

func TestDevice03(t *testing.T) {
	env := sm2go.NewTestEnv() // TestEnv構造体

	ConfigureLeftButton(leftButton)
	ConfigureMiddleButton(middleButton)
	ConfigureRightButton(rightButton)

	ConfigureMonitor(screen)

	ConfigureAlarm(buzzer)

	ConfigureLog(logTest)

	// goroutine(base.go Task())
	env.Add(sm2go.Continue, func() {
		for {
			time.Sleep(10 * time.Millisecond)
			OneStep()

		}
	},
	)

	// goroutine(user operation)
	env.Add(sm2go.Done, func() {
		logger.Println("----------------leftButtonPush-------------")
		env.Sleep(1 * time.Second)
		leftButton.Push()
		env.Sleep(50 * time.Millisecond)
		leftButton.Release()

		logger.Println("----------------middleButtonPush-------------")
		env.Sleep(1 * time.Second)
		middleButton.Push()
		env.Sleep(50 * time.Millisecond)
		middleButton.Release()

		logger.Println("----------------rightButtonPush to Start-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(100 * time.Millisecond)
		rightButton.Release()
		env.Sleep(5 * time.Second)

		logger.Println("----------------rightButtonPush to Stop-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(100 * time.Millisecond)
		rightButton.Release()
		env.Sleep(5 * time.Second)

		logger.Println("----------------middleButtonPush-------------")
		env.Sleep(1 * time.Second)
		middleButton.Push()
		env.Sleep(100 * time.Millisecond)
		middleButton.Release()

		logger.Println("----------------rightButtonPush to Start-------------")
		env.Sleep(1 * time.Second)
		rightButton.Push()
		env.Sleep(200 * time.Millisecond)
		rightButton.Release()
		env.Sleep(5 * time.Second)

	},
	)

	env.Set(1)
	env.Go()
}
