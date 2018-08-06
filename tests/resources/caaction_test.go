package nitro

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaaction(t *testing.T) {
	client := nitro.NewNitroClient("http://10.2.0.232", "nsroot", "Charlie")

	resource, tearDown := setup_caaction(t, client)

	if resource == nil {
		return
	}

	err := client.AddCaaction(*resource)

	assert.NoError(t, err)

	exists, err := client.ExistsCaaction(resource.Name + "-unknown")

	assert.NoError(t, err)
	assert.Equal(t, exists, false)

	exists, err = client.ExistsCaaction(resource.Name)

	assert.NoError(t, err)
	assert.Equal(t, exists, true)

	res, err := client.GetCaaction(resource.Name)

	assert.NoError(t, err)
	assert.NotNil(t, res)

	list, err := client.ListCaaction()

	assert.NoError(t, err)
	assert.NotNil(t, list)

	// err = client.UpdateCaaction(*resource)

	// assert.NoError(t, err)

	err = client.RenameCaaction(resource.Name, resource.Name+"-rename")

	assert.NoError(t, err)

	err = client.DeleteCaaction(resource.Name + "-rename")

	assert.NoError(t, err)
	if tearDown != nil {
		tearDown()
	}
}
