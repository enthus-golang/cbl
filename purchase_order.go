package cbl

import (
	"encoding/xml"
)

type PurchaseOrder struct {
	XMLName           xml.Name          `xml:"PurchaseOrder"`
	OrderHeader       OrderHeader       `xml:"OrderHeader"`
	ListOfOrderDetail ListOfOrderDetail `xml:"ListOfOrderDetail"`
	OrderSummary      OrderSummary      `xml:"OrderSummary"`
}

type OrderSummary struct {
	XMLName      xml.Name `xml:"OrderSummary"`
	TotalAmount  float64  `xml:"TotalAmount"`
	TotalLineNum float64  `xml:"TotalLineNum"`
}
