package service

import (
	"demo-unit-test/entity"
	"demo-unit-test/repository"
	"errors"
)

/**
 * Created by muhammad.khadafi on 28/07/2023
 */

type EmployeeService struct {
	EmployeeRepository repository.EmployeeRepository
}

func CalculateTakeHomePay(basicSalary int, allowance int) int {
	return basicSalary + allowance
}

func GetEmployeeFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func (service EmployeeService) GetEmployee(id int) (*entity.Employee, error) {
	employee := service.EmployeeRepository.FindById(id)
	if employee == nil {
		return nil, errors.New("Employee not found")
	} else {
		return employee, nil
	}
}
