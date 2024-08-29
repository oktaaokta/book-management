# Book Management API

## Overview

The Book Management API provides a backend service for managing books. It supports two primary functionalities:
1. **Get a List of Books by Genre / Subject:** Fetches book information from an external API based on a specified genre or subject.
2. **Submit a Book Pickup Schedule:** Allows users to submit a schedule for picking up books.

- **Method:** `POST`
- **Endpoint:** `/api/books/pickup`

**Request Payload:**

The request payload should be a JSON object with the following structure:

```json
{
    "edition": "{{edition_key}}",
    "pickup_date": "2006-01-02T15:04:05Z",
    "return_date": "2006-01-02T15:04:05Z"
}