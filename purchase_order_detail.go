package cbl

import "encoding/xml"

type ListOfOrderDetail struct {
	XMLName     xml.Name      `xml:"ListOfOrderDetail"`
	OrderDetail []OrderDetail `xml:"OrderDetail"`
}

type OrderDetail struct {
	XMLName                xml.Name               `xml:"OrderDetail"`
	BaseItemDetail         BaseItemDetail         `xml:"BaseItemDetail"`
	RequestedDeliveryDate  string                 `xml:"RequestedDeliveryDate"`
	Tax                    Tax                    `xml:"Tax"`
	BuyerExpectedUnitPrice BuyerExpectedUnitPrice `xml:"BuyerExpectedUnitPrice"`
}

type BaseItemDetail struct {
	XMLName                xml.Name        `xml:"BaseItemDetail"`
	LineItemNum            int             `xml:"LineItemNum"`
	SupplierPartNum        SupplierPartNum `xml:"SupplierPartNum"`
	CommodityCode          string          `xml:"CommodityCode"`
	ItemDescription        string          `xml:"ItemDescription"`
	Quantity               Quantity        `xml:"Quantity"`
	FinalRecipient         FinalRecipient  `xml:"FinalRecipient"`
	OffCatalogFlag         bool            `xml:"OffCatalogFlag"`
	CostCenter             string          `xml:"CostCenter"`
	CostType               string          `xml:"CostType"`
	SpecialDeprecationFlag bool            `xml:"SpecialDeprecationFlag"`
	LowValueAssetFlag      bool            `xml:"LowValueAssetFlag"`
	SpecialHandlingFlag    bool            `xml:"SpecialHandlingFlag"`
}

type SupplierPartNum struct {
	XMLName xml.Name `xml:"SupplierPartNum"`
	PartID  string   `xml:"PartID"`
}

type Quantity struct {
	XMLName       xml.Name `xml:"Quantity"`
	Qty           float64  `xml:"Qty"`
	UnitOfMeasure string   `xml:"UnitOfMeasure"`
}

type FinalRecipient struct {
	XMLName          xml.Name         `xml:"FinalRecipient"`
	NameAddress      NameAddress      `xml:"NameAddress"`
	ReceivingContact ReceivingContact `xml:"ReceivingContact"`
}

type ReceivingContact struct {
	XMLName     xml.Name `xml:"ReceivingContact"`
	ContactName string   `xml:"ContactName"`
	Telephone   string   `xml:"Telephone"`
	Email       string   `xml:"Email"`
}

type Tax struct {
	XMLName       xml.Name `xml:"Tax"`
	TaxPercent    float64  `xml:"TaxPercent"`
	TaxAmount     float64  `xml:"TaxAmount"`
	TaxableAmount float64  `xml:"TaxableAmount"`
}

type BuyerExpectedUnitPrice struct {
	XMLName   xml.Name `xml:"BuyerExpectedUnitPrice"`
	UnitPrice float64  `xml:"UnitPrice"`
}
