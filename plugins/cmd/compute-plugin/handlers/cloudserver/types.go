package cloudserver

// TODO: Define types for Cloudserver based on OpenAPI specification
// These types should match the Aruba Cloud API schema

type CloudserverDto struct {
	Metadata   *MetadataDto `json:"metadata,omitempty"`
	Properties interface{}  `json:"properties,omitempty"` // TODO: Define specific properties type
}

type MetadataDto struct {
	Name     string       `json:"name,omitempty"`
	Location *LocationDto `json:"location,omitempty"`
	Tags     []string     `json:"tags,omitempty"`
}

type LocationDto struct {
	Value string `json:"value,omitempty"`
}

type CloudserverResponseDto struct {
	Metadata   *MetadataResponseDto `json:"metadata,omitempty"`
	Status     *StatusResponseDto   `json:"status,omitempty"`
	Properties interface{}          `json:"properties,omitempty"` // TODO: Define specific properties type
}

type MetadataResponseDto struct {
	ID           string               `json:"id,omitempty"`
	URI          string               `json:"uri,omitempty"`
	Name         string               `json:"name,omitempty"`
	Location     *LocationResponseDto `json:"location,omitempty"`
	Project      *ProjectResponseDto  `json:"project,omitempty"`
	Tags         []string             `json:"tags,omitempty"`
	Category     *CategoryResponseDto `json:"category,omitempty"`
	CreationDate string               `json:"creationDate,omitempty"`
	CreatedBy    string               `json:"createdBy,omitempty"`
	UpdateDate   string               `json:"updateDate,omitempty"`
	UpdatedBy    string               `json:"updatedBy,omitempty"`
	Version      string               `json:"version,omitempty"`
}

type LocationResponseDto struct {
	Code    string `json:"code,omitempty"`
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
}

type ProjectResponseDto struct {
	ID string `json:"id,omitempty"`
}

type CategoryResponseDto struct {
	Name     string               `json:"name,omitempty"`
	Provider *ProviderResponseDto `json:"provider,omitempty"`
	Type     *TypeResponseDto     `json:"type,omitempty"`
}

type ProviderResponseDto struct {
	Name string `json:"name,omitempty"`
}

type TypeResponseDto struct {
	Name string `json:"name,omitempty"`
}

type StatusResponseDto struct {
	State            string `json:"state,omitempty"`
	Message          string `json:"message,omitempty"`
	ProvisioningDate string `json:"provisioningDate,omitempty"`
}
