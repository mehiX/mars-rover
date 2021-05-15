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

# Run a series of commands
./roverctl -n Kata -x 4 -y 5 -d N -c FFFRBFLBRFF

# Run a series of commands with a delay of 2s between commands
./roverctl -n Kata -x 4 -y 5 -d N -c FFFRBFLBRFF -delay 2s

# Run in interactive mode without delay
./roverctl -n Kata -x 4 -y 5 -d N

# Run in interactive mode with a delay of 1s between commands
./roverctl -n Kata -x 4 -y 5 -d N -delay 1s
```

## TODO

Ideas to extend the program:
- add units to each command. The commands string should look like: 12F3BLR. Later do the same for L and R

