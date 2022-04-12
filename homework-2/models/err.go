package models

import (
	"errors"
)

var (
	NotInbooks = errors.New("Given value not in Books..")

	HigherThanStock = errors.New("Given count is higher than stock number..")

	ErrList = errors.New("List command does not take arg..")

	ExpectedDef = errors.New("Expected 'buy','delete','search' or 'list'..")

	ExpectedBuy      = errors.New("Expected valid buy arguments for command 'buy' :> 'buy' 'int' 'int'..")
	ExpectedBuyArg   = errors.New("Expected valid buy arguments for command 'buy'..")
	ExpectedValidBuy = errors.New("Expected valid buy argument(id) for command 'buy', Entered value not in Books..")

	ExpectedSearchArg = errors.New("Expected search argument for command 'search'..")

	ExpectedDelete      = errors.New("Expected valid delete argument for command 'delete' :> 'delete' 'int'..")
	ExpectedDeleteArg   = errors.New("Expected valid delete argument for command 'delete'..")
	ExpectedValidDelete = errors.New("Expected valid delete argument for command 'delete', Entered value not in Books..")
)
