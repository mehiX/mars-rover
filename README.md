# Trivial Mars Rover simulation

This is an example for the implementation of the Command pattern.

It simulates a remote control that would send commands to a Rover. The commands come in the form of ASCII characters:
- F -> move forward 1 unit
- B -> move backwards 1 unit
- L -> rotate left by 90deg
- R -> rotate right by 90deg

## Build the project

```bash
git clone https://github.com/mehiX/mars-rover.git
cd mars-rover

go build ./cmd/roverctl
```

## Run

```bash
# Check available flags
./roverctl -h

# Example run
./roverctl -x 4 -y 5 -d N -c FFFRBFLBRFF
```

## TODO

Ideas to extend the program:
- add units to each command. The commands string should look like: 12F3BLR. Later do the same for L and R
- allow for graceful shutdown by canceling the context on CTRL-C
- create an interactive version of the command that would allow to input commands and receive position updates

