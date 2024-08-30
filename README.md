# Book Management API

## Overview

The Book Management API provides a backend service for managing books. It supports two primary functionalities:
1. **Get a List of Books by Genre / Subject:** Fetches book information from an external API based on a specified genre or subject.
- **Method:** `GET`
- **Endpoint:** `/api/books`
- **Request Parameters:**
    - `subject` (optional): The subject to filter books by.
- **URL:**
    - `http://127.0.0.1:8000/api/books?subject=science_fiction`

2. **Submit a Book Pickup Schedule:** Allows users to submit a schedule for picking up books.

- **Method:** `POST`
- **Endpoint:** `/api/books/pickup`
- **Example URL:**
    - `http://127.0.0.1:8000/api/books/pickup`

**Request Payload:**

The request payload should be a JSON object with the following structure:

```json
{
    "edition": "{{edition_key}}",
    "pickup_date": "2006-01-02T15:04:05Z",
    "return_date": "2006-01-02T15:04:05Z"
}
```

## Running the Application

To run the Book Management API on a Linux system, follow these steps:

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/oktaaokta/book-management.git
   cd book-management
   ```
2. **Run the Service:**
    After the 
    ```
    make run
    ```
