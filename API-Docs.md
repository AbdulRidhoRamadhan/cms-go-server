# CMS-Go-Server API Documentation

## Base URL

`http://localhost:80`

## Endpoints

### Authentication

#### Login

- **URL**: `/users/login`
- **Method**: `POST`
- **Description**: Authenticate user and return access token.
- **Request Body**:

  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```

  ```json
  {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "role": "Admin"
  }
  ```

### Public Routes

#### Get All Products

- **URL**: `/pub`
- **Method**: `GET`
- **Description**: Retrieve all products with optional filtering and pagination.
- **Query Parameters**:
  - `search`: Search by product name.
  - `sort`: Sort by creation date (`asc` or `desc`).
  - `categoryId`: Filter by category ID.
  - `categoryName`: Filter by category name.
  - `page`: Page number (default: 1).
  - `limit`: Number of items per page (default: 10).
- **Response**:
  ```json
  {
    "products": [
      {
        "id": 1,
        "name": "Product 1",
        "description": "Description of Product 1",
        "price": 100000,
        "stock": 10,
        "imgUrl": "http://example.com/image.jpg",
        "categoryId": 1,
        "authorId": 1,
        "Category": {
          "id": 1,
          "name": "Category 1",
          "authorId": 1,
          "Author": {
            "id": 1,
            "username": "admin",
            "email": "admin@example.com",
            "role": "Admin",
            "phoneNumber": "1234567890",
            "address": "123 Main St"
          }
        },
        "User": {
          "id": 1,
          "username": "admin",
          "email": "admin@example.com",
          "role": "Admin",
          "phoneNumber": "1234567890",
          "address": "123 Main St"
        },
        "createdAt": "2023-10-01T00:00:00Z",
        "updatedAt": "2023-10-01T00:00:00Z"
      }
    ],
    "totalPages": 1,
    "currentPage": 1
  }
  ```

#### Get Product by ID

- **URL**: `/pub/:id`
- **Method**: `GET`
- **Description**: Retrieve a single product by its ID.
- **Response**:
  ```json
  {
    "productById": {
      "id": 1,
      "name": "Product 1",
      "description": "Description of Product 1",
      "price": 100000,
      "stock": 10,
      "imgUrl": "http://example.com/image.jpg",
      "categoryId": 1,
      "authorId": 1,
      "Category": {
        "id": 1,
        "name": "Category 1",
        "authorId": 1,
        "Author": {
          "id": 1,
          "username": "admin",
          "email": "admin@example.com",
          "role": "Admin",
          "phoneNumber": "1234567890",
          "address": "123 Main St"
        }
      },
      "User": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "Admin",
        "phoneNumber": "1234567890",
        "address": "123 Main St"
      },
      "createdAt": "2023-10-01T00:00:00Z",
      "updatedAt": "2023-10-01T00:00:00Z"
    }
  }
  ```

#### Get All Categories

- **URL**: `/pub/categories`
- **Method**: `GET`
- **Description**: Retrieve all categories.
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "Category 1",
      "authorId": 1,
      "Author": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "Admin",
        "phoneNumber": "1234567890",
        "address": "123 Main St"
      },
      "createdAt": "2023-10-01T00:00:00Z",
      "updatedAt": "2023-10-01T00:00:00Z"
    }
  ]
  ```

### Protected Routes (Require Authentication)

#### Add User (Admin Only)

- **URL**: `/users/add-user`
- **Method**: `POST`
- **Description**: Register a new user (Admin only).
- **Request Body**:
  ```json
  {
    "username": "newuser",
    "email": "newuser@example.com",
    "password": "password123",
    "phoneNumber": "1234567890",
    "address": "123 Main St",
    "role": "Staff"
  }
  ```
- **Response**:
  ```json
  {
    "id": 2,
    "username": "newuser",
    "email": "newuser@example.com",
    "phoneNumber": "1234567890",
    "address": "123 Main St",
    "role": "Staff"
  }
  ```

#### Create Product

- **URL**: `/products`
- **Method**: `POST`
- **Description**: Create a new product.
- **Request Body**:
  ```json
  {
    "name": "New Product",
    "description": "Description of New Product",
    "price": 150000,
    "stock": 5,
    "imgUrl": "http://example.com/new_image.jpg",
    "categoryId": 1
  }
  ```
- **Response**:
  ```json
  {
    "id": 2,
    "name": "New Product",
    "description": "Description of New Product",
    "price": 150000,
    "stock": 5,
    "imgUrl": "http://example.com/new_image.jpg",
    "categoryId": 1,
    "authorId": 1,
    "Category": {
      "id": 1,
      "name": "Category 1",
      "authorId": 1,
      "Author": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "Admin",
        "phoneNumber": "1234567890",
        "address": "123 Main St"
      }
    },
    "User": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "Admin",
      "phoneNumber": "1234567890",
      "address": "123 Main St"
    },
    "createdAt": "2023-10-01T00:00:00Z",
    "updatedAt": "2023-10-01T00:00:00Z"
  }
  ```

#### Get All Products

