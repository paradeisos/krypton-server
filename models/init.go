package models

var (
	mongo *Model
)

func SetupModel(model *Model) {
	mongo = model
}

func Model() *Model {
	return mongo
}
