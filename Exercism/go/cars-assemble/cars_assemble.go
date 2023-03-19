package cars

const oneCarCost = 10000
const tenCarsCost = 95000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) / 100 * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var succesfulCars float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(succesfulCars / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var groupsOfTen int = carsCount / 10
	var remainder int = carsCount % 10
	return (uint)(groupsOfTen*tenCarsCost + remainder*oneCarCost)
}
