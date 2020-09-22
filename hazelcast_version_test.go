package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHazelcastVersionServiceOp_List(t *testing.T) {
	//given
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	defer server.Close()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		var request GraphQLQuery
		json.NewDecoder(r.Body).Decode(&request)

		if strings.Contains(request.Query, "hazelcastVersions") {
			fmt.Fprint(w, `{"data":{"response":[{"version":"4.0"},{"version":"3.12.6"},{"version":"3.12.5"},{"version":"3.12.4"},{"version":"3.12.3"},{"version":"3.12.2"},{"version":"3.12.1"},{"version":"3.12"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dW51c0BoYXplbGNhc3QuY29tIiwicm9sZXMiOlt7InRlYW1JZCI6IjMiLCJhdXRob3JpdHkiOiJURUFNX0FETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9GSU5BTkNFIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkFETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJVU0VSIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkRFRElDQVRFRF9VU0VSIn0seyJ0ZWFtSWQiOiIyIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjoiMiIsImF1dGhvcml0eSI6IlRFQU1fRklOQU5DRSJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJBQ0NPVU5USU5HIn1dLCJ0b2tlbiI6IjE1YjY5MWQxLThmOWUtNGQ4Zi04NzNkLTk4ZWI0NGU0ODk5NSIsImV4cCI6MTc1NzQyODg3MH0.HM3vLZbR4H8LIu0Quqm3dqwCj6V_XAYtaUGg5ZQkeefgvMA1LIoxJRyPgZYhJgJJ_aHPnBZ08wJwCrFADGHitA"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	hazelcastVersions, _, _ := NewHazelcastVersionService(client).List(context.TODO())

	//then
	assert.Len(t, *hazelcastVersions, 8)
}

func ExampleHazelcastVersionService_list() {
	client, _, _ := New()
	hazelcastVersions, _, _ := client.HazelcastVersion.List(context.Background())
	fmt.Printf("Result: %#v", hazelcastVersions)
	//Output:Result: &[]models.EnterpriseHazelcastVersion{models.EnterpriseHazelcastVersion{Version:"4.0", UpgradeableVersions:[]string{}}, models.EnterpriseHazelcastVersion{Version:"3.12.6", UpgradeableVersions:[]string{}}, models.EnterpriseHazelcastVersion{Version:"3.12.5", UpgradeableVersions:[]string{"3.12.6"}}, models.EnterpriseHazelcastVersion{Version:"3.12.4", UpgradeableVersions:[]string{"3.12.6", "3.12.5"}}, models.EnterpriseHazelcastVersion{Version:"3.12.3", UpgradeableVersions:[]string{"3.12.6", "3.12.5", "3.12.4"}}, models.EnterpriseHazelcastVersion{Version:"3.12.2", UpgradeableVersions:[]string{"3.12.6", "3.12.5", "3.12.4", "3.12.3"}}, models.EnterpriseHazelcastVersion{Version:"3.12.1", UpgradeableVersions:[]string{"3.12.6", "3.12.5", "3.12.4", "3.12.3", "3.12.2"}}, models.EnterpriseHazelcastVersion{Version:"3.12", UpgradeableVersions:[]string{"3.12.6", "3.12.5", "3.12.4", "3.12.3", "3.12.2", "3.12.1"}}}
}
