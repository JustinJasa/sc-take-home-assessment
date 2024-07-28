package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
	Limit int    // Number of items per page
	Token string // Pagination token for the next page
}

type FetchFolderResponse struct {
	Folders []*Folder
	Token   string // Token for the next page, empty if no more results
}
