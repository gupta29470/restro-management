# golang-restaurant-management

## User

### User Model

```
type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	First_Name   *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_Name    *string            `json:"last_name" validate:"required,min=2,max=100"`
	Email        *string            `json:"email" validate:"required"`
	Password     *string            `json:"password" validate:"required"`
	Phone_Number *string            `jsson:"phone_number" validate:"required"`
	Avatar       *string            `json:"avatar"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	Created_At   time.Time          `json:"created_at"`
	Updated_At   time.Time          `json:"updated_at"`
	User_ID      string             `json:"user_id"`
}
```

### API Endpoints

##### 1. Signup
> Endpoint: /users/signup
> <br>
> Method: POST
> <br>
> Request Payload:

```
{
  "first_name": "aakash",
  "last_name": "gupta",
  "email": "aakash.gupta@gmail.com",
  "password": "123456789",
  "phone_number": "1234567890",
  "avatar": "image"
}
```

##### 2. Login
> Endpoint: /users/login
> <br>
> Method: POST
> <br>
> Request Payload:

```
{
  "email": "aakash.gupta@gmail.com",
  "password": "123456789"
}
```

> Response:
```
{
  "ID": "65b7df93cbb9a80a38bbb812",
  "first_name": "aakash",
  "last_name": "gupta",
  "email": "aakash.gupta@gmail.com",
  "password": "$2a$14$D8BzvwcfKgGwbu8tZq2QB.5BWaGr.IxIVZ6G96fkTpa0e3kJyfBMG",
  "Phone_Number": "1234567890",
  "avatar": "image",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFha2FzaC5ndXB0YUBnbWFpbC5jb20iLCJGaXJzdF9OYW1lIjoiYWFrYXNoIiwiTGFzdF9OYW1lIjoiZ3VwdGEiLCJQaG9uZV9OdW1iZXIiOiIxMjM0NTY3ODkwIiwiVXNlcl9JRCI6IjY1YjdkZjkzY2JiOWE4MGEzOGJiYjgxMiIsImV4cCI6MTcwNjk5NTU4Nn0.DuvP7-LOXapZcwKnff6zIDHbo-Oc5_kQIB5Ne_zFwXA",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJQaG9uZV9OdW1iZXIiOiIiLCJVc2VyX0lEIjoiIiwiZXhwIjoxNzA3MTUzOTg2fQ.aQnQI_-maQmMNUFg1yclZOqETYZZzOwu2_jN9M40Hv8",
  "created_at": "2024-01-29T17:25:39Z",
  "updated_at": "2024-01-29T17:25:39Z",
  "user_id": "65b7df93cbb9a80a38bbb812"
}
```

##### 3. Get Users
> Endpoint: /users
> <br>
> Method: GET
> <br>
> Response
```

