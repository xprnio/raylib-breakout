package state

type ChallengeEvent struct {
  Target string
}

func NewChallenge(name string) ChallengeEvent {
  return ChallengeEvent{
    Target: name,
  }
}

type Player struct {
  Name string
  IsLocal bool
}
