package pkg

import "golang.org/x/net/context"

var _ ScraperServer = (*MyScraper)(nil) // for interface implementation

type MyScraper struct {
	UnimplementedScraperServer
}


func (m MyScraper) FindCompanyByINN(ctx context.Context, request *Request) (*Response, error) {
	//TODO implement me
	panic("implement me")
}

