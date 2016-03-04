package utils

type TimeUnit string

/*
const (
	MILLISECONDS TimeUnit = iota
	SECONDS
	MINUTES
	HOURS
	DAYS
	WEEKS
	MONTHS
	YEARS
)

var unitString = []string{
	"miiliseconds",
	"seconds",
	"minutes",
	"hours",
	"days",
	"weeks",
	"months",
	"years",
}
*/

const (
	MILLISECONDS TimeUnit = "milliseconds"
	SECONDS               = "seconds"
	MINUTES               = "minutes"
	HOURS                 = "hours"
	DAYS                  = "days"
	WEEKS                 = "weeks"
	MONTHS                = "months"
	YEARS                 = "years"
)

/*
func (unit TimeUnit) String() string {
	return unitString[unit]
}
*/
