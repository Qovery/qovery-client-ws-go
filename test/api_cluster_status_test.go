/*
websocket-gateway

Testing ClusterStatusAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package qovery-ws

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/qovery/qovery-client-ws-go"
)

func Test_qovery-ws_ClusterStatusAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClusterStatusAPIService HandleClusterStatusRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var cluster string

		resp, httpRes, err := apiClient.ClusterStatusAPI.HandleClusterStatusRequest(context.Background(), organization, cluster).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
