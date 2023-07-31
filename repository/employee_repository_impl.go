package repository

import "demo-unit-test/entity"

/**
 * Created by muhammad.khadafi on 28/07/2023
 */

type EmployeeRepositoryImpl struct {
	// DB object, etc..
}

func (repository *EmployeeRepositoryImpl) FindById(id int) *entity.Employee {
	// real implementation..
	panic("implement me")
}
