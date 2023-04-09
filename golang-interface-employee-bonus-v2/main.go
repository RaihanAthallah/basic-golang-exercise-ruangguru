package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (j *Junior) GetBonus() float64 {
	return 1 * float64(j.BaseSalary) * float64(j.WorkingMonth) / 12 // TODO: replace this
}

func (s *Senior) GetBonus() float64 {
	return (2 * float64(s.BaseSalary) * float64(s.WorkingMonth) / 12) + (s.PerformanceRate * float64(s.BaseSalary)) // TODO: replace this
}

func (m *Manager) GetBonus() float64 {
	return (2 * float64(m.BaseSalary) * float64(m.WorkingMonth) / 12) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary)) // TODO: replace this
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	sum := 0.0
	for _, employee := range employees {
		sum += employee.GetBonus()
	}

	return sum // TODO: replace this
}

func main() {

}
