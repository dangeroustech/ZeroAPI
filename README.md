# ZeroAPI

Just like the official one, but for selfhosted controllers too :+1:

## Endpoints TODO

### Node (Pretty sure this will only be local to the node this binary is on)

/node

- [x] GET Node Info

/node/***

- [ ] GET Node Network Memberships
- [ ] GET Node's Moons
- [ ] GET Node's Peers
- [ ] PUT Join Network
- [ ] PUT Update Network Features (global, managed, dns, etc)

### Network (need to check if some of these are 100% viable)

/network

- [ ] GET Network Members
- [ ] POST Create Network
- [ ] DELETE Remove Network

/network/***

- [ ] GET Network Node Info
- [ ] GET Network Routes
- [ ] PUT Approve Node
- [ ] PUT Modify Node Properties (Bridge Allowance, etc.)
- [ ] PUT Add Network Route
- [ ] DELETE Remove Node
