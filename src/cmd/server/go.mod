module github.com/DrDowellsHead/Go_Bootcamp_4/src/cmd/server

go 1.23.0

require (
	github.com/DrDowellsHead/Go_Bootcamp_4/src/di v0.0.0
	go.uber.org/fx v1.24.0
)

replace (
	github.com/DrDowellsHead/Go_Bootcamp_4/src/datasource => ../../datasource
	github.com/DrDowellsHead/Go_Bootcamp_4/src/domain => ../../domain
	github.com/DrDowellsHead/Go_Bootcamp_4/src/web => ../../web
	github.com/DrDowellsHead/Go_Bootcamp_4/src/di => ../../di
)
