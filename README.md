## Randomness Beacons

This project is designed to provide a secure, zero-trust randomness beacon for
use in generating random seeds. There are several SHA256 implementations used
for benchmarking purposes.

## Building and Releasing
Deps:
```
- gox
```

Run:
```
bash make_release.sh
```

## Security design
In order to create a trusted, secure randomness beacon, this project is
designed to calculate the SHA256(merkle_root) in order to provide a trustworthy
seed to randomly shuffle a list. In order to do this, the SHA256 function is
applied 2^33 times to the merkle_root of a predetermined and fixed block in the
Bitcoin blockchain. 

In order to tamper with the beacon, an adversary (or me)
would need to have the computing resources to reliably beat the miners on the
Bitcoin blockchain, while calculating 2^33 more SHA256 hashes. This means that
for an adversary it would be **at least 2^32 times harder than the difficulty
needed to reliably mine a single block** to carry out a successful attack (Bitcoin applies SHA256
twice). The adversary would also be subject to all the normal time constraints
that normal miners are subjected to (needs to be completed in under 10 minutes).

## Known Bugs
The hashes from different implementations don't match up.
