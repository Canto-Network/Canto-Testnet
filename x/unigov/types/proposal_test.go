package types

import (
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"reflect"
	"testing"
)

func TestNewLendingMarketProposal(t *testing.T) {
	type args struct {
		title       string
		description string
		m           *LendingMarketMetadata
	}
	tests := []struct {
		name string
		args args
		want types.Content
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLendingMarketProposal(tt.args.title, tt.args.description, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLendingMarketProposal() = %v, want %v", got, tt.want)
			}
		})
	}
}
