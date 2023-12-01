package mongo

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/HamzaGo5911/wanclouds-employee-hub/db"
	"github.com/HamzaGo5911/wanclouds-employee-hub/models"
)

func Test_client_AddEmployee(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	type args struct {
		employee *models.Employee
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add employee in db",
			args: args{employee: &models.Employee{
				ID:        "123",
				Name:      "hamza",
				Email:     "hamza@wanlouds.net",
				Phone:     "12345567",
				Team:      "Golang",
				JobTitle:  "Software Engineer",
				StartDate: time.Now(),
				Salary:    75000.00,
				Benefits:  []string{"Health Insurance", "Paid Time Off"},
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewClient(db.Option{})
			_, err := m.AddEmployee(tt.args.employee)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddCloudProvider() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func Test_client_GetEmployeeByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})

	employee := &models.Employee{
		ID:        "123",
		Name:      "hamza",
		Email:     "hamza@wanlouds.net",
		Phone:     "12345567",
		Team:      "Golang",
		JobTitle:  "Software Engineer",
		StartDate: time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		Salary:    75000.00,
		Benefits:  []string{"Health Insurance", "Paid Time Off"},
	}

	_, _err := c.AddEmployee(employee)
	if _err != nil {
		t.Fatalf("Error adding employee: %v", _err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Employee
		wantErr bool
	}{
		{
			name:    "success - get employee from db",
			args:    args{id: employee.ID},
			want:    employee,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetEmployeeByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEmployeeByID error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEmployeeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_DeleteEmployee(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})
	employee := &models.Employee{
		ID:        "123",
		Name:      "hamza",
		Email:     "hamza@wanlouds.net",
		Phone:     "12345567",
		Team:      "Golang",
		JobTitle:  "Software Engineer",
		StartDate: time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
		Salary:    75000.00,
		Benefits:  []string{"Health Insurance", "Paid Time Off"},
	}

	_, _ = c.AddEmployee(employee)
	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete employee from db",
			args:    args{id: employee.ID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteEmployee(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_ListEmployee(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success - List employee from db",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.ListEmployee()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_AddOffice(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	type args struct {
		office *models.Office
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add office in db",
			args: args{office: &models.Office{
				ID:          "456",
				EmployeeID:  "789",
				OfficeRoom:  "Room 405",
				MeetingRoom: "Room 206",
				LunchArea:   "gallery",
				TeaArea:     "Kitchen",
				PlayingArea: "Recreation Room",
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewClient(db.Option{})
			_, err := m.AddOffice(tt.args.office)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddOffice() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func Test_client_OfficeByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})

	office := &models.Office{
		ID:          "456",
		EmployeeID:  "789",
		OfficeRoom:  "Room 405",
		MeetingRoom: "Room 206",
		LunchArea:   "gallery",
		TeaArea:     "Kitchen",
		PlayingArea: "Recreation Room",
	}

	_, _err := c.AddOffice(office)
	if _err != nil {
		t.Fatalf("Error adding office: %v", _err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Office
		wantErr bool
	}{
		{
			name:    "success - get office from db",
			args:    args{id: office.ID},
			want:    office,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetOfficeByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOfficeByID error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOfficeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_DeleteOffice(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})
	office := &models.Office{
		ID:          "456",
		EmployeeID:  "789",
		OfficeRoom:  "Room 405",
		MeetingRoom: "Room 206",
		LunchArea:   "gallery",
		TeaArea:     "Kitchen",
		PlayingArea: "Recreation Room",
	}

	_, _ = c.AddOffice(office)
	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete office from db",
			args:    args{id: office.ID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteOffice(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteOffice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_ListOffice(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success - List office from db",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.ListOffice()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListOffice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_AddHome(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	type args struct {
		home *models.Home
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - add home in db",
			args: args{home: &models.Home{
				ID:            "1",
				EmployeeID:    "123",
				ResidencyRoom: "Room A",
				Dinner:        "Home-cooked",
				DinnerRoom:    "Kitchen",
				DiningRoom:    "Dining Hall",
				CommonRoom:    "Living Room",
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewClient(db.Option{})
			_, err := m.AddHome(tt.args.home)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddHome() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func Test_client_GetHomeByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})

	home := &models.Home{
		ID:            "1",
		EmployeeID:    "123",
		ResidencyRoom: "Room A",
		Dinner:        "Home-cooked",
		DinnerRoom:    "Kitchen",
		DiningRoom:    "Dining Hall",
		CommonRoom:    "Living Room",
	}

	_, _err := c.AddHome(home)
	if _err != nil {
		t.Fatalf("Error adding home: %v", _err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Home
		wantErr bool
	}{
		{
			name:    "success - get home from db",
			args:    args{id: home.ID},
			want:    home,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetHomeByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHomeByID error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHomeByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_DeleteHome(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})
	home := &models.Home{
		ID:            "1",
		EmployeeID:    "123",
		ResidencyRoom: "Room A",
		Dinner:        "Home-cooked",
		DinnerRoom:    "Kitchen",
		DiningRoom:    "Dining Hall",
		CommonRoom:    "Living Room",
	}

	_, _ = c.AddHome(home)
	type args struct {
		id string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete home from db",
			args:    args{id: home.ID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteHome(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteHome() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_ListHome(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	c, _ := NewClient(db.Option{})
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success - List home from db",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.ListHome()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListHome() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
