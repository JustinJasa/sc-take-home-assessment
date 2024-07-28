package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("Valid organisation ID", func(t *testing.T) {
		orgID := uuid.FromStringOrNil(folders.DefaultOrgID)
		req := &folders.FetchFolderRequest{OrgID: orgID}

		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.Folders)
		for _, folder := range resp.Folders {
			assert.Equal(t, orgID, folder.OrgId)
		}
	})

	t.Run("Invalid organisation ID", func(t *testing.T) {
		orgID := uuid.Must(uuid.NewV4())
		req := &folders.FetchFolderRequest{OrgID: orgID}

		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, resp.Folders)

	})

	t.Run("Zero UUID", func(t *testing.T) {
		req := &folders.FetchFolderRequest{OrgID: uuid.Nil}

		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, resp.Folders)
	})
}
