# go-boggle

A [Boggle](https://boardgamegeek.com/boardgame/1293/boggle) solver in Go.

## Usage
To solve the Boggle grid [boggle-example.txt](./assets/boggle-example.txt) using
the word list [scrabble.fr.txt](./assets/scrabble.fr.txt):
```bash
make build
./cmd/boggle/boggle -w ./assets/scrabble.fr.txt -b ./assets/boggle-example.txt
```

Check help:
```bash
./cmd/boggle/boggle -h
```

To remove the executable:
```bash
make clean
```
