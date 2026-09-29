package smtp

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnhancedCodeToPart(t *testing.T) {
	testCases := []struct {
		code EnhancedCode
		want string
	}{
		{EnhancedCode{2, 0, 0}, "2.0.0 "},
		{EnhancedCode{5, 1, 1}, "5.1.1 "},
		{EnhancedCode{5, 7, 23}, "5.7.23 "},
		{EnhancedCode{4, 7, 23}, "4.7.23 "},
		{EnhancedCode{5, 999, 999}, "5.999.999 "},
		{NoEnhancedCode, ""},
	}
	for _, tc := range testCases {
		t.Run(tc.want, func(t *testing.T) {
			assert.Equal(t, tc.want, string(tc.code.ToPart()))
		})
	}
}

func TestStatusWriteToMultiDigitEnhancedCode(t *testing.T) {
	var buf bytes.Buffer
	_, err := NewStatusM(550, EnhancedCode{5, 7, 23}, []string{"SPF", "fail"}).WriteTo(&buf)
	require.NoError(t, err)
	assert.Equal(t, "550-5.7.23 SPF\r\n550 5.7.23 fail\r\n", buf.String())
}
