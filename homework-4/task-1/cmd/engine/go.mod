module go-core.course/homework-3/task-1/cmd/engine

go 1.15

replace go-core.course/homework-4/task-1/pkg/scan v0.0.0 => ../../pkg/scan
require go-core.course/homework-4/task-1/pkg/scan v0.0.0

replace go-core.course/homework-4/task-1/pkg/spider v0.0.0 => ../../pkg/spider
require go-core.course/homework-4/task-1/pkg/spider v0.0.0 

replace go-core.course/homework-4/task-1/pkg/dummyScan v0.0.0 => ../../pkg/dummyScan
require go-core.course/homework-4/task-1/pkg/dummyScan v0.0.0 

replace go-core.course/homework-4/task-1/pkg/index v0.0.0 => ../../pkg/index
require go-core.course/homework-4/task-1/pkg/index v0.0.0 