[
  {
    "ID": "65b1d66a0754e4fbd10025f5",
    "first_name": "Aakash",
    "last_name": "Gupta",
    "email": "aa@gmail.com",
    "password": "$2a$14$Li1UpP44ZFWy2Kwfa40GHeIaNg4x06Cd0.2Zambro3S6jgbHBOuTe",
    "Phone_Number": "123456789",
    "avatar": null,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFhQGdtYWlsLmNvbSIsIkZpcnN0X05hbWUiOiJBYWthc2giLCJMYXN0X05hbWUiOiJHdXB0YSIsIlBob25lX051bWJlciI6IjEyMzQ1Njc4OSIsIlVzZXJfSUQiOiI2NWIxZDY2YTA3NTRlNGZiZDEwMDI1ZjUiLCJleHAiOjE3MDY1OTk5Nzh9.eL1XfelLt2K5hUnEiAKaJGHf1cGTyv9-COUmpmhE2KI",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJQaG9uZV9OdW1iZXIiOiIiLCJVc2VyX0lEIjoiIiwiZXhwIjoxNzA2NzU4Mzc4fQ.p-DblQg3vrNGgpfbSSaP8Olzu12ey1V4WHjtZ1p4z5w",
    "created_at": "2024-01-25T03:32:58Z",
    "updated_at": "2024-01-25T03:32:58Z",
    "user_id": "65b1d66a0754e4fbd10025f5"
  },
  {
    "ID": "65b1d75fce0a7da5f8d77504",
    "first_name": "Aakash",
    "last_name": "Gupta",
    "email": "a1@gmail.com",
    "password": "$2a$14$IWHKWGVbjQYxmdBJJ0sIhONdPpXqcNpe5E3C7vOIGT/ifr4uKLS9i",
    "Phone_Number": "123456789",
    "avatar": null,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImExQGdtYWlsLmNvbSIsIkZpcnN0X05hbWUiOiJBYWthc2giLCJMYXN0X05hbWUiOiJHdXB0YSIsIlBob25lX051bWJlciI6IjEyMzQ1Njc4OSIsIlVzZXJfSUQiOiI2NWIxZDc1ZmNlMGE3ZGE1ZjhkNzc1MDQiLCJleHAiOjE3MDY2MDAyMjN9.6pG8jMrEsC0bUrkivln5tEGdso5YRcP3KG6YgQzd310",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJQaG9uZV9OdW1iZXIiOiIiLCJVc2VyX0lEIjoiIiwiZXhwIjoxNzA2NzU4NjIzfQ.JafcJgK1_4tgr-LN94OIbQ8f0zSSQxL40n82H1zE4B4",
    "created_at": "2024-01-25T03:37:03Z",
    "updated_at": "2024-01-25T03:37:03Z",
    "user_id": "65b1d75fce0a7da5f8d77504"
  },
  {
    "ID": "65b1d878c3532651e9a518e8",
    "first_name": "Aakash",
    "last_name": "Gupta",
    "email": "a2@gmail.com",
    "password": "$2a$14$vqgGVste3k7M7SqBXESzyuOaHnXHrGfe8bp1NCoDfgdbbKQ2PzEyS",
    "Phone_Number": "123456789",
    "avatar": null,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImEyQGdtYWlsLmNvbSIsIkZpcnN0X05hbWUiOiJBYWthc2giLCJMYXN0X05hbWUiOiJHdXB0YSIsIlBob25lX051bWJlciI6IjEyMzQ1Njc4OSIsIlVzZXJfSUQiOiI2NWIxZDg3OGMzNTMyNjUxZTlhNTE4ZTgiLCJleHAiOjE3MDY2MDA1MDR9.1E9NnO2GFnKrY2iBvDvCqD7LG-XeY3NKg7vvoi9Z3zE",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJQaG9uZV9OdW1iZXIiOiIiLCJVc2VyX0lEIjoiIiwiZXhwIjoxNzA2NzU4OTA0fQ.vstXpcD-w5HEw9zjnEkWUGGOAnU4YmRpSP21yZaQlzc",
    "created_at": "2024-01-25T03:41:44Z",
    "updated_at": "2024-01-25T03:41:44Z",
    "user_id": "65b1d878c3532651e9a518e8"
  },
  {
    "ID": "65b1d9020503cf98847b338b",
    "first_name": "Aakash",
    "last_name": "Gupta",
    "email": "a3@gmail.com",
    "password": "$2a$14$/ThXUx8nWLhc3EMf/huJ/uWVKgcVJgMcrCrdmRqsBrwuSol.VKZfe",
    "Phone_Number": "123456789",
    "avatar": null,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImEzQGdtYWlsLmNvbSIsIkZpcnN0X05hbWUiOiJBYWthc2giLCJMYXN0X05hbWUiOiJHdXB0YSIsIlBob25lX051bWJlciI6IjEyMzQ1Njc4OSIsIlVzZXJfSUQiOiI2NWIxZDkwMjA1MDNjZjk4ODQ3YjMzOGIiLCJleHAiOjE3MDY2MDE0OTF9.4w2oqBzwMzYVnuEGXEV9xKJYxfpnokXyjsNEsggawic",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJQaG9uZV9OdW1iZXIiOiIiLCJVc2VyX0lEIjoiIiwiZXhwIjoxNzA2NzU5MDQyfQ.POB6t4FHInPTTjyQA2KWZPZaoqI90nWpCOtCzk2nX-A",
    "created_at": "2024-01-25T03:44:02Z",
    "updated_at": "2024-01-25T03:58:11Z",
    "user_id": "65b1d9020503cf98847b338b"
  },
  {
    "ID": "65b7df93cbb9a80a38bbb812",
    "first_name": "aakash",
    "last_name": "gupta",
    "email": "aakash.gupta@gmail.com",
    "password": "$2a$14$D8BzvwcfKgGwbu8tZq2QB.5BWaGr.IxIVZ6G96fkTpa0e3kJyfBMG",
    "Phone_Number": "1234567890",
    "avatar": "image",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFha2FzaC5ndXB0YUBnbWFpbC5jb20iLCJGaXJzdF9OYW1lIjoiYWFrYXNoIiwiTGFzdF9OYW1lIjoiZ3VwdGEiLCJQaG9uZV9OdW1iZXIiOiIxMjM0NTY3ODkwIiwiVXNlcl9JRCI6IjY1YjdkZjkzY2JiOWE4MGEzOGJiYjgxMiIsImV4cCI6MTcwNjk5NTU4Nn0.DuvP7-LOXapZcwKnff6zIDHbo-Oc5_kQIB5Ne_zFwXA",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJQaG9uZV9OdW1iZXIiOiIiLCJVc2VyX0lEIjoiIiwiZXhwIjoxNzA3MTUzOTM5fQ.w_P6e9Gly50OVUU3n_AkXJIU61GUQ4w45Cw3mgDhB8M",
    "created_at": "2024-01-29T17:25:39Z",
    "updated_at": "2024-01-29T17:26:26Z",
    "user_id": "65b7df93cbb9a80a38bbb812"
  }
]
```

##### 4. Get User
> Endpoint: /users/:user_id
> <br>
> Method: GET
> <br>
> Response

```
{
  "ID": "65b1d66a0754e4fbd10025f5",
  "first_name": "Aakash",
  "last_name": "Gupta",
  "email": "aa@gmail.com",
  "password": "$2a$14$Li1UpP44ZFWy2Kwfa40GHeIaNg4x06Cd0.2Zambro3S6jgbHBOuTe",
  "Phone_Number": "123456789",
  "avatar": null,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFhQGdtYWlsLmNvbSIsIkZpcnN0X05hbWUiOiJBYWthc2giLCJMYXN0X05hbWUiOiJHdXB0YSIsIlBob25lX051bWJlciI6IjEyMzQ1Njc4OSIsIlVzZXJfSUQiOiI2NWIxZDY2YTA3NTRlNGZiZDEwMDI1ZjUiLCJleHAiOjE3MDY1OTk5Nzh9.eL1XfelLt2K5hUnEiAKaJGHf1cGTyv9-COUmpmhE2KI",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X05hbWUiOiIiLCJMYXN0X05hbWUiOiIiLCJQaG9uZV9OdW1iZXIiOiIiLCJVc2VyX0lEIjoiIiwiZXhwIjoxNzA2NzU4Mzc4fQ.p-DblQg3vrNGgpfbSSaP8Olzu12ey1V4WHjtZ1p4z5w",
  "created_at": "2024-01-25T03:32:58Z",
  "updated_at": "2024-01-25T03:32:58Z",
  "user_id": "65b1d66a0754e4fbd10025f5"
}
```

<br>
<br>

## Menu

### Menu Model

```
type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name" validate:"required"`
	Category   string             `json:"category" validate:"required"`
	Start_Date *time.Time         `json:"start_date"`
	End_Date   *time.Time         `json:"end_date"`
	Created_At time.Time          `json:"created_at"`
	Updated_At time.Time          `json:"updated_at"`
	Menu_ID    string             `json:"menu_id"`
}
```

### API Endpoints

##### 1. Add Menu
> Endpoint: /menus
> <br>
> Method: POST
> <br>
> Request Payload:
```
{
  "name": "Food",
  "category": "Mughlai"
}
```

##### 2. Get Menus
> Endpoint: /menus
> <br>
> Method: GET
> <br>
> Response:
```
[
  {
    "_id": "65b1dda4d116a73f754f6c35",
    "category": "Punjabi",
    "created_at": "2024-01-25T04:03:48Z",
    "end_date": null,
    "menu_id": "65b1dda4d116a73f754f6c35",
    "name": "Food",
    "start_date": null,
    "updated_at": "2024-01-25T13:26:39Z"
  },
  {
    "_id": "65b260f2d116a73f754f6c37",
    "category": "Mughlai",
    "created_at": "2024-01-25T13:24:02Z",
    "end_date": null,
    "menu_id": "65b260f2d116a73f754f6c37",
    "name": "Food",
    "start_date": null,
    "updated_at": "2024-01-25T13:24:02Z"
  }
]
```

##### 3. Get Menu
> Endpoint: /menus/:menu_id
> <br>
> Method: GET
> <br>
> Response:
```
{
  "ID": "65b260f2d116a73f754f6c37",
  "name": "Food",
  "category": "Mughlai",
  "start_date": null,
  "end_date": null,
  "created_at": "2024-01-25T13:24:02Z",
  "updated_at": "2024-01-25T13:24:02Z",
  "menu_id": "65b260f2d116a73f754f6c37"
}
```


##### 4. Update Menu
> Endpoint: /menus/:menu_id
> <br>
> Method: PATCH
> <br>
> Request Payload:
```
{
  "name": "Food",
  "category": "Punjabi"
}
```

<br>
<br>

## Food

### Food Model

```
type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" validate:"required"`
	Food_Image *string            `json:"food_image" validate:"required"`
	Created_At time.Time          `json:"created_at"`
	Updated_At time.Time          `json:"updated_at"`
	Food_ID    string             `json:"food_id"`
	Menu_ID    *string            `json:"menu_id" validate:"required"`
}
```

### API Endpoints

##### 1. Add Food
> Endpoint: /foods
> <br>
> Method: POST
> <br>
> Request Payload:
```
{
  "name": "Dahi Cham Cham",
  "price": 350,
  "food_image": "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxleHBsb3JlLWZlZWR8NHx8fGVufDB8fHx8fA%3D%3D",
  "menu_id": "65b260f2d116a73f754f6c37"
}
```

##### 2. Get Foods
> Endpoint: /foods
> <br>
> Method: GET
> <br>
> Response:
```
[
  {
    "ID": "65b36e21d1bf83f8fe68f3bb",
    "name": "Dahi Vade",
    "price": 350,
    "food_image": "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxleHBsb3JlLWZlZWR8NHx8fGVufDB8fHx8fA%3D%3D",
    "created_at": "2024-01-26T08:32:33Z",
    "updated_at": "2024-01-26T08:32:33Z",
    "food_id": "65b36e21d1bf83f8fe68f3bb",
    "menu_id": "65b260f2d116a73f754f6c37"
  },
  {
    "ID": "65b36e29d1bf83f8fe68f3bd",
    "name": "Dahi Cham Cham",
    "price": 350,
    "food_image": "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxleHBsb3JlLWZlZWR8NHx8fGVufDB8fHx8fA%3D%3D",
    "created_at": "2024-01-26T08:32:41Z",
    "updated_at": "2024-01-26T08:32:41Z",
    "food_id": "65b36e29d1bf83f8fe68f3bd",
    "menu_id": "65b260f2d116a73f754f6c37"
  }
]
```

##### 3. Get Food
> Endpoint: /foods/:food_id
> <br>
> Method: GET
> <br>
> Response:
```
{
  "ID": "65b26324d116a73f754f6c3b",
  "name": "Kalimiri Kabab",
  "price": 339,
  "food_image": "https://images.unsplash.com/photo-1567620905732-2d1ec7ab7445?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxleHBsb3JlLWZlZWR8NHx8fGVufDB8fHx8fA%3D%3D",
  "created_at": "2024-01-25T13:33:24Z",
  "updated_at": "2024-01-25T13:40:22Z",
  "food_id": "65b26324d116a73f754f6c3b",
  "menu_id": "65b260f2d116a73f754f6c37"
}
```


##### 4. Update Food
> Endpoint: /foods/:food_id
> <br>
> Method: PATCH
> <br>
> Request Payload:
```
{
  "price": 339
}
```


<br>
<br>

## Table

### Table Model

```
type Table struct {
	ID               primitive.ObjectID `bson:"_id"`
	Number_Of_Guests *int               `json:"number_of_guests" validate:"required"`
	Table_Number     *int               `json:"table_number" validate:"required"`
	Created_At       time.Time          `json:"created_at"`
	Updated_At       time.Time          `json:"updated_at"`
	Table_ID         string             `json:"table_id"`
}
```

### API Endpoints

##### 1. Add Table
> Endpoint: /tables
> <br>
> Method: POST
> <br>
> Request Payload:
```
{
  "number_of_guests": 2,
  "table_number": 3
}
```

##### 2. Get Tables
> Endpoint: /tables
> <br>
> Method: GET
> <br>
> Response:
```
[
  {
    "_id": "65b265c798cd2627196bf909",
    "created_at": "2024-01-25T13:44:39Z",
    "number_of_guests": 4,
    "table_id": "65b265c798cd2627196bf909",
    "table_number": 1,
    "updated_at": "2024-01-25T13:44:39Z"
  },
  {
    "_id": "65b265d298cd2627196bf90b",
    "created_at": "2024-01-25T13:44:50Z",
    "number_of_guests": 6,
    "table_id": "65b265d298cd2627196bf90b",
    "table_number": 4,
    "updated_at": "2024-01-25T13:44:50Z"
  },
  {
    "_id": "65b265d898cd2627196bf90d",
    "created_at": "2024-01-25T13:44:56Z",
    "number_of_guests": 2,
    "table_id": "65b265d898cd2627196bf90d",
    "table_number": 3,
    "updated_at": "2024-01-25T13:44:56Z"
  },
  {
    "_id": "65f09936ac33ec1b6edadd26",
    "created_at": "2024-03-12T18:04:38Z",
    "number_of_guests": 2,
    "table_id": "65f09936ac33ec1b6edadd26",
    "table_number": 3,
    "updated_at": "2024-03-12T18:04:38Z"
  }
]
```

##### 3. Get Table
> Endpoint: /tables/:table_id
> <br>
> Method: GET
> <br>
> Response:
```
{
  "ID": "65b265d298cd2627196bf90b",
  "number_of_guests": 6,
  "table_number": 4,
  "created_at": "2024-01-25T13:44:50Z",
  "updated_at": "2024-01-25T13:44:50Z",
  "table_id": "65b265d298cd2627196bf90b"
}
```


##### 4. Update Table
> Endpoint: /tables/:table_id
> <br>
> Method: PATCH
> <br>
> Request Payload:
```
{
  "table_number": 4
}
```
