// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

func TestCheckServerlessV2MaxCapacity(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		maxCapacity float64
		support     *types.ServerlessV2FeaturesSupport
		wantErr     bool
	}{
		"no support info":      {maxCapacity: 384},
		"nil max capacity":     {maxCapacity: 384, support: &types.ServerlessV2FeaturesSupport{}},
		"within limit":         {maxCapacity: 300, support: &types.ServerlessV2FeaturesSupport{MaxCapacity: aws.Float64(384)}},
		"at limit":             {maxCapacity: 384, support: &types.ServerlessV2FeaturesSupport{MaxCapacity: aws.Float64(384)}},
		"above 256 limit":      {maxCapacity: 300, support: &types.ServerlessV2FeaturesSupport{MaxCapacity: aws.Float64(256)}, wantErr: true},
		"above 384 limit":      {maxCapacity: 384.5, support: &types.ServerlessV2FeaturesSupport{MaxCapacity: aws.Float64(384)}, wantErr: true},
		"256 on 256 engine ok": {maxCapacity: 256, support: &types.ServerlessV2FeaturesSupport{MaxCapacity: aws.Float64(256)}},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := checkServerlessV2MaxCapacity(tc.maxCapacity, "aurora-postgresql", "17.9", tc.support)
			if (err != nil) != tc.wantErr {
				t.Fatalf("got err = %v, wantErr = %t", err, tc.wantErr)
			}
		})
	}
}
