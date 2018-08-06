package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRewritepolicy(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_rewritepolicy(t, client)

	if resource == nil {
		return
	}

	err := client.AddRewritepolicy(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsRewritepolicy(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsRewritepolicy(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetRewritepolicy(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListRewritepolicy()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateRewritepolicy(*resource)

	// assert.NoError(t, err)

	err = client.RenameRewritepolicy(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteRewritepolicy(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
