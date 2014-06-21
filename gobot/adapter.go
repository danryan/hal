package gobot

// type Adapter struct {
// }

type Adapter interface {
	Send(env *Envelope, strings []string) error
	Emote(env *Envelope, strings []string) error
	Reply(env *Envelope, strings []string) error
	Topic(env *Envelope, strings []string) error
}

// func (a *Adapter) Send(env *Envelope, strings []string) error {
// 	return nil
// }

// func (a *Adapter) Emote(env *Envelope, strings []string) error {
// 	return nil
// }

// func (a *Adapter) Reply(env *Envelope, strings []string) error {
// 	return nil
// }

// func (a *Adapter) Topic(env *Envelope, strings []string) error {
// 	return nil
// }
