package main

import types "github.com/alanwade2001/spa-common"

// Initiation s
type Initiation struct {
	ID                  string `bson:"_id"`
	Customer            Company
	States              WorkflowStates
	GroupHeader         GroupHeader
	PaymentInstructions PaymentInstructions
}

// User s
type User struct {
	Email string
	Role  string
}

// Company s
type Company struct {
	ID                string
	Name              string
	InitiatingPartyID string
}

// Initiations a
type Initiations []Initiation

// GroupHeader s
type GroupHeader types.GroupHeader

// PaymentInstruction s
type PaymentInstruction types.PaymentInstruction

// PaymentInstructions s
type PaymentInstructions []PaymentInstruction

// WorkflowState s
type WorkflowState struct {
	User  User
	State string
}

// WorkflowStates a
type WorkflowStates []WorkflowState

// AccountReference s
type AccountReference types.AccountReference
