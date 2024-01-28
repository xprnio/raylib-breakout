package events

type ExitEvent struct {}

func NewExitEvent() ExitEvent {
  return ExitEvent{}
}
