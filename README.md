# Trivial Mars Rover simulation

This is an example for the implementation of the Command pattern

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

