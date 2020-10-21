module go-core.course/homework-3/task-1/cmd/engine

go 1.15

replace go-core.course/homework-3/task-1/pkg/spider v0.0.0 => ../../pkg/spider

replace go-core.course/homework-3/task-1/pkg/dummySpider v0.0.0 => ../../pkg/dummySpider

require go-core.course/homework-3/task-1/pkg/spider v0.0.0 

require go-core.course/homework-3/task-1/pkg/dummySpider v0.0.0 
