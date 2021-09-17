package main

import "fmt"

//button кнопка
type button struct {
    command command
}

//command команда
type command interface {
    execute()
}

//device команды устройства
type device interface {
    on()
    off()
}

//offCommand отключение
type offCommand struct {
    device device
}

//onCommand включение
type onCommand struct {
    device device
}

//tv статус запуска ТВ
type tv struct {
    isRunning bool
}


func main() {
    tv := &tv{}

    onCommand := &onCommand{
        device: tv,
    }

    offCommand := &offCommand{
        device: tv,
    }

    onButton := &button{
        command: onCommand,
    }
    onButton.press()

    offButton := &button{
        command: offCommand,
    }
    offButton.press()
}

//press нажать кнопку
func (b *button) press() {
    b.command.execute()
}

//execute выполнить команду отключения
func (c *offCommand) execute() {
    c.device.off()
}

//execute выполнить команду включения
func (c *onCommand) execute() {
    c.device.on()
}

//on метод включения
func (t *tv) on() {
    t.isRunning = true
    fmt.Println("Turning tv on")
}

//off метод выключения
func (t *tv) off() {
    t.isRunning = false
    fmt.Println("Turning tv off")
}