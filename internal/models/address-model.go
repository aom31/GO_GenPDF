package models

type AddressInfo struct {
	NameCompany             string `json:"nameCompany"`
	Address                 string `json:"address"`
	TaxIdentificationNumber int    `json:"taxIdentificationNumber"`
}
