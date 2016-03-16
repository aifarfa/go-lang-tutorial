package main

import (
  "testing"
)

func TestCreateUser(test *testing.T) {
  usr := CreateUser()
  if(usr.Age != 0){
    test.Error("expect Age", usr.Age, "to be", 0 )
  }
}
