# Assignment | Week 2

> Book listing and searching app

The app reads a list of top-selling books of all time. User can list all the books in the list and also can search if a book is on the list or not by its name.

## Features

- All of the books on the list can be printed.
- A certain book can be searched on the list.

## Configuration

The app can execute two different commands given by the user on the terminal.

### list Arguement

With the "list" arguement, the user can print all the books on the top-selling books list. Example usage can be seen below.

```
go run main.go list
```

### search Arguement

With the "search" arguement, the user can search if a certain book of his/her choice is on the list or not. After search arguement, the name of the book to be searched should be written. If this book is on the list. The app prints its name. If it is not, the app gives error. Example usage can be seen below.

```
go run main.go search <bookName>
```

## Links

- Project repository: https://github.com/Picus-Security-Golang-Bootcamp/homework-1-week-2-cagrikilicoglu
