package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"reflect"
	"testing"
)

func createFullMetadata(account string, propId uint64, values []uint64, calldatas []string, signatures []string) LendingMarketMetadata {
	return LendingMarketMetadata{
		Account:    nil,
		PropId:     0,
		Values:     nil,
		Calldatas:  nil,
		Signatures: nil,
	}
}

func createMetadata(account string, propId uint64, values []uint64, calldatas []string, signatures []string) LendingMarketMetadata {
	return createFullMetadata(account, propId, values, calldatas, signatures)
}

func TestLendingMarketProposal(t *testing.T) {
	validMetaData := LendingMarketMetadata{
		Account:    nil,
		PropId:     0,
		Values:     nil,
		Calldatas:  nil,
		Signatures: nil,
	}

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
		// Valid tests
		{
			name: "Test",
			args: struct {
				title       string
				description string
				m           *LendingMarketMetadata
			}{
				title:       "Test",
				description: "Test",
				m:           &validMetaData,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		fmt.Println(NewLendingMarketProposal(tt.args.title, tt.args.description, tt.args.m))
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLendingMarketProposal(tt.args.title, tt.args.description, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLendingMarketProposal() = %v, want %v", got, tt.want)
			}
		})
	}
}
