// Code generated by goctl. DO NOT EDIT.
package types

type RequestEmpty struct {
	Name string `json:"name,optional"`
}

type Request struct {
	Name      string `json:"name,optional"`
	NameSpace string `json:"name_space,optional"`
}

type RequestCreateDeployment struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Replicas  int32  `json:"replicas"`
	Image     string `json:"image"`
	Labels    string `json:"labels"`
	Ports     string `json:"ports"`
}

type RequestCreateService struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	AppName   string `json:"appName"`
	Labels    string `json:"labels"`
	Ports     string `json:"ports"`
}

type RequestCreateIngress struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	Host      string `json:"host"`
}

type Response struct {
	Message string `json:"message"`
}
