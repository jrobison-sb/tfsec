package secrets

import (
	"testing"

	"github.com/aquasecurity/tfsec/internal/pkg/testutil"
)

func Test_AWSSensitiveLocals(t *testing.T) {
	expectedCode := "general-secrets-no-plaintext-exposure"

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode string
		mustExcludeResultCode string
	}{
		{
			name: "check sensitive local with value",
			source: `
 locals {
 	password = "secret"
 }`,
			mustIncludeResultCode: expectedCode,
		},
		{
			name: "check non-sensitive local",
			source: `
 locals {
 	something = "something"
 }`,
			mustExcludeResultCode: expectedCode,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			results := testutil.ScanHCL(test.source, t)
			if test.mustIncludeResultCode != "" {
				testutil.AssertRuleFound(t, test.mustIncludeResultCode, results, "false negative found")
			}
			if test.mustExcludeResultCode != "" {
				testutil.AssertRuleNotFound(t, test.mustExcludeResultCode, results, "false positive found")
			}
		})
	}

}
