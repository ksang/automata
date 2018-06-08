package rrul

import (
	"testing"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		output string
	}{
		{
			`MIGRATED TCP STREAM TEST from (null) (0.0.0.0) port 0 AF_INET to (null) () port 0 AF_INET : demo
Interim result: 18963.02 10^6bits/s over 0.200 seconds ending at 1528437374.250
Interim result: 19186.45 10^6bits/s over 0.200 seconds ending at 1528437374.450
Interim result: 17561.28 10^6bits/s over 0.200 seconds ending at 1528437374.650
Interim result: 16563.53 10^6bits/s over 0.200 seconds ending at 1528437374.850
Interim result: 15017.10 10^6bits/s over 0.200 seconds ending at 1528437375.050
Interim result: 17649.69 10^6bits/s over 0.200 seconds ending at 1528437375.250
Interim result: 16904.74 10^6bits/s over 0.200 seconds ending at 1528437375.450
Interim result: 15180.72 10^6bits/s over 0.200 seconds ending at 1528437375.650
Interim result: 15445.21 10^6bits/s over 0.200 seconds ending at 1528437375.850
Interim result: 15768.22 10^6bits/s over 0.200 seconds ending at 1528437376.050
Interim result: 15817.45 10^6bits/s over 0.200 seconds ending at 1528437376.250
Interim result: 17242.11 10^6bits/s over 0.200 seconds ending at 1528437376.450
Interim result: 15835.75 10^6bits/s over 0.200 seconds ending at 1528437376.650
Interim result: 14944.91 10^6bits/s over 0.200 seconds ending at 1528437376.850
Interim result: 15426.48 10^6bits/s over 0.200 seconds ending at 1528437377.050
Interim result: 11131.38 10^6bits/s over 0.000 seconds ending at 1528437377.051
16493.2`,
		},
		{
			`MIGRATED TCP MAERTS TEST from (null) (0.0.0.0) port 0 AF_INET to (null) () port 0 AF_INET : demo
Interim result: 16221.47 10^6bits/s over 0.200 seconds ending at 1528437377.302
Interim result: 18751.98 10^6bits/s over 0.200 seconds ending at 1528437377.502
Interim result: 18408.76 10^6bits/s over 0.200 seconds ending at 1528437377.702
Interim result: 18068.85 10^6bits/s over 0.200 seconds ending at 1528437377.902
Interim result: 17252.26 10^6bits/s over 0.200 seconds ending at 1528437378.103
Interim result: 17132.44 10^6bits/s over 0.200 seconds ending at 1528437378.303
Interim result: 13603.21 10^6bits/s over 0.200 seconds ending at 1528437378.503
Interim result: 15619.20 10^6bits/s over 0.200 seconds ending at 1528437378.703
Interim result: 15112.73 10^6bits/s over 0.200 seconds ending at 1528437378.903
Interim result: 14304.46 10^6bits/s over 0.200 seconds ending at 1528437379.103
Interim result: 14271.71 10^6bits/s over 0.200 seconds ending at 1528437379.303
Interim result: 14642.54 10^6bits/s over 0.200 seconds ending at 1528437379.503
Interim result: 17142.59 10^6bits/s over 0.200 seconds ending at 1528437379.703
Interim result: 17738.76 10^6bits/s over 0.200 seconds ending at 1528437379.903
Interim result: 15496.79 10^6bits/s over 0.199 seconds ending at 1528437380.102
16247.49`,
		},
		{
			`MIGRATED UDP REQUEST/RESPONSE TEST from (null) (0.0.0.0) port 0 AF_INET to (null) () port 0 AF_INET : demo : first burst 0
Interim result:    0.24 10^6bits/s over 0.200 seconds ending at 1528437380.338
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437380.538
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437380.738
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437380.938
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437381.138
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437381.338
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437381.538
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437381.738
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437381.938
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437382.138
Interim result:    0.24 10^6bits/s over 0.200 seconds ending at 1528437382.338
Interim result:    0.29 10^6bits/s over 0.200 seconds ending at 1528437382.538
Interim result:    0.29 10^6bits/s over 0.200 seconds ending at 1528437382.738
Interim result:    0.25 10^6bits/s over 0.200 seconds ending at 1528437382.938
Interim result:    0.24 10^6bits/s over 0.199 seconds ending at 1528437383.138
   0.25`,
		},
	}
	for _, te := range tests {
		ret, err := MarshalOutput([]byte(te.output))
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
		t.Logf("result: %#v", ret)
	}
}
