### Features:

1. List all the books.
2. Get book by ID.


### Get all books:

* Response: 
```
{
    "status": 200,
    "data": [
        {
            "id": 1,
            "author": {
                "firstname": "Jorge Luis",
                "lastname": "Borges"
            },
            "title": "El Aleph",
            "price": 42.75,
            "isbn": "84-206-1933-7",
            "stock": 0
        },
        {
            "id": 2,
            "author": {
                "firstname": "Stanislaw",
                "lastname": "Lem"
            },
            "title": "Solaris",
            "price": 65.2,
            "isbn": "0156027607",
            "stock": 15
        }
    ]
}
```

### Get book By ID:

* Response: 
```
{
    "status": 200,
    "data": {
        "id": 1,
        "author": {
            "firstname": "Isaac",
            "lastname": "Asimov"
        },
        "title": "Fundation",
        "price": 28.5,
        "isbn": "0-553-29335-4",
        "stock": 9
    }
}
```