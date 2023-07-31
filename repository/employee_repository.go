package repository

import "demo-unit-test/entity"

/**
 * Created by muhammad.khadafi on 28/07/2023
 */

type EmployeeRepository interface {
	FindById(id int) *entity.Employee
}
