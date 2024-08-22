package main

type resp struct {
	field string
}

type respWrapper struct {
	resp *resp
}

type requestAndChan struct {
	resc chan respWrapper
	gone chan struct{}
}

func main() {
	reqch := make(chan requestAndChan, 1)
	closech := make(chan struct{}, 1)

	go func() {
		rc := <-reqch
		println("receive reqch")

		resp := &resp{field: "field"}

		select {
		case rc.resc <- respWrapper{resp: resp}:
			println("return respWrapper")
		case <-rc.gone:
		}
	}()

	gone := make(chan struct{}, 1)
	resc := make(chan respWrapper, 1)

	reqch <- requestAndChan{
		resc: resc,
		gone: gone,
	}

	select {
	case <-closech:
		println("resp is nil")
	case rc := <-resc:
		println("receive rc")
		println("resp: ", rc.resp.field)
	}
}
