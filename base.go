package didar

import "time"

type Fields map[string]interface{}

type Contact struct {
	Id           string   `json:"Id,omitempty"`
	FirstName    string   `json:"FirstName"`
	LastName     string   `json:"LastName"`
	Email        string   `json:"Email,omitempty"`
	MobilePhone  string   `json:"MobilePhone,omitempty"`
	WorkPhone    string   `json:"WorkPhone,omitempty"`
	CompanyName  string   `json:"CompanyName,omitempty"`
	CompanyId    string   `json:"CompanyId,omitempty"`
	Fields       Fields   `json:"Fields,omitempty"`
	SegmentIds   []string `json:"SegmentIds,omitempty"`
	OwnerId      string   `json:"OwnerId,omitempty"`
	CustomerCode string   `json:"CusomterCode,omitempty"`
}

type ContactRequest struct {
	Contact Contact `json:"Contact"`
}

type ContactResponse struct {
	Response struct {
		Id string `json:"Id"`
	} `json:"Response"`
}

type Deal struct {
	Title           string     `json:"Title"`
	Description     string     `json:"Description"`
	TaxPercent      int        `json:"TaxPercent"`
	ContactId       string     `json:"ContactId"`
	RegisterDate    *time.Time `json:"RegisterDate"`
	DealItems       []DealItem `json:"DealItems"`
	PipelineStageId string     `json:"PipelineStageId"`
	Fields          Fields     `json:"Fields"`
}
type DealItem struct {
	Description string  `json:"Description"`
	Discount    int     `json:"Discount"`
	ProductId   string  `json:"ProductId"`
	Quantity    float64 `json:"Quantity"`
	UnitPrice   int64   `json:"UnitPrice"`
}
type DealRequest struct {
	Deal Deal `json:"Deal"`
}

type DealResponse struct {
	Response struct {
		Id string `json:"Id"`
	} `json:"Response"`
}
