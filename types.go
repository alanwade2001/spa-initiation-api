package main

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
type GroupHeader struct {
	MessageID            string
	NumberOfTransactions int
	ControlSum           int
	CreationDateTime     string
	InitiatingPartyID    string
}

// PaymentInstruction s
type PaymentInstruction struct {
	PaymentID            string
	NumberOfTransactions int
	ControlSum           int
	RequestedExcutionDate string
	Debtor               Account
}

// PaymentInstructions s
type PaymentInstructions []PaymentInstruction

// WorkflowState s
type WorkflowState struct {
	User  User
	State string
}

// WorkflowStates a
type WorkflowStates []WorkflowState

// Account s
type Account struct {
	IBAN string
	BIC  string
	Name string
}
