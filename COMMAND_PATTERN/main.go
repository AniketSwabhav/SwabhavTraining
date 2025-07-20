package main

import "fmt"

// Receiver: Light
type Light struct {
	location string
}

func (l *Light) On() {
	fmt.Println(l.location + " light is ON")
}

func (l *Light) Off() {
	fmt.Println(l.location + " light is OFF")
}

// Command Interface
type Command interface {
	Execute()
	Undo()
}

// Concrete Command: LightOnCommand
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

// Invoker: Remote Control
type RemoteControl struct {
	slots     [7]Command
	undoStack []Command
}

func (rc *RemoteControl) SetCommand(slot int, command Command) {
	rc.slots[slot] = command
}

func (rc *RemoteControl) PressButton(slot int) {
	if rc.slots[slot] != nil {
		rc.slots[slot].Execute()
		rc.undoStack = append(rc.undoStack, rc.slots[slot])
	}
}

func (rc *RemoteControl) PressUndo() {
	if len(rc.undoStack) > 0 {
		last := rc.undoStack[len(rc.undoStack)-1]
		last.Undo()
		rc.undoStack = rc.undoStack[:len(rc.undoStack)-1]
	}
}

func main() {
	remote := &RemoteControl{}
	livingRoomLight := &Light{location: "Living Room"}

	lightOn := &LightOnCommand{light: livingRoomLight}
	remote.SetCommand(0, lightOn)

	remote.PressButton(0) // -> Living Room light is ON
	remote.PressUndo()    // -> Living Room light is OFF
}
