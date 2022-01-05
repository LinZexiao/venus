// FETCHED FROM LOTUS: builtin/system/state.go.template

package system

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/venus/pkg/types/specactors/adt"

	system2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/system"
)

var _ State = (*state2)(nil)

func load2(store adt.Store, root cid.Cid) (State, error) {
	out := state2{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make2(store adt.Store) (State, error) {
	out := state2{store: store}
	out.State = system2.State{}
	return &out, nil
}

type state2 struct {
	system2.State
	store adt.Store
}

func (s *state2) GetState() interface{} {
	return &s.State
}
