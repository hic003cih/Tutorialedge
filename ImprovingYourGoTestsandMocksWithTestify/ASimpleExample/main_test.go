/* package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}
*/

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
}
func TestStatusNotDown(t *testing.T) {
	assert.NotEqual(t, status, "down")
}
