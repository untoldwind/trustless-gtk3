package daemon

import (
	"encoding/json"
	"net/http"

	"github.com/awnumar/memguard"

	"github.com/leanovate/microtools/logging"
	"github.com/leanovate/microtools/rest"
	"github.com/untoldwind/trustless/api"
	"github.com/untoldwind/trustless/secrets"
)

// MasterKeyResource is a REST resource representing the master key of the secret store
type MasterKeyResource struct {
	rest.ResourceBase
	logger  logging.Logger
	secrets secrets.Secrets
}

// NewMasterKeyResource creates a new MasterKeyResource
func NewMasterKeyResource(secrets secrets.Secrets, logger logging.Logger) *MasterKeyResource {
	return &MasterKeyResource{
		logger:  logger.WithField("resource", "masterkey"),
		secrets: secrets,
	}
}

// Self link to the resource
func (MasterKeyResource) Self() rest.Link {
	return rest.SimpleLink("/v1/masterkey")
}

// Get the status of the master key
func (r *MasterKeyResource) Get(request *http.Request) (interface{}, error) {
	status, err := r.secrets.Status(request.Context())
	if err != nil {
		return nil, err
	}
	return &api.MasterKey{
		Locked:     status.Locked,
		AutolockAt: status.AutolockAt,
	}, nil
}

// Update unlocks the master keys.
func (r *MasterKeyResource) Update(request *http.Request) (interface{}, error) {
	var unlock api.MasterKeyUnlock

	defer request.Body.Close()
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&unlock); err != nil {
		return nil, rest.HTTPBadRequest.WithDetails(err.Error())
	}
	locked, err := memguard.NewImmutableFromBytes([]byte(unlock.Passphrase))
	if err != nil {
		return nil, err
	}
	defer locked.Destroy()
	unlock.Passphrase = string(locked.Buffer())
	if err := r.secrets.Unlock(request.Context(), unlock.Name, unlock.Email, unlock.Passphrase); err != nil {
		return nil, rest.HTTPInternalServerError(err)
	}
	return nil, nil
}

// Delete locks the master key
func (r *MasterKeyResource) Delete(request *http.Request) (interface{}, error) {
	r.secrets.Lock(request.Context())
	return nil, nil
}
