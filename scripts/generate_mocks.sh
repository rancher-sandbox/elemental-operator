#!/bin/sh

go install go.uber.org/mock/mockgen@latest

mockgen -destination=pkg/register/mocks/mock_client.go -package=mocks github.com/rancher/elemental-operator/pkg/register Client
