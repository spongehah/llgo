package _demo
//
//import "net/url"
//
//type Transport struct {
//	altProto    atomic.Value // of nil or map[string]RoundTripper, key is URI scheme
//	reqMu       sync.Mutex
//	reqCanceler map[cancelKey]func(error)
//	Proxy       func(*Request) (*url.URL, error)
//
//	connsPerHostMu   sync.Mutex
//	connsPerHost     map[connectMethodKey]int
//	connsPerHostWait map[connectMethodKey]wantConnQueue // waiting getConns
//
//	// MaxConnsPerHost optionally limits the total number of
//	// connections per host, including connections in the dialing,
//	// active, and idle states. On limit violation, dials will block.
//	//
//	// Zero means no limit.
//	MaxConnsPerHost int
//}
//
//type RoundTripper interface {
//	RoundTrip(*Request) (*Response, error)
//}
//
//var DefaultTransport RoundTripper = &Transport{
//	//Proxy: ProxyFromEnvironment,
//	Proxy: nil,
//}
//
//func (t *Transport) queueForDial(w *wantConn) {
//	w.beforeDial()
//
//	if t.MaxConnsPerHost <= 0 {
//		go t.dialConnFor(w)
//		return
//	}
//
//	t.connsPerHostMu.Lock()
//	defer t.connsPerHostMu.Unlock()
//
//	if n := t.connsPerHost[w.key]; n < t.MaxConnsPerHost {
//		if t.connsPerHost == nil {
//			t.connsPerHost = make(map[connectMethodKey]int)
//		}
//		t.connsPerHost[w.key] = n + 1
//		go t.dialConnFor(w)
//		return
//	}
//
//	if t.connsPerHostWait == nil {
//		t.connsPerHostWait = make(map[connectMethodKey]wantConnQueue)
//	}
//	q := t.connsPerHostWait[w.key]
//	q.cleanFront()
//	q.pushBack(w)
//	t.connsPerHostWait[w.key] = q
//}
//
//
//// A wantConn records state about a wanted connection
//// (that is, an active call to getConn).
//// The conn may be gotten by dialing or by finding an idle connection,
//// or a cancellation may make the conn no longer wanted.
//// These three options are racing against each other and use
//// wantConn to coordinate and agree about the winning outcome.
//type wantConn struct {
//	cm    connectMethod
//	key   connectMethodKey // cm.key()
//	ctx   context.Context  // context for dial
//	ready chan struct{}    // closed when pc, err pair is delivered
//
//	// hooks for testing to know when dials are done
//	// beforeDial is called in the getConn goroutine when the dial is queued.
//	// afterDial is called when the dial is completed or canceled.
//	beforeDial func()
//	afterDial  func()
//
//	mu  sync.Mutex // protects pc, err, close(ready)
//	pc  *persistConn
//	err error
//}
//
//// A wantConnQueue is a queue of wantConns.
//type wantConnQueue struct {
//	// This is a queue, not a deque.
//	// It is split into two stages - head[headPos:] and tail.
//	// popFront is trivial (headPos++) on the first stage, and
//	// pushBack is trivial (append) on the second stage.
//	// If the first stage is empty, popFront can swap the
//	// first and second stages to remedy the situation.
//	//
//	// This two-stage split is analogous to the use of two lists
//	// in Okasaki's purely functional queue but without the
//	// overhead of reversing the list when swapping stages.
//	head    []*wantConn
//	headPos int
//	tail    []*wantConn
//}
//
//// len returns the number of items in the queue.
//func (q *wantConnQueue) len() int {
//	return len(q.head) - q.headPos + len(q.tail)
//}
//
//// pushBack adds w to the back of the queue.
//func (q *wantConnQueue) pushBack(w *wantConn) {
//	q.tail = append(q.tail, w)
//}
//
//// popFront removes and returns the wantConn at the front of the queue.
//func (q *wantConnQueue) popFront() *wantConn {
//	if q.headPos >= len(q.head) {
//		if len(q.tail) == 0 {
//			return nil
//		}
//		// Pick up tail as new head, clear tail.
//		q.head, q.headPos, q.tail = q.tail, 0, q.head[:0]
//	}
//	w := q.head[q.headPos]
//	q.head[q.headPos] = nil
//	q.headPos++
//	return w
//}
//
//// peekFront returns the wantConn at the front of the queue without removing it.
//func (q *wantConnQueue) peekFront() *wantConn {
//	if q.headPos < len(q.head) {
//		return q.head[q.headPos]
//	}
//	if len(q.tail) > 0 {
//		return q.tail[0]
//	}
//	return nil
//}
//
//// cleanFront pops any wantConns that are no longer waiting from the head of the
//// queue, reporting whether any were popped.
//func (q *wantConnQueue) cleanFront() (cleaned bool) {
//	for {
//		w := q.peekFront()
//		if w == nil || w.waiting() {
//			return cleaned
//		}
//		q.popFront()
//		cleaned = true
//	}
//}
//
