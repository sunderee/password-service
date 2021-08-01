package tests

import (
	"fmt"
	"testing"

	"github.com/sunderee/password-service/service"
)

func TestPasswordGeneration(test *testing.T) {
	var vanillaOptions *service.Options = new(service.Options)
	vanillaOptions.HasUppercase = true
	vanillaOptions.HasNumbers = false
	vanillaOptions.HasSpecials = false
	vanillaOptions.PasswordLength = 8

	var vanillaPassword string = service.GeneratePassword(*vanillaOptions)
	fmt.Println("Vanilla password: " + vanillaPassword)
	if len(vanillaPassword) != 8 {
		test.Fail()
	}

	var strongOptions *service.Options = new(service.Options)
	strongOptions.HasUppercase = true
	strongOptions.HasNumbers = true
	strongOptions.HasSpecials = true
	strongOptions.PasswordLength = 15

	var strongPassword string = service.GeneratePassword(*strongOptions)
	fmt.Println("String password: " + strongPassword)
	if len(strongPassword) != 15 {
		test.Fail()
	}
}
