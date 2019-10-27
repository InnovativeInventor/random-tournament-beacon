## Commitment

This piece of code will be run to determine the tournament parings:
```python
import random

hash = HASH
matches = list(range(18))
random.seed(hash)
random.shuffle(matches)

print(matches)
```
The nth person in the tournament list will be the `matches[n]` in the shuffled
list.

The first half of the shuffled list will be the first bracket of the
tournament, and the second half of the list will be the second bracket of the
tournament.

The hash used will be the merkle root of Bitcoin block 601250 and path to the
executable will be:
```
build/std_darwin_amd64
```
with commit hash:
```
fd92606acf2342311091a1dace539dd0174198e9
```
