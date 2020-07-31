package main

import "fmt"

type Car struct {
	ID           int
	Make         string
	Model        string
	YearReleased int
}

func (car Car) String() string {
	return fmt.Sprintf(
		"Make:\t\t%q\n"+
			"Model:\t\t%q\n"+
			"Released:\t%v\n", car.Make, car.Model, car.YearReleased)
}

var cars = []Car{
	Car{
		ID:           1,
		Make:         "Alfa Romeo",
		Model:        "GIULIETTA",
		YearReleased: 2010,
	},
	Car{
		ID:           2,
		Make:         "Alfa Romeo",
		Model:        "4C",
		YearReleased: 2013,
	},
	Car{
		ID:           3,
		Make:         "Alfa Romeo",
		Model:        "GIULIA",
		YearReleased: 2015,
	},
	Car{
		ID:           4,
		Make:         "Alfa Romeo",
		Model:        "STELVIO",
		YearReleased: 2016,
	},
	Car{
		ID:           5,
		Make:         "Maserati",
		Model:        "Ghibli",
		YearReleased: 2013,
	},
	Car{
		ID:           6,
		Make:         "Maserati",
		Model:        "Levante",
		YearReleased: 2018,
	},
	Car{
		ID:           7,
		Make:         "Ferrari",
		Model:        "812 Superfast",
		YearReleased: 2019,
	},
	Car{
		ID:           8,
		Make:         "Ferrari",
		Model:        "SF90 Stradale",
		YearReleased: 2018,
	},
	Car{
		ID:           9,
		Make:         "Ferrari",
		Model:        "GTC4Lusso",
		YearReleased: 2019,
	},
	Car{
		ID:           10,
		Make:         "Ferrari",
		Model:        "Ferrari 488 Pista",
		YearReleased: 2020,
	},
}
