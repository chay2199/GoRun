Unlike strings in Go where s1 = s2 will just add a discriptor, in array a1 = a2 will actually copy the memory chunks.

Slices in Go have variable length, backed by some array. Slices are passed by reference. Slices are not comparable and they cannot be used as a map key.
