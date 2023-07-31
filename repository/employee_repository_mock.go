package repository

import (
	"demo-unit-test/entity"
	"github.com/stretchr/testify/mock"
)

/**
 * Created by muhammad.khadafi on 28/07/2023
 */

type EmployeeRepositoryMock struct {
	Mock mock.Mock
}

func (repository *EmployeeRepositoryMock) FindById(id int) *entity.Employee {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		employee := arguments.Get(0).(entity.Employee)
		return &employee
	}
}
