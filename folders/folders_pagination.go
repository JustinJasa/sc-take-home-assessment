/**

Explanation of code:

The getfolders function is designed to retrieve a paginated list of folders.

It does this byfirst fetching all folders for the given organisation,
it then determines the number of folders to return by limit
which is defaulted to 10.

the function calculates the starting point of the page of results,
it decondes the token to determine the starting index

it then returns a slice of folders based on the start index and limit,
if there are more folders beyond the page, the function generates a token for the next page

it returns a list of all folders, and the token for the next page,
if there are more folders beyond the page.

**/

package folders

import (
	"fmt"
)

// DefaultLimit is the default number of items per page if not specified
const DefaultLimit = 10

// GetFolders retrieves a paginated list of folders for a given organization
func GetFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Fetch all folders for the organization
	allFolders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch folders: %w", err)
	}

	// Set default limit if not provided
	limit := req.Limit
	if limit <= 0 {
		limit = DefaultLimit
	}

	// Calculate start index based on the token
	startIndex := 0
	if req.Token != "" {
		startIndex, err = decodeToken(req.Token)
		if err != nil {
			return nil, fmt.Errorf("invalid token: %w", err)
		}
	}

	// Calculate end index
	endIndex := startIndex + limit
	if endIndex > len(allFolders) {
		endIndex = len(allFolders)
	}

	// Get the subset of folders for this page
	paginatedFolders := allFolders[startIndex:endIndex]

	// Generate next token if there are more results
	var nextToken string
	if endIndex < len(allFolders) {
		nextToken = encodeToken(endIndex)
	}

	return &FetchFolderResponse{
		Folders: paginatedFolders,
		Token:   nextToken,
	}, nil
}

// encodeToken converts an index to a token string
func encodeToken(index int) string {
	return fmt.Sprintf("%d", index)
}

// decodeToken converts a token string to an index
func decodeToken(token string) (int, error) {
	var index int
	_, err := fmt.Sscanf(token, "%d", &index)
	return index, err
}
