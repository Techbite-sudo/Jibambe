# Book Library GraphQL API

This is a GraphQL API for managing a collection of books in a library. It provides queries to fetch books, authors, and genres, as well as mutations to create, update, and delete books.

## Getting Started

To use this API, you'll need a GraphQL client. You can use a tool like Altair GraphQL Client or GraphQL Playground to send queries and mutations to the API.

## Queries

### `books`

This query returns a list of all books in the library.

**Query:**

```graphql
{
  books {
    id
    title
    author
    publishedYear
  }
}
```

**Example Response:**

```json
{
  "data": {
    "books": [
      {
        "id": "ae8a941d-ea10-412d-97bb-409bbc88a8fa",
        "title": "The Pearl",
        "author": "Mike",
        "publishedYear": 2020
      },
      {
        "id": "0b14f521-b319-4884-9247-666d1b6df8af",
        "title": "Blossoms",
        "author": "Boniface",
        "publishedYear": 2014
      }
    ]
  }
}
```

### `book`

This query fetches a specific book by its ID.

**Query:**

```graphql
{
  book(id: "ae8a941d-ea10-412d-97bb-409bbc88a8fa") {
    id
    title
    author
    publishedYear
  }
}
```

**Example Response:**

```json
{
  "data": {
    "book": {
      "id": "ae8a941d-ea10-412d-97bb-409bbc88a8fa",
      "title": "The Pearl",
      "author": "Mike",
      "publishedYear": 2020
    }
  }
}
```

## Mutations

### `createBook`

This mutation creates a new book in the library.

**Mutation:**

```graphql
mutation {
  createBook(input: { title: "Blossoms", author: "Boniface", publishedYear: 2014 }) {
    id
    title
    author
    publishedYear
  }
}
```

**Example Response:**

```json
{
  "data": {
    "createBook": {
      "id": "0b14f521-b319-4884-9247-666d1b6df8af",
      "title": "Blossoms",
      "author": "Boniface",
      "publishedYear": 2014
    }
  }
}
```

### `updateBook`

This mutation updates an existing book by its ID.

**Mutation:**

```graphql
mutation {
  updateBook(
    id: "328545c8-68ba-4701-af8b-2282c1e573f3"
    input: { title: "Blossoms", author: "Boniface", publishedYear: 2014 }
  ) {
    id
    title
    author
    publishedYear
  }
}
```

**Example Response:**

```json
{
  "data": {
    "updateBook": {
      "id": "328545c8-68ba-4701-af8b-2282c1e573f3",
      "title": "Blossoms",
      "author": "Boniface",
      "publishedYear": 2014
    }
  }
}
```

### `deleteBook`

This mutation deletes a book from the library by its ID.

**Mutation:**

```graphql
mutation {
  deleteBook(id: "328545c8-68ba-4701-af8b-2282c1e573f3")
}
```

**Example Response:**

```json
{
  "data": {
    "deleteBook": true
  }
}
```

## Error Handling

If an error occurs during the execution of a query or mutation, the API will return an error response with a descriptive error message. For example, if you try to fetch a book with an invalid ID, you might receive a response like this:

```json
{
  "errors": [
    {
      "message": "book not found"
    }
  ],
  "data": null
}
```

## License

This project is licensed under the [MIT License](LICENSE).

Feel free to modify or extend this README file according to your specific requirements or additional features you might have in your API.