package types

import (
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/suite"
	"reflect"
	"testing"
)

type ProposalTestSuite struct {
	suite.Suite
}

func TestProposalTestSuite(t *testing.T) {
	suite.Run(t, new(ProposalTestSuite))
}

func (suite *ProposalTestSuite) TestKeysTypes() {
	suite.Require().Equal("unigov", (&LendingMarketProposal{}).ProposalRoute())
	suite.Require().Equal("Lending-Market", (&LendingMarketProposal{}).ProposalType())
	suite.Require().Equal("unigov", (&TreasuryProposal{}).ProposalRoute())
	suite.Require().Equal("Treasury", (&TreasuryProposal{}).ProposalType())
}

func createLendingMarketFullMetadata(account []string, propId uint64, values []uint64, calldatas []string, signatures []string) LendingMarketMetadata {
	return LendingMarketMetadata{
		Account:    account,
		PropId:     propId,
		Values:     values,
		Calldatas:  calldatas,
		Signatures: signatures,
	}
}

func createLendingMarketMetadata(account []string, propId uint64, values []uint64, calldatas []string, signatures []string) LendingMarketMetadata {
	return createLendingMarketFullMetadata(account, propId, values, calldatas, signatures)
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
		name       string
		args       args
		want       types.Content
		expectPass bool
	}{
		// Valid tests
		{
			name: "Valid Test",
			args: struct {
				title       string
				description string
				m           *LendingMarketMetadata
			}{
				title:       "Test",
				description: "Valid",
				m:           &validMetaData,
			},
			want:       nil,
			expectPass: false,
		},
		// Invalid tests
		//{
		//	name: "Invalid Test",
		//	args: struct {
		//		title       string
		//		description string
		//		m           *LendingMarketMetadata
		//	}{
		//		title:       "Test",
		//		description: "Invalid",
		//		m:           createLendingMarketMetadata([]string{}, 0, []uint64{}, []string{}, []string{}),
		//	},
		//	want: nil,
		//	expectPass: true,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewLendingMarketProposal(tt.args.title, tt.args.description, tt.args.m)
			e := tx.ValidateBasic()

			if tt.expectPass {
				if e != nil {
					t.Errorf("Error: %v", e.Error())
				}
			} else {
				if e == nil {
					t.Errorf("Error: %v", e.Error())
				}
			}

			if !reflect.DeepEqual(tx, tt.want) {
				t.Errorf("NewLendingMarketProposal() = %v, want %v", tx, tt.want)
			}
		})
	}
}

func createTreasuryFullMetadata(propId uint64, recipient string, amount uint64, denom string) TreasuryProposalMetadata {
	return TreasuryProposalMetadata{
		PropID:    propId,
		Recipient: recipient,
		Amount:    amount,
		Denom:     denom,
	}
}

func createTreasuryMetadata(propId uint64, recipient string, amount uint64, denom string) TreasuryProposalMetadata {
	return createTreasuryMetadata(propId, recipient, amount, denom)
}

func TestNewTreasuryProposal(t *testing.T) {
	validMetaData := TreasuryProposalMetadata{
		PropID:    0,
		Recipient: "",
		Amount:    0,
		Denom:     "",
	}

	type args struct {
		title       string
		description string
		tm          *TreasuryProposalMetadata
	}

	tests := []struct {
		name       string
		args       args
		want       types.Content
		expectPass bool
	}{
		// Valid tests
		{
			name: "Valid Test",
			args: struct {
				title       string
				description string
				tm          *TreasuryProposalMetadata
			}{
				title:       "Test",
				description: "Valid",
				tm:          &validMetaData,
			},
			want:       nil,
			expectPass: false,
		},
		// Invalid tests
		//{
		//	name: "Invalid Test",
		//	args: struct {
		//		title       string
		//		description string
		//		tm           *TreasuryProposalMetaData
		//	}{
		//		title:       "Test",
		//		description: "Invalid",
		//		tm:           createLendingMarketMetadata([]string{}, 0, []uint64{}, []string{}, []string{}),
		//	},
		//	want: nil,
		//	expectPass: true,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := NewTreasuryProposal(tt.args.title, tt.args.description, tt.args.tm)
			e := tx.ValidateBasic()

			if tt.expectPass {
				if e != nil {
					t.Errorf("Error: %v", e.Error())
				}
			} else {
				if e == nil {
					t.Errorf("Error: %v", e.Error())
				}
			}

			if !reflect.DeepEqual(tx, tt.want) {
				t.Errorf("NewTreasuryProposal() = %v, want %v", tx, tt.want)
			}
		})
	}
}

func TestTreasuryProposal_FromTreasuryToLendingMarket(t *testing.T) {
	validMetaData := TreasuryProposalMetadata{
		PropID:    0,
		Recipient: "",
		Amount:    0,
		Denom:     "",
	}

	type args struct {
		title       string
		description string
		m           *TreasuryProposalMetadata
	}

	tests := []struct {
		name       string
		args       args
		want       *LendingMarketProposal
		expectPass bool
	}{
		// Valid tests
		{
			name: "Valid Test",
			args: struct {
				title       string
				description string
				m           *TreasuryProposalMetadata
			}{
				title:       "Test",
				description: "Valid",
				m:           &validMetaData,
			},
			want:       nil,
			expectPass: false,
		},
		// Invalid tests
		//{
		//	name: "Invalid Test",
		//	args: struct {
		//		title       string
		//		description string
		//		m           *LendingMarketMetadata
		//	}{
		//		title:       "Test",
		//		description: "Invalid",
		//		m:           createLendingMarketMetadata([]string{}, 0, []uint64{}, []string{}, []string{}),
		//	},
		//	want: nil,
		//	expectPass: true,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tp := &TreasuryProposal{
				Title:       tt.args.title,
				Description: tt.args.description,
				Metadata:    tt.args.m,
			}
			tx := tp.FromTreasuryToLendingMarket()
			e := tx.ValidateBasic()

			if tt.expectPass {
				if e != nil {
					t.Errorf("Error: %v", e.Error())
				}
			} else {
				if e == nil {
					t.Errorf("Error: %v", e.Error())
				}
			}

			if !reflect.DeepEqual(tx, tt.want) {
				t.Errorf("FromTreasuryToLendingMarket() = %v, want %v", tx, tt.want)
			}
		})
	}
}
