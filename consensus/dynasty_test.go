package consensus

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDynasty_NewDynasty(t *testing.T) {
	dynasty := NewDynasty()
	assert.Empty(t,dynasty.miners)
}

func TestDynasty_NewDynastyWithMiners(t *testing.T) {
	tests := []struct{
		name 		string
		input 		[]string
		expected	[]string
	}{
		{
			name: 		"NonEmptyInput",
			input:		[]string{"m1","m2","m3"},
			expected:	[]string{"m1","m2","m3"},
		},
		{
			name: 		"EmptyInput",
			input:		[]string{},
			expected:	[]string{},
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			dynasty:= NewDynastyWithMiners(tt.input)
			assert.Equal(t, tt.expected, dynasty.miners)
		})
	}
}

func TestDynasty_AddMiner(t *testing.T) {
	tests := []struct{
		name 		string
		input 		string
		expected	string
	}{
		{
			name: 		"NonEmptyInput",
			input:		"m1",
			expected:	"m1",
		},
		{
			name: 		"EmptyInput",
			input:		"",
			expected:	"",
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			dynasty:= NewDynasty()
			dynasty.AddMiner(tt.input)
			assert.Equal(t, tt.expected, dynasty.miners[len(dynasty.miners)-1])
		})
	}
}

func TestDynasty_AddMultipleMiners(t *testing.T) {
	tests := []struct{
		name 		string
		input 		[]string
		expected	[]string
	}{
		{
			name: 		"NonEmptyInput",
			input:		[]string{"m1","m2","m3"},
			expected:	[]string{"m1","m2","m3"},
		},
		{
			name: 		"EmptyInput",
			input:		[]string{},
			expected:	nil,
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			dynasty:= NewDynasty()
			dynasty.AddMultipleMiners(tt.input)
			assert.Equal(t, tt.expected, dynasty.miners)
		})
	}
}

func TestDynasty_GetMinerIndex(t *testing.T) {
	tests := []struct{
		name 			string
		initialMiners 	[]string
		miner 			string
		expected		int
	}{
		{
			name: 			"minerCouldBeFound",
			initialMiners:	[]string{"m1","m2","m3"},
			miner: 			"m1",
			expected:		0,
		},
		{
			name: 			"minerCouldNotBeFound",
			initialMiners:	[]string{"m1","m2","m3"},
			miner: 			"m4",
			expected:		-1,
		},
		{
			name: 			"EmptyInput",
			initialMiners:	[]string{"m1","m2","m3"},
			miner: 			"",
			expected:		-1,
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			dynasty:= NewDynastyWithMiners(tt.initialMiners)
			index := dynasty.GetMinerIndex(tt.miner)
			assert.Equal(t, tt.expected, index)
		})
	}
}

func TestDynasty_GetNextMintTimeByIndex(t *testing.T) {
	tests := []struct{
		name 		string
		index 		int
		now 		int64
		expected	int64
	}{
		{
			name: 		"NextMintTimeInThisDynasty",
			index:		2,
			now: 		46,
			expected:	75,
		},
		{
			name: 		"NextMintTimeInNextDynasty",
			index:		1,
			now: 		61,
			expected:	105,
		},
		{
			name: 		"InvalidIndexInput",
			index:		-6,
			now: 		61,
			expected:	-1,
		},
		{
			name: 		"InvalidNowInput",
			index:		2,
			now: 		-1,
			expected:	-1,
		},
		{
			name: 		"IndexInputExceedsMaxSize",
			index:		5,
			now: 		44,
			expected:	-1,
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			dynasty:= NewDynasty()
			nextMintTime := dynasty.GetNextMintTimeByIndex(tt.index, tt.now)
			assert.Equal(t, tt.expected, nextMintTime)
		})
	}
}

func TestDynasty_GetNextMintTime(t *testing.T) {
	tests := []struct{
		name 			string
		initialMiners 	[]string
		miner 			string
		index 			int
		now 			int64
		expected		int64
	}{
		{
			name: 			"NextMintTimeInThisDynasty",
			initialMiners:	[]string{"m1","m2","m3"},
			miner: 			"m3",
			now: 			46,
			expected:		75,
		},
		{
			name: 			"NextMintTimeInNextDynasty",
			initialMiners:	[]string{"m1","m2","m3"},
			miner: 			"m2",
			now: 			61,
			expected:		105,
		},
		{
			name: 			"EmptyInput",
			initialMiners:	[]string{"m1","m2","m3"},
			miner: 			"",
			now: 			61,
			expected:		-1,
		},
		{
			name: 			"InvalidNowInput",
			initialMiners:	[]string{"m1","m2","m3"},
			miner: 			"m2",
			now: 			0,
			expected:		-1,
		},
		{
			name: 			"minerNotFoundInDynasty",
			initialMiners:	[]string{"m1","m2","m3"},
			miner: 			"m5",
			now: 			90,
			expected:		-1,
		},
	}

	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			dynasty:= NewDynastyWithMiners(tt.initialMiners)
			nextMintTime := dynasty.GetNextMintTime(tt.miner, tt.now)
			assert.Equal(t, tt.expected, nextMintTime)
		})
	}
}