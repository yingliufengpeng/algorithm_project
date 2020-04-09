package datastruct

// TuopuElement is an element of a linked tuopu.
type TuopuElement struct {
	// left and right pointers in the tuopu node elements.
	upstream []*TuopuElement
	// The tuopu to which this element belongs.
	tuopu *Tuopu
	// The value stored with this element.
	Key interface{}
}

type LinkData map[interface{}]*TuopuElement

// Tuopu a tuopu graph
type Tuopu struct {
	link map[interface{}](LinkData)
	node map[interface{}](*TuopuElement)
	len  int
}

// NewTuopu new a tuopu
func NewTuopu() *Tuopu {
	return &Tuopu{
		link: make(map[interface{}]LinkData),
		node: make(map[interface{}]*TuopuElement),
	}
}

// getset
func (tp *Tuopu) getset(k interface{}) (e *TuopuElement) {
	if nil != tp.node[k] {
		e = tp.node[k]
	} else {
		e = &TuopuElement{
			upstream: make([]*TuopuElement, 0),
			tuopu:    tp,
			Key:      k,
		}
		tp.node[k] = e
	}
	return e
}

// getset
func (tp *Tuopu) getlink(k interface{}) (l LinkData) {
	if nil == tp.link[k] {
		e := tp.getset(k)
		tp.link[k] = make(LinkData)
		tp.link[k][k] = e
	}
	l = tp.link[k]
	return l
}

// Push push a node input
func (tp *Tuopu) Push(k interface{}, vv ...interface{}) (e *TuopuElement) {
	knode := tp.getset(k)
	klink := tp.getlink(k)
	for _, v := range vv {
		vnode := tp.getset(v)
		klink[v] = vnode
		knode.upstream = append(knode.upstream, vnode)
	}
	tp.link[k] = klink
	tp.len += len(klink)
	return e
}

// Get get a node and link from Tuopu
func (tp *Tuopu) Get(k interface{}) (e *TuopuElement, upstream []*TuopuElement) {
	if nil != tp.link[k] {
		e = tp.link[k][k]
		upstream = tp.node[k].upstream
	}
	return e, upstream
}

// BeGet get a node and link from Tuopu
func (tp *Tuopu) BeGet(k interface{}) (e *TuopuElement, upstream []*TuopuElement) {
	if nil != tp.link[k] {
		e = tp.link[k][k]
		upstream = tp.node[k].upstream
	}
	return e, upstream
}
