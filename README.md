# Simple Library App
A Simple Library App to retrieve list of book and schedule a pickup time for borrow a book.

### Feature
- List Books: Allows users to retrieve a list of books available in the library.
- List Pickup Schedules: Allows users to retrieve a list of pickup schedules.
- Create Pickup Schedule: Allows users to schedule a pickup time to borrow a book.

## Getting Started

### Prerequisite
- Go: Version 1.20 or higher

### Installation
- Install dependency

```bash
go mod tidy
```

- Copy env.sample

```bash
cp env.sample .env
```

#### Running in Local

- Start the application

```bash
go run cmd/main.go
```

## Usage

### API Endpoint

| Endpoint | Method    | Description    |
| :---:   | :---: | :---: |
| `/api/books` | `GET` | Retrieve list of book |
| `/api/pickup-schedule` | `GET` | Retrieve list of pickup schedule |
| `/api/pickup-schedule/create` | `POST` | Create a pickup schedule |

#### List Books

##### Example Request

```bash
curl --location --request GET 'localhost:8080/api/books?subjects=Fiction' \
--header 'Content-Type: application/json'
```

##### Example Response

```bash
{
  "data": [
    {
        "Title": "Wuthering Heights",
        "Authors": [
            "Emily Brontë"
        ],
        "EditionNumber": "9798373104548",
        "IsAvailable": true
    },
    {
        "Title": "Chronicles of Avonlea",
        "Authors": [
            "Lucy Maud Montgomery"
        ],
        "EditionNumber": "9781511627740",
        "IsAvailable": true
    },
    {
        "Title": "Le Petit Prince",
        "Authors": [
            "Antoine de Saint-Exupéry"
        ],
        "EditionNumber": "8459912019",
        "IsAvailable": true
    }
  ]
}
```

#### List Pickup Schedules

##### Example Request

```bash
curl --location --request GET 'localhost:8080/api/pickup-schedule' \
--header 'Content-Type: application/json'
```

##### Example Response

```bash
{
  "data": [
    {
        "ID": 1,
        "Book": {
            "Title": "Book of Modern Puzzles",
            "Authors": [
                "Gerald L. Kaufman"
            ],
            "EditionNumber": "0486201430",
            "IsAvailable": false
        },
        "DateTime": "2025-01-25T10:00:00Z"
    },
    {
        "ID": 2,
        "Book": {
            "Title": "Harry Potter and the Deathly Hallows",
            "Authors": [
                "J. K. Rowling"
            ],
            "EditionNumber": "9780545139700",
            "IsAvailable": false
        },
        "DateTime": "2025-01-23T06:45:00Z"
    }
  ]
}
```

#### Create Pickup Schedule

##### Example Request

```bash
curl --location 'localhost:8080/api/pickup-schedule/create' \
--header 'Content-Type: application/json' \
--data '{
    "edition_number": "0486201430",
    "datetime": "2025-01-25T10:00:00Z"
}'
```

##### Example Response

```bash
{
  "Schedule": {
    "ID": 1,
    "Book": {
      "Title": "Book of Modern Puzzles",
      "Authors": [
        "Gerald L. Kaufman"
      ],
      "EditionNumber": "0486201430",
      "IsAvailable": false
    },
    "DateTime": "2025-01-25T10:00:00Z"
  },
  "Message": "Pickup schedule created!"
}
```