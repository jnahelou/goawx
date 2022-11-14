package awx

// MeService implements awx login apis.
type MeService struct {
	client *Client
}

type MeResponse struct {
	Pagination
	Results []*Me `json:"results"`
}

const MeAPIEndpoint = "/api/v2/me/"

// Me get user info from awx servers.
func (l *MeService) Me() ([]*Me, *MeResponse, error) {
	result := new(MeResponse)
	resp, err := l.client.Requester.GetJSON(MeAPIEndpoint, result, nil)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}
