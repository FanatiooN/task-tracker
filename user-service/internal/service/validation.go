package service

import "regexp"

var validEmail = regexp.MustCompile(`^[A-Za-z]+[@][A-Za-z]+[.][A-Za-z]{1,3}$`)
