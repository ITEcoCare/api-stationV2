package database

import (
	"api-station/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBCon() *gorm.DB {

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed Connect Database")
		log.Fatal(err.Error())
		panic(err.Error())
	}

	db.AutoMigrate(
		models.Province{},
		models.City{},
		models.District{},
		models.Village{},
		models.IdentityCard{},
		models.Bank{},
		models.Company{},
		models.Branch{},
		models.Warehouse{},
		models.Team{},
		models.Role{},
		models.JobPosition{},
		models.User{},
		models.Employee{},
		models.EmployeeAddress{},
		models.EmployeeEducation{},
		models.EmployeeContact{},
		models.EmployeeAttendance{},
		models.CustomerGroup{},
		models.Customer{},
		models.CustomerAddress{},
		models.CustomerPic{},
		models.Project{},
		models.ProjectCustomShift{},
		models.ProjectSectionRoom{},
		models.ProjectRequestEmployee{},
		models.ProjectEmployeeAssignment{},
		models.ProjectEmployeeAttendace{},
		models.ProjectEmployeeCheckpoint{},
		models.ProjectTop{},
		models.ProjectInvoice{},
		models.ProjectPaymentReceive{},
		// models.Permission{},
		// models.ModuleApp{},
		// models.ModulePermission{},
	)

	return db
}
