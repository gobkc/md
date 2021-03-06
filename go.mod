module github.com/gobkc/md

go 1.14

require (
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/gobkc/cmd-parse v0.0.5
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace gopkg.in/yaml.v2 v2.3.0 => github.com/go-yaml/yaml v0.0.0-20200506230838-0b1645d91e85
