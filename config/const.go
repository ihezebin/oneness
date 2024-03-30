package config

type DataType string

const (
	DataTypeJson       DataType = "json"
	DataTypeToml       DataType = "toml"
	DataTypeYaml       DataType = "yaml"
	DataTypeYml        DataType = "yml"
	DataTypeProperties DataType = "properties"
	DataTypeProps      DataType = "props"
	DataTypeProp       DataType = "prop"
	DataTypeHcl        DataType = "hcl"
	DataTypeDotenv     DataType = "dotenv"
	DataTypeEnv        DataType = "env"
	DataTypeIni        DataType = "ini"
)

type DestTagName string

const (
	DestTagNameJson   DestTagName = "json"
	DestTagNameConfig DestTagName = "config"
)

const defaultDestTagName DestTagName = DestTagNameJson
const defaultFileName string = "config"
