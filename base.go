package didar

type ContactField map[string]interface{}

type Contact struct {
	Id           string       `json:"Id,omitempty"`
	FirstName    string       `json:"FirstName"`
	LastName     string       `json:"LastName"`
	Email        string       `json:"Email,omitempty"`
	MobilePhone  string       `json:"MobilePhone,omitempty"`
	WorkPhone    string       `json:"WorkPhone,omitempty"`
	CompanyName  string       `json:"CompanyName,omitempty"`
	CompanyId    string       `json:"CompanyId,omitempty"`
	Fields       ContactField `json:"Fields,omitempty"`
	SegmentIds   []string     `json:"SegmentIds,omitempty"`
	OwnerId      string       `json:"OwnerId,omitempty"`
	CustomerCode string       `json:"CusomterCode,omitempty"`
}

type ContactRequest struct {
	Contact Contact `json:"Contact"`
}

type ContactResponse struct {
	Response struct {
		Id string `json:"Id"`
	} `json:"Response"`
}
