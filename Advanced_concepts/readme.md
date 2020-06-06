A type assertion notation; x.(T). Where x is an inteface type 
and T is the type. The actual value stored in x is of type T
and T must satisfy the interface type of x.
type assertion helps to 
1. check whether an interface value keeps a particular type, 
returning two values: the underlying value and a bool value
that tells you whether the type assertion was succesful or not.