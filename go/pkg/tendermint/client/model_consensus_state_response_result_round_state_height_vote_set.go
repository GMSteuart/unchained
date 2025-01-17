/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ConsensusStateResponseResultRoundStateHeightVoteSet struct for ConsensusStateResponseResultRoundStateHeightVoteSet
type ConsensusStateResponseResultRoundStateHeightVoteSet struct {
	Round *int32 `json:"round,omitempty"`
	Prevotes []string `json:"prevotes,omitempty"`
	PrevotesBitArray *string `json:"prevotes_bit_array,omitempty"`
	Precommits []string `json:"precommits,omitempty"`
	PrecommitsBitArray *string `json:"precommits_bit_array,omitempty"`
}

// NewConsensusStateResponseResultRoundStateHeightVoteSet instantiates a new ConsensusStateResponseResultRoundStateHeightVoteSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsensusStateResponseResultRoundStateHeightVoteSet() *ConsensusStateResponseResultRoundStateHeightVoteSet {
	this := ConsensusStateResponseResultRoundStateHeightVoteSet{}
	return &this
}

// NewConsensusStateResponseResultRoundStateHeightVoteSetWithDefaults instantiates a new ConsensusStateResponseResultRoundStateHeightVoteSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsensusStateResponseResultRoundStateHeightVoteSetWithDefaults() *ConsensusStateResponseResultRoundStateHeightVoteSet {
	this := ConsensusStateResponseResultRoundStateHeightVoteSet{}
	return &this
}

// GetRound returns the Round field value if set, zero value otherwise.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetRound() int32 {
	if o == nil || o.Round == nil {
		var ret int32
		return ret
	}
	return *o.Round
}

// GetRoundOk returns a tuple with the Round field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetRoundOk() (*int32, bool) {
	if o == nil || o.Round == nil {
		return nil, false
	}
	return o.Round, true
}

// HasRound returns a boolean if a field has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) HasRound() bool {
	if o != nil && o.Round != nil {
		return true
	}

	return false
}

// SetRound gets a reference to the given int32 and assigns it to the Round field.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) SetRound(v int32) {
	o.Round = &v
}

// GetPrevotes returns the Prevotes field value if set, zero value otherwise.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetPrevotes() []string {
	if o == nil || o.Prevotes == nil {
		var ret []string
		return ret
	}
	return o.Prevotes
}

// GetPrevotesOk returns a tuple with the Prevotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetPrevotesOk() ([]string, bool) {
	if o == nil || o.Prevotes == nil {
		return nil, false
	}
	return o.Prevotes, true
}

// HasPrevotes returns a boolean if a field has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) HasPrevotes() bool {
	if o != nil && o.Prevotes != nil {
		return true
	}

	return false
}

// SetPrevotes gets a reference to the given []string and assigns it to the Prevotes field.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) SetPrevotes(v []string) {
	o.Prevotes = v
}

// GetPrevotesBitArray returns the PrevotesBitArray field value if set, zero value otherwise.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetPrevotesBitArray() string {
	if o == nil || o.PrevotesBitArray == nil {
		var ret string
		return ret
	}
	return *o.PrevotesBitArray
}

// GetPrevotesBitArrayOk returns a tuple with the PrevotesBitArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetPrevotesBitArrayOk() (*string, bool) {
	if o == nil || o.PrevotesBitArray == nil {
		return nil, false
	}
	return o.PrevotesBitArray, true
}

// HasPrevotesBitArray returns a boolean if a field has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) HasPrevotesBitArray() bool {
	if o != nil && o.PrevotesBitArray != nil {
		return true
	}

	return false
}

// SetPrevotesBitArray gets a reference to the given string and assigns it to the PrevotesBitArray field.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) SetPrevotesBitArray(v string) {
	o.PrevotesBitArray = &v
}

// GetPrecommits returns the Precommits field value if set, zero value otherwise.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetPrecommits() []string {
	if o == nil || o.Precommits == nil {
		var ret []string
		return ret
	}
	return o.Precommits
}

// GetPrecommitsOk returns a tuple with the Precommits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetPrecommitsOk() ([]string, bool) {
	if o == nil || o.Precommits == nil {
		return nil, false
	}
	return o.Precommits, true
}

// HasPrecommits returns a boolean if a field has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) HasPrecommits() bool {
	if o != nil && o.Precommits != nil {
		return true
	}

	return false
}

// SetPrecommits gets a reference to the given []string and assigns it to the Precommits field.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) SetPrecommits(v []string) {
	o.Precommits = v
}

// GetPrecommitsBitArray returns the PrecommitsBitArray field value if set, zero value otherwise.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetPrecommitsBitArray() string {
	if o == nil || o.PrecommitsBitArray == nil {
		var ret string
		return ret
	}
	return *o.PrecommitsBitArray
}

// GetPrecommitsBitArrayOk returns a tuple with the PrecommitsBitArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) GetPrecommitsBitArrayOk() (*string, bool) {
	if o == nil || o.PrecommitsBitArray == nil {
		return nil, false
	}
	return o.PrecommitsBitArray, true
}

// HasPrecommitsBitArray returns a boolean if a field has been set.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) HasPrecommitsBitArray() bool {
	if o != nil && o.PrecommitsBitArray != nil {
		return true
	}

	return false
}

// SetPrecommitsBitArray gets a reference to the given string and assigns it to the PrecommitsBitArray field.
func (o *ConsensusStateResponseResultRoundStateHeightVoteSet) SetPrecommitsBitArray(v string) {
	o.PrecommitsBitArray = &v
}

func (o ConsensusStateResponseResultRoundStateHeightVoteSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Round != nil {
		toSerialize["round"] = o.Round
	}
	if o.Prevotes != nil {
		toSerialize["prevotes"] = o.Prevotes
	}
	if o.PrevotesBitArray != nil {
		toSerialize["prevotes_bit_array"] = o.PrevotesBitArray
	}
	if o.Precommits != nil {
		toSerialize["precommits"] = o.Precommits
	}
	if o.PrecommitsBitArray != nil {
		toSerialize["precommits_bit_array"] = o.PrecommitsBitArray
	}
	return json.Marshal(toSerialize)
}

type NullableConsensusStateResponseResultRoundStateHeightVoteSet struct {
	value *ConsensusStateResponseResultRoundStateHeightVoteSet
	isSet bool
}

func (v NullableConsensusStateResponseResultRoundStateHeightVoteSet) Get() *ConsensusStateResponseResultRoundStateHeightVoteSet {
	return v.value
}

func (v *NullableConsensusStateResponseResultRoundStateHeightVoteSet) Set(val *ConsensusStateResponseResultRoundStateHeightVoteSet) {
	v.value = val
	v.isSet = true
}

func (v NullableConsensusStateResponseResultRoundStateHeightVoteSet) IsSet() bool {
	return v.isSet
}

func (v *NullableConsensusStateResponseResultRoundStateHeightVoteSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsensusStateResponseResultRoundStateHeightVoteSet(val *ConsensusStateResponseResultRoundStateHeightVoteSet) *NullableConsensusStateResponseResultRoundStateHeightVoteSet {
	return &NullableConsensusStateResponseResultRoundStateHeightVoteSet{value: val, isSet: true}
}

func (v NullableConsensusStateResponseResultRoundStateHeightVoteSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsensusStateResponseResultRoundStateHeightVoteSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


