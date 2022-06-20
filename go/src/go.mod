module api

go 1.18

require visdb v0.0.0-00010101000000-000000000000

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	gorm.io/driver/mysql v1.3.4 // indirect
	gorm.io/gorm v1.23.6 // indirect
)

replace visdb => ./utils
