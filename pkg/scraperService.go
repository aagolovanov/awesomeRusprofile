package pkg

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ ScraperServer = (*MyScraper)(nil) // for interface implementation

type MyScraper struct {
	UnimplementedScraperServer
}

func (m MyScraper) FindCompanyByINN(ctx context.Context, request *Request) (*Response, error) {
	company, err := GetMainInfo(request.INN)
	if err != nil {
		if err.Error() == "NotFound" {
			return nil, status.Error(codes.NotFound, "No companies with provided INN")
		} else {
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
	}

	kpp, err := GetCompanyKPP(company)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return &Response{
		INN:  company.INN,
		KPP:  kpp,
		NAME: company.Name,
		FIO:  company.FIO,
	}, nil
}
