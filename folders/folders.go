/**
Component 1:
- What does this code do?

the code in this file is retrieving folders from a given organisation.
the GetAllFolders function takes a FetchFolderRequest and returns a FetchFolderResponse.
it uses FetchAllFoldersByOrgID to retrieve the folders, then processes data to return it in the expected format

The FetchAllFoldersByOrgID function takes an orgID and returns a slice of folders based on the provided organisationID.

- What improvements can be made to the code?

1. Remove unused variables
2. Unnecessary type conversions
	- code converts folders from []*Folder to []Folder and back to []*Folder.
3. Error handling
	- The error returned by FetchAllFoldersByOrgID is not handled, this can be checked and handled appropriately.
**/

package folders

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// revised functions

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	folders, err := FetchAllFoldersByOrgID(req.OrgID)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch folders: %w", err)
	}

	return &FetchFolderResponse{Folders: folders}, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