- **URL**: `/products`
- **Method**: `GET`
- **Description**: Retrieve all products.
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "Product 1",
      "description": "Description of Product 1",
      "price": 100000,
      "stock": 10,
      "imgUrl": "http://example.com/image.jpg",
      "categoryId": 1,
      "authorId": 1,
      "Category": {
        "id": 1,
        "name": "Category 1",
        "authorId": 1,
        "Author": {
          "id": 1,
          "username": "admin",
          "email": "admin@example.com",
          "role": "Admin",
          "phoneNumber": "1234567890",
          "address": "123 Main St"
        }
      },
      "User": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "Admin",
        "phoneNumber": "1234567890",
        "address": "123 Main St"
      },
      "createdAt": "2023-10-01T00:00:00Z",
      "updatedAt": "2023-10-01T00:00:00Z"
    }
  ]
  ```

#### Get Product by ID

- **URL**: `/products/:id`
- **Method**: `GET`
- **Description**: Retrieve a single product by its ID.
- **Response**:
  ```json
  {
    "id": 1,
    "name": "Product 1",
    "description": "Description of Product 1",
    "price": 100000,
    "stock": 10,
    "imgUrl": "http://example.com/image.jpg",
    "categoryId": 1,
    "authorId": 1,
    "Category": {
      "id": 1,
      "name": "Category 1",
      "authorId": 1,
      "Author": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "Admin",
        "phoneNumber": "1234567890",
        "address": "123 Main St"
      }
    },
    "User": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "Admin",
      "phoneNumber": "1234567890",
      "address": "123 Main St"
    },
    "createdAt": "2023-10-01T00:00:00Z",
    "updatedAt": "2023-10-01T00:00:00Z"
  }
  ```

#### Update Product

- **URL**: `/products/:id`
- **Method**: `PUT`
- **Description**: Update a product by its ID.
- **Request Body**:
  ```json
  {
    "name": "Updated Product",
    "description": "Updated Description",
    "price": 200000,
    "stock": 20,
    "imgUrl": "http://example.com/updated_image.jpg",
    "categoryId": 1
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "name": "Updated Product",
    "description": "Updated Description",
    "price": 200000,
    "stock": 20,
    "imgUrl": "http://example.com/updated_image.jpg",
    "categoryId": 1,
    "authorId": 1,
    "Category": {
      "id": 1,
      "name": "Category 1",
      "authorId": 1,
      "Author": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "Admin",
        "phoneNumber": "1234567890",
        "address": "123 Main St"
      }
    },
    "User": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "Admin",
      "phoneNumber": "1234567890",
      "address": "123 Main St"
    },
    "createdAt": "2023-10-01T00:00:00Z",
    "updatedAt": "2023-10-01T00:00:00Z"
  }
  ```

#### Delete Product

- **URL**: `/products/:id`
- **Method**: `DELETE`
- **Description**: Delete a product by its ID.
- **Response**:
  ```json
  {
    "message": "Product success to delete"
  }
  ```

#### Upload Product Image

- **URL**: `/products/upload/:id`
- **Method**: `PATCH`
- **Description**: Upload an image for a product.
- **Request Body**:
  - `image`: Image file (JPEG, PNG, GIF, WEBP).
- **Response**:
  ```json
  {
    "message": "Image Product success to update",
    "imgUrl": "http://example.com/new_image.jpg"
  }
  ```

### Categories

#### Create Category

- **URL**: `/categories`
- **Method**: `POST`
- **Description**: Create a new category.
- **Request Body**:
  ```json
  {
    "name": "New Category"
  }
  ```
- **Response**:
  ```json
  {
    "id": 2,
    "name": "New Category",
    "authorId": 1,
    "Author": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "Admin",
      "phoneNumber": "1234567890",
      "address": "123 Main St"
    },
    "createdAt": "2023-10-01T00:00:00Z",
    "updatedAt": "2023-10-01T00:00:00Z"
  }
  ```

#### Get All Categories

- **URL**: `/categories`
- **Method**: `GET`
- **Description**: Retrieve all categories.
- **Response**:
  ```json
  [
    {
      "id": 1,
      "name": "Category 1",
      "authorId": 1,
      "Author": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "role": "Admin",
        "phoneNumber": "1234567890",
        "address": "123 Main St"
      },
      "createdAt": "2023-10-01T00:00:00Z",
      "updatedAt": "2023-10-01T00:00:00Z"
    }
  ]
  ```

#### Update Category

- **URL**: `/categories/:id`
- **Method**: `PUT`
- **Description**: Update a category by its ID.
- **Request Body**:
  ```json
  {
    "name": "Updated Category"
  }
  ```
- **Response**:
  ```json
  {
    "message": "Category updated successfully"
  }
  ```

## Error Responses

All error responses follow the same format:

```json
{
  "message": "Error message"
}
```

### Common Error Messages

- `BAD_REQUEST`: Invalid input format.
- `NOT_FOUND`: Data not found.
- `INTERNAL_SERVER_ERROR`: Internal server error.
- `UNAUTHENTICATED`: Invalid email or password.
- `UNAUTHORIZED`: Invalid token.
- `FORBIDDEN`: You don't have access.
- `EMAIL_ALREADY_EXISTS`: Email already exists.
- `PASSWORD_LENGTH`: Password must be at least 5 characters.
- `PRICE_MIN`: Minimum price is Rp.100.000,00.
- `CATEGORY_NOT_FOUND`: Category not found.
- `NAME_REQUIRED`: Product name is required.
- `DESCRIPTION_REQUIRED`: Product description is required.
- `PRICE_REQUIRED`: Product price is required.
- `CATEGORY_REQUIRED`: Category is required.
- `IMAGE_REQUIRED`: Product image is required.
- `INVALID_IMAGE_TYPE`: Invalid image type. Supported types: JPEG, PNG, GIF, WEBP.
- `INVALID_EMAIL_FORMAT`: Enter the correct email type
