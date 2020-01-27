package services

import (
	"log"
	"testing"
)

func TestRoleAccessService_GetUserAccess(t *testing.T) {
	res := InitRoleAccessService().GetUserAccess(1, "RestoSave")

	log.Println(res)
}
