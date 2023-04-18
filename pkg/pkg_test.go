package pkg

import (
	"context"
	"reflect"
	"testing"
)

type test struct {
	name string
	args args
	want *Response
}

type args struct {
	ctx context.Context
	inn string
}

func TestSomething(t *testing.T) {
	tests := []test{
		{
			name: "Магнолия",
			args: args{
				ctx: context.Background(),
				inn: "7751012274",
			},
			want: &Response{
				INN:  "7751012274",
				KPP:  "775101001",
				NAME: "ООО \"Магнолия\"",
				FIO:  "Амирджанов Шамай Рафаилович",
			},
		},
	}

	for _, tst := range tests {
		t.Run("test", func(t *testing.T) {
			scraper := MyScraper{}
			company, err := scraper.FindCompanyByINN(tst.args.ctx, &Request{INN: tst.args.inn})

			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(company, tst.want) {
				t.Errorf("===\nwanted: %v\ngot:    %v\n===\n", tst.want, company)
			}
		})
	}
}
