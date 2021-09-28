package provider

var (
	p IProvider
)

type Provider struct {

}

func (pr *Provider) User() interface{}{
	return nil
}

type IProvider interface {
	User() interface{}
}

func UseStorageProvider()IProvider{
	if p == nil{
		p = new(Provider)
	}
	return p
}