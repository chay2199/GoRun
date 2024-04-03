A bit about strings

When we calculate length of a string in Go then it is actually the length of the byte array. So there are some chars that can take more than 1 byte to store a char. In those cases the length might not come as expected. For example the length of énum will be 5 and not 4.