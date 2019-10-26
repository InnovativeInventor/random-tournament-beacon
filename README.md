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

## Known Bugs
The hashes from different implementations don't match up.
