package mqs

import (
	"github.com/nats-io/nats.go"
)

var conn *nats.Conn

//Init Init
func Init(url string) error {
	nc, err := nats.Connect(url)
	if err != nil {
		return err
	}

	conn = nc

	return nil
}

// Publish publishes the data argument to the given subject. The data
// argument is left untouched and needs to be correctly interpreted on
// the receiver.
func Publish(subj string, data []byte) error {
	return conn.Publish(subj, data)
}

// Subscribe will express interest in the given subject. The subject
// can have wildcards (partial:*, full:>). Messages will be delivered
// to the associated MsgHandler.
func Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return conn.Subscribe(subj, cb)
}

// SubscribeSync will express interest on the given subject. Messages will
// be received synchronously using Subscription.NextMsg().
func SubscribeSync(subj string) (*nats.Subscription, error) {
	return conn.SubscribeSync(subj)
}

// QueueSubscribe creates an asynchronous queue subscriber on the given subject.
// All subscribers with the same queue name will form the queue group and
// only one member of the group will be selected to receive any given
// message asynchronously.
func QueueSubscribe(subj, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return conn.QueueSubscribe(subj, queue, cb)
}

// QueueSubscribeSync creates a synchronous queue subscriber on the given
// subject. All subscribers with the same queue name will form the queue
// group and only one member of the group will be selected to receive any
// given message synchronously using Subscription.NextMsg().
func QueueSubscribeSync(subj, queue string) (*nats.Subscription, error) {
	return conn.QueueSubscribeSync(subj, queue)
}

// ChanSubscribe will express interest in the given subject and place
// all messages received on the channel.
// You should not close the channel until sub.Unsubscribe() has been called.
func ChanSubscribe(subj string, ch chan *nats.Msg) (*nats.Subscription, error) {
	return conn.ChanSubscribe(subj, ch)
}

// ChanQueueSubscribe will express interest in the given subject.
// All subscribers with the same queue name will form the queue group
// and only one member of the group will be selected to receive any given message,
// which will be placed on the channel.
// You should not close the channel until sub.Unsubscribe() has been called.
// Note: This is the same than QueueSubscribeSyncWithChan.
func ChanQueueSubscribe(subj, group string, ch chan *nats.Msg) (*nats.Subscription, error) {
	return conn.ChanQueueSubscribe(subj, group, ch)
}

// QueueSubscribeSyncWithChan will express interest in the given subject.
// All subscribers with the same queue name will form the queue group
// and only one member of the group will be selected to receive any given message,
// which will be placed on the channel.
// You should not close the channel until sub.Unsubscribe() has been called.
// Note: This is the same than ChanQueueSubscribe.
func QueueSubscribeSyncWithChan(subj, queue string, ch chan *nats.Msg) (*nats.Subscription, error) {
	return conn.QueueSubscribeSyncWithChan(subj, queue, ch)
}

// Close will close the connection to the server. This call will release
// all blocking calls, such as Flush() and NextMsg()
func Close() {
	conn.Close()
}

// IsClosed tests if a Conn has been closed.
func IsClosed() bool {
	return conn.IsClosed()
}

// IsReconnecting tests if a Conn is reconnecting.
func IsReconnecting() bool {
	return conn.IsReconnecting()
}

// IsConnected tests if a Conn is connected.
func IsConnected() bool {
	return conn.IsConnected()
}

// Drain will put a connection into a drain state. All subscriptions will
// immediately be put into a drain state. Upon completion, the publishers
// will be drained and can not publish any additional messages. Upon draining
// of the publishers, the connection will be closed. Use the ClosedCB()
// option to know when the connection has moved from draining to closed.
func Drain() error {
	return conn.Drain()
}

// IsDraining tests if a Conn is in the draining state.
func IsDraining() bool {
	return conn.IsDraining()
}

// Servers returns the list of known server urls, including additional
// servers discovered after a connection has been established.  If
// authentication is enabled, use UserInfo or Token when connecting with
// these urls.
func Servers() []string {
	return conn.Servers()
}

// DiscoveredServers returns only the server urls that have been discovered
// after a connection has been established. If authentication is enabled,
// use UserInfo or Token when connecting with these urls.
func DiscoveredServers() []string {
	return conn.DiscoveredServers()
}

// Status returns the current state of the connection.
func Status() nats.Status {
	return conn.Status()
}

// // Connect to a server
// nc, _ := nats.Connect(nats.DefaultURL)

// // Simple Publisher
// nc.Publish("foo", []byte("Hello World"))

// // Simple Async Subscriber
// nc.Subscribe("foo", func(m *nats.Msg) {
//     fmt.Printf("Received a message: %s\n", string(m.Data))
// })

// // Responding to a request message
// nc.Subscribe("request", func(m *nats.Msg) {
//     m.Respond([]byte("answer is 42"))
// })

// // Simple Sync Subscriber
// sub, err := nc.SubscribeSync("foo")
// m, err := sub.NextMsg(timeout)

// // Channel Subscriber
// ch := make(chan *nats.Msg, 64)
// sub, err := nc.ChanSubscribe("foo", ch)
// msg := <- ch

// // Unsubscribe
// sub.Unsubscribe()

// // Drain
// sub.Drain()

// // Requests
// msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)

// // Replies
// nc.Subscribe("help", func(m *nats.Msg) {
//     nc.Publish(m.Reply, []byte("I can help!"))
// })

// // Drain connection (Preferred for responders)
// // Close() not needed if this is called.
// nc.Drain()

// // Close connection
// nc.Close()
