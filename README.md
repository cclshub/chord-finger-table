# chord-finger-table

#### Things that have been implemented as of 31/10/24:

**Assumption**
For this checkoff, we are going to assume that each table somehow has a copy of all the nodes in the system (this assumption is required to populate the finger tables as the nodes are "spawned" simultaneously and we do not have a mechanism to handle node entering and leaving the system). 


functions:
1. initNode in a non-distributed way
2. populateKeys (without predecessor)
3. populatefTable (which requires a slice of all the nodes in the system)
4. main
5. populateFingerTable(n)
6. n.findSuccessor(k) (this needs to be API-fied)
7. closestPrecedingNode (don't need to touch it)

#### Things to implement
- [ ] initNode in a distributed way
- [ ] add predecessor to populateKeys
- [ ] API-fy n.findSuccessor(k)
- [ ] integrate this with sean's grpc code
- [ ] prettify the output on the console
- [ ] add more comments to the functions