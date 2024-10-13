package kaivestbinance

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	client := NewKaivestBinanceClient("http://localhost:8027")

	resPairsWithUsdt, err := client.GetPairsWithUsdt()
	require.NoError(t, err)
	fmt.Println("resPairsWithUsdt", resPairsWithUsdt)

	resTicker, err := client.GetSpotBookTicker(strings.Join(resPairsWithUsdt[:100], ","))
	require.NoError(t, err)
	fmt.Println("resTicker", resTicker)
}
