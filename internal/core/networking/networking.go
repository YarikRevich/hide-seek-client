package networking

var instance *Networking

type Networking struct {
	lanserver *LANServer
	dialer *Dialer
}

func (n *Networking) LANServer() *LANServer{
	return n.lanserver
}

func (n *Networking) Dialer() *Dialer{
	return n.dialer
}

func UseNetworking() *Networking{
	if instance == nil{
		instance = &Networking{
			lanserver: NewLANServer(),
			dialer: NewDialer(),
		}
	}
	return instance
}