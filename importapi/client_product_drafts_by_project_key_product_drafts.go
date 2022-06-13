package importapi

// Generated file, please do not change!!!

type ByProjectKeyProductDraftsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyProductDraftsRequestBuilder) ImportContainers() *ByProjectKeyProductDraftsImportContainersRequestBuilder {
	return &ByProjectKeyProductDraftsImportContainersRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyProductDraftsRequestBuilder) ImportSinkKeyWithImportSinkKeyValue(importSinkKey string) *ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyRequestBuilder {
	return &ByProjectKeyProductDraftsImportSinkKeyByImportSinkKeyRequestBuilder{
		importSinkKey: importSinkKey,
		projectKey:    rb.projectKey,
		client:        rb.client,
	}
}
