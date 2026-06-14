package main

import "errors"

var (
	ErrDivideByZero = errors.New("tidak bisa membagi dengan 0")
	ErrNegativeInput = errors.New("input tidak boleh negatif")
	ErrEmptyName    = errors.New("nama tidak boleh kosong")
	ErrNegativeAge  = errors.New("umur tidak boleh negatif")
)
