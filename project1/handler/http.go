package handler

import "challenge/project1/service"

type HttpServer struct {
	service service.ServiceInterface
}

func NewHttpServer(service service.ServiceInterface) HttpServer {
	return HttpServer{service: service}
}
