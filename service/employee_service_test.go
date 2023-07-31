package service

import (
	"demo-unit-test/entity"
	"demo-unit-test/repository"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

/**
 * Created by muhammad.khadafi on 28/07/2023
 */

var employeeRepository = &repository.EmployeeRepositoryMock{Mock: mock.Mock{}}
var employeeService = EmployeeService{EmployeeRepository: employeeRepository}

func TestCalculateTakeHomePay(t *testing.T) {
	result := CalculateTakeHomePay(5000000, 2000000)
	assert.Equal(t, 7000000, result, "Result must be 7000000")
	fmt.Println("line from TestCalculateTakeHomePay executed")
}

func TestGetEmployeeFullName(t *testing.T) {
	result := GetEmployeeFullName("Budi", "Santoso")
	require.Equal(t, "Budi Santoso", result, "Result must be Budi Santoso")
	fmt.Println("line from TestGetEmployeeFullName executed")
}

func TestGetEmployee_GetNotFound(t *testing.T) {

	// program mock
	employeeRepository.Mock.On("FindById", 1).Return(nil)

	category, err := employeeService.GetEmployee(1)
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestGetEmployee_GetSuccess(t *testing.T) {
	employee := entity.Employee{
		Id:   2,
		Nama: "Jaja",
	}
	employeeRepository.Mock.On("FindById", 2).Return(employee)

	result, err := employeeService.GetEmployee(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, employee.Id, result.Id)
	assert.Equal(t, employee.Nama, result.Nama)
}
