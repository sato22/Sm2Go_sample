package timer

// Please do not edit this file.

// package name to import

const (
	debug = true
)

type OneState int

const (
	AlarmOn OneState = iota
	Countdown
	TimerSet
	Idle
)

type Eod int

const (
	Entry Eod = iota
	Do
	Exit
)

var oneeod Eod
var oneCurrentState OneState

func OneStep() {
	switch oneCurrentState {
	case AlarmOn:
		if oneeod == Entry {
			alarmonEntry()
			oneeod = Do
		}
		if oneeod == Do {
			alarmonDo()
			if AlarmOffCond() {
				oneCurrentState = Idle
				if debug {
					logger.Println("State is changed: AlarmOn to Idle")
				}
				oneeod = Exit
			}
		}
		if oneeod == Exit {
			alarmonExit()
			oneeod = Entry
		}
	case Countdown:
		if oneeod == Entry {
			countdownEntry()
			oneeod = Do
		}
		if oneeod == Do {
			countdownDo()
			if EndCond() {
				// EndCondFunc()

				oneCurrentState = AlarmOn
				if debug {
					logger.Println("State is changed: Countdown to AlarmOn")
				}
				oneeod = Exit
			}
			if StopCond() {
				// StopCondFunc()

				oneCurrentState = TimerSet
				if debug {
					logger.Println("State is changed: Countdown to TimerSet")
				}
				oneeod = Exit
			}
		}
		if oneeod == Exit {
			countdownExit()
			oneeod = Entry

		}
	case TimerSet:
		if oneeod == Entry {
			timersetEntry()
			oneeod = Do
		}
		if oneeod == Do {
			timersetDo()
			if StartCond() {
				oneCurrentState = Countdown
				if debug {
					logger.Println("State is changed: TimerSet to Countdown")
				}
				oneeod = Exit
			}
		}
		if oneeod == Exit {
			timersetExit()
			oneeod = Entry
		}
	case Idle:
		if oneeod == Entry {
			idleEntry()
			oneeod = Do
		}
		if oneeod == Do {
			idleDo()
			if SetTimeCond() {
				oneCurrentState = TimerSet
				if debug {
					logger.Println("State is changed: Idle to TimerSet")
				}
				oneeod = Exit
			}
		}
		if oneeod == Exit {
			idleExit()
			oneeod = Entry
		}
	}
}

func init() {
	oneCurrentState = Idle
	oneeod = Entry
}
