package property

//PropertyChangedValidator ...
type PropertyChangedValidator interface {

    //Validate ...
    Validate(newValue string) error
}