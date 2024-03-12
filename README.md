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
