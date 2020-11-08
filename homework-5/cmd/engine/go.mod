module go-core.course/cmd/engine

go 1.15

replace go-core.course/pkg/scan v0.0.0 => ../../pkg/scan
require go-core.course/pkg/scan v0.0.0

replace go-core.course/pkg/spider v0.0.0 => ../../pkg/spider
require go-core.course/pkg/spider v0.0.0 

replace go-core.course/pkg/dummyScan v0.0.0 => ../../pkg/dummyScan
require go-core.course/pkg/dummyScan v0.0.0 

replace go-core.course/pkg/index v0.0.0 => ../../pkg/index
require go-core.course/pkg/index v0.0.0 
