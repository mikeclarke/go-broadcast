package broadcast

type Message interface{}
type MessageChannel chan Message

type Broadcaster struct {
	inc       chan Message
	registryc chan MessageChannel
	listeners []MessageChannel
}

func (b *Broadcaster) Write(v Message) {
	b.inc <- v
}

func (b *Broadcaster) Listen() (chan Message) {
	c := make(chan Message)
	b.registryc <- c
	return c
}

func (b *Broadcaster) loop() {
	for {
		select {
		case v := <-b.inc:
			for _, c := range(b.listeners) {
				c <- v
			}
		case c := <-b.registryc:
			b.listeners = append(b.listeners, c)
		}
	}
}

func NewBroadcaster() *Broadcaster {
	b := new(Broadcaster)
	go b.loop()
	return b
}
