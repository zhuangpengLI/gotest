package bridge

import "fmt"

// 桥接模式分离抽象部分和实现部分。使得两部分独立扩展。
// 桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。
// 策略模式使抽象部分和实现部分分离，可以独立变化。

//抽象接口
type AbstractMessage interface {
	SendMessage(text, to string)
}

//实现接口
type MessageImplementer interface {
	Send(text, to string)
}

type MessageSMS struct{}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

//实际的实现函数
func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}

type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

//实际的实现函数
func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}

//桥接对象
type CommonMessage struct {
	method MessageImplementer
}

//可传入不同消息
func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

//抽象接口实际调用实现接口
func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

//桥接对象
type UrgencyMessage struct {
	method MessageImplementer
}

//可传入不同消息
func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

//抽象接口实际调用实现接口
func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
