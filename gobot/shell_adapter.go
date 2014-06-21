package gobot

import (
// "gobot/gobot"
)

type ShellAdapter struct {
	// gobot.Adapter
}

func (a *ShellAdapter) Send(env *Envelope, strings []string) error {
	return nil
}

func (a *ShellAdapter) Emote(env *Envelope, strings []string) error {
	return nil
}

func (a *ShellAdapter) Reply(env *Envelope, strings []string) error {
	return nil
}

func (a *ShellAdapter) Topic(env *Envelope, strings []string) error {
	return nil
}
