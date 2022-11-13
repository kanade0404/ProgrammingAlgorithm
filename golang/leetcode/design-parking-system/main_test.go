package design_parking_system

import "testing"

func TestParkingSystem(t *testing.T) {
	tests := []struct {
		name      string
		big       int
		medium    int
		small     int
		operation []int
		wants     []bool
	}{
		{
			name:      "test1",
			big:       1,
			medium:    1,
			small:     0,
			operation: []int{1, 2, 3, 1},
			wants:     []bool{true, true, false, false},
		},
	}
	for _, tt := range tests {
		p := Constructor(tt.big, tt.medium, tt.small)
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < len(tt.operation); i++ {
				if got := p.AddCar(tt.operation[i]); got != tt.wants[i] {
					t.Errorf("AddCar(%d)got: %t, want: %t", tt.operation[i], got, tt.wants[i])
				}
			}
		})
	}
}
