package cbl

import "encoding/xml"

type OrderHeader struct {
	XMLName                xml.Name       `xml:"OrderHeader"`
	POIssuedDate           DateTime       `xml:"POIssuedDate"`
	ShipByDate             *DateTime      `xml:"ShipByDate"`
	OrderReference         OrderReference `xml:"OrderReference"`
	OrderParty             OrderParty     `xml:"OrderParty"`
	OrderCurrency          string         `xml:"OrderCurrency"`
	OrderLanguage          string         `xml:"OrderLanguage"`
	Payment                Payment        `xml:"Payment"`
	PartialShipmentAllowed bool           `xml:"PartialShipmentAllowed"`
}

type OrderReference struct {
	XMLName                 xml.Name                 `xml:"OrderReference"`
	AccountCode             *AccountCode             `xml:"AccountCode"`
	BuyerReferenceNumber    BuyerReferenceNumber     `xml:"BuyerRefNum"`
	SupplierReferenceNumber *SupplierReferenceNumber `xml:"SupplierRefNum"`
	ListOfReferenceCoded    ListOfReferenceCoded     `xml:"ListOfReferenceCoded"`
}

type AccountCode struct {
	XMLName   xml.Name  `xml:"AccountCode"`
	Reference Reference `xml:"Reference"`
}

type BuyerReferenceNumber struct {
	XMLName   xml.Name  `xml:"BuyerRefNum"`
	Reference Reference `xml:"Reference"`
}

type SupplierReferenceNumber struct {
	XMLName   xml.Name  `xml:"SupplierRefNum"`
	Reference Reference `xml:"Reference"`
}

type Reference struct {
	XMLName         xml.Name  `xml:"Reference"`
	ReferenceNumber string    `xml:"RefNum"`
	ReferenceDate   *DateTime `xml:"RefDate"`
}

type ListOfReferenceCoded struct {
	XMLName        xml.Name         `xml:"ListOfReferenceCoded"`
	ReferenceCoded []ReferenceCoded `xml:"ReferenceCoded"`
}

type ReferenceCoded struct {
	XMLName      xml.Name `xml:"ReferenceCoded"`
	RefNum       string   `xml:"RefNum"`
	RefCode      string   `xml:"RefCode"`
	RefCodeOther string   `xml:"RefCodeOther,omitempty"`
}

type OrderParty struct {
	XMLName       xml.Name      `xml:"OrderParty"`
	BuyerParty    BuyerParty    `xml:"BuyerParty"`
	SupplierParty SupplierParty `xml:"SupplierParty"`
	ShipToParty   ShipToParty   `xml:"ShipToParty"`
	BillToParty   BillToParty   `xml:"BillToParty"`
}

type BuyerParty struct {
	XMLName     xml.Name    `xml:"BuyerParty"`
	NameAddress NameAddress `xml:"NameAddress"`
}

type SupplierParty struct {
	XMLName     xml.Name    `xml:"SupplierParty"`
	NameAddress NameAddress `xml:"NameAddress"`
}

type ShipToParty struct {
	XMLName     xml.Name    `xml:"ShipToParty"`
	NameAddress NameAddress `xml:"NameAddress"`
}

type BillToParty struct {
	XMLName     xml.Name    `xml:"BillToParty"`
	NameAddress NameAddress `xml:"NameAddress"`
}

type NameAddress struct {
	XMLName         xml.Name `xml:"NameAddress"`
	Name1           string   `xml:"Name1"`
	Name2           string   `xml:"Name2,omitempty"`
	Name3           string   `xml:"Name3,omitempty"`
	Address1        string   `xml:"Address1"`
	Address2        string   `xml:"Address2,omitempty"`
	Address3        string   `xml:"Address3,omitempty"`
	Address4        string   `xml:"Address4,omitempty"`
	Address5        string   `xml:"Address5,omitempty"`
	City            string   `xml:"City"`
	StateOrProvince string   `xml:"StateOrProvince"`
	PostalCode      string   `xml:"PostalCode"`
	Country         string   `xml:"Country"`
}

type Payment struct {
	XMLName     xml.Name `xml:"Payment"`
	PaymentMean string   `xml:"PaymentMean"`
}
