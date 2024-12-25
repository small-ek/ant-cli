package template

func RequestCommon() string {
	return "package request\n\ntype IdsRequest struct {\n\tIds []int `json:\"ids\" form:\"ids\" binding:\"required\"` //标识\n}"

}
