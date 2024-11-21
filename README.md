# Spotfinder

An API to display locations/spots of different institutions, places in Sierra Leone

### Tasks to complete
- [X] Display all locations
- [X] Add new locations
- [ ] Filter locations
  - [ ] By category
  - [ ] By tags _(not added yet)_(*)
- [ ] Add user authentications to update locations(*)
- [ ] Mark location as valid/invalid - then delete if invalid(*)
- [ ] Add a flag to seed data on startup.

> tasks with * are undecided yet

### Requirements
Tools required to contribute or build the project locally.

- Go - the go programming language [website](https://go.dev) _(basic knowledge is ok)_
- Make - comes installed in UNIX system _(optional)_
- SQLC - generates SQL boilerplate _(required if you plan to modify SQL statements)_ [website](https://sqlc.dev)
- Goose - migration tool [website](https://pressly.github.io/goose/)

> There are binary files for all platforms in the release page. Use them if you don't plan to contribute but want to run the final API locally.
> The database is SQLite for easy access to run and test for non-go developers and locally.
> for windows users without `make` installed, just type `go run ./cmd/api` and the project will be up and running.
