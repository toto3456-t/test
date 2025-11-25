package test
import "gorm.io/gorm"

type Test struct{
	gorm.Model
	test string
}
