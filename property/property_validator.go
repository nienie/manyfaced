package property

//ChangedValidator ...
type ChangedValidator interface {

	//Validate ...
	Validate(newValue string) error
}
