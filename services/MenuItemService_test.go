package services

import (
	"resto-be/models/dto"
	"testing"
)

func TestMenuItemServiceInterface_UploadImage(t *testing.T) {
	req:= dto.UploadImageMenuItemRequestDto{}
	InitializeMenuItemServiceInterface().UploadImage(req)
}
