package design_parking_system

type ParkingSystem struct {
	bigSlot    int
	bigNum     int
	mediumSlot int
	mediumNum  int
	smallSlot  int
	smallNum   int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big,
		0,
		medium,
		0,
		small,
		0,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 {
		if this.bigNum+1 <= this.bigSlot {
			this.bigNum += 1
			return true
		} else {
			return false
		}
	}
	if carType == 2 {
		if this.mediumNum+1 <= this.mediumSlot {
			this.mediumNum += 1
			return true
		} else {
			return false
		}
	}
	if this.smallNum+1 <= this.smallSlot {
		this.smallNum += 1
		return true
	} else {
		return false
	}
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
