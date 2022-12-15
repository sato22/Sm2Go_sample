// This file was generated by a program.
// Please do not edit this file directly.
package mypackage

// package names to be imported

type stm0State uint8

const (
	AlarmOn stm0State = iota
	Countdown
	TimerSet
	Idle
)

var stm0Eod Eod

var stm0CurrentState stm0State

func init() {
	stm0Initialize()
}

func stm0Initialize() {
	stm0CurrentState = Idle
}

func Entrystm0Task() {
	stm0Task()
}

func stm0Task() {
	switch stm0CurrentState {
	case AlarmOn:
		if stm0Eod == Entry {
			AlarmOnEntry()
			stm0Eod = Do
		}
		if stm0Eod == Do {
			AlarmOnDo()
			if AlarmOffCond() {
				AlarmOffAction()
				stm0CurrentState = Idle
				stm0Eod = Exit
			}
		}
		if stm0Eod == Exit {
			AlarmOnExit()
			stm0Eod = Entry
		}
	case Countdown:
		if stm0Eod == Entry {
			CountdownEntry()
			stm0Eod = Do
		}
		if stm0Eod == Do {
			CountdownDo()
			if EndCond() {
				EndAction()
				stm0CurrentState = AlarmOn
				stm0Eod = Exit
			}
			if StopCond() {
				StopAction()
				stm0CurrentState = TimerSet
				stm0Eod = Exit
			}
		}
		if stm0Eod == Exit {
			CountdownExit()
			stm0Eod = Entry
		}
	case TimerSet:
		if stm0Eod == Entry {
			TimerSetEntry()
			stm0Eod = Do
		}
		if stm0Eod == Do {
			TimerSetDo()
			if StartCond() {
				StartAction()
				stm0CurrentState = Countdown
				stm0Eod = Exit
			}
		}
		if stm0Eod == Exit {
			TimerSetExit()
			stm0Eod = Entry
		}
	case Idle:
		if stm0Eod == Entry {
			IdleEntry()
			stm0Eod = Do
		}
		if stm0Eod == Do {
			IdleDo()
			if SetTimeCond() {
				SetTimeAction()
				stm0CurrentState = TimerSet
				stm0Eod = Exit
			}
		}
		if stm0Eod == Exit {
			IdleExit()
			stm0Eod = Entry
		}
	}
}
