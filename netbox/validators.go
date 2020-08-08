package netbox

import (
	"net"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func stringLenBetween(min, max int) schema.SchemaValidateDiagFunc {
	return func(i interface{}, k cty.Path) diag.Diagnostics {
		v, ok := i.(string)
		if !ok {
			return diag.Errorf("expected type of %s to be string", k)
		}

		if len(v) < min || len(v) > max {
			return diag.Errorf("expected length of %s to be in the range (%d - %d), got %s", k, min, max, v)
		}

		return nil
	}
}

func stringInSlice(valid []string) schema.SchemaValidateDiagFunc {
	return func(i interface{}, k cty.Path) diag.Diagnostics {
		v, ok := i.(string)
		if !ok {
			return diag.Errorf("expected type of %s to be string", k)
		}

		for _, str := range valid {
			if v == str {
				return nil
			}
		}

		return diag.Errorf("expected %s to be one of %v, got %s", k, valid, v)
	}
}

func isCIDR(i interface{}, k cty.Path) diag.Diagnostics {
	v, ok := i.(string)
	if !ok {
		return diag.Errorf("expected type of %s to be string", k)
	}

	if _, _, err := net.ParseCIDR(v); err != nil {
		return diag.Errorf("expected %q to be a valid IPv4 Value, got %v: %v", k, i, err)
	}

	return nil
}
