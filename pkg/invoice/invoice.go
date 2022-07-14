package invoice

import (
	"github.com/Jullentuss/composition/pkg/customer"
	"github.com/Jullentuss/composition/pkg/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}
