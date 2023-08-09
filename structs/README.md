In golang, there are two types we need to worry about regarding handling of pointers

Value Types and Reference Types

### For Value Types such as :

1. int
2. float
3. string
4. bool
5. struct

we will need to use pointers to handle manipulation of data in memory

### For Refernce Types such as:

1. slices
2. maps
3. channels
4. pointers
5. functions

pointers will not be needed since this will be handle by golang itself since they already have a pointer as part of their references.
