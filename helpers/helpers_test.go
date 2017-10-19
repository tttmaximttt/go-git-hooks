package helpers_test

import (
  "testing"
  "github.com/tttmaximttt/hooks/helpers"
)
// TODO
func extractDescription_WithDescription(t *testing.T) {
  rawBranch := "HS-2342-fdsdf-efcbg-efe-gfg"
  expected := "HS-2342"

  actual := helpers.ExtractDescription(rawBranch)

  if actual != expected {
    t.Errorf("incorrect actual result actual: %s", actual)
  }
}

func extractDescription_WithoutDescription(t *testing.T) {
  rawBranch := "HS-2342"

  actual := helpers.ExtractDescription(rawBranch)

  if actual != rawBranch {
    t.Errorf("incorrect actual result actual: %s", actual)
  }
}