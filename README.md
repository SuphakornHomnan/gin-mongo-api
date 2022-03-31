# Implement RESTful API by Gin & mongodb

| Service               | Method | Route         | Body                      |
| --------------------- | ------ | ------------- | ------------------------- |
| Create user           | POST   | /user         | { Name, Location, Title } |
| Get user by userId    | GET    | /user/:userId |                           |
| Get all user          | GET    | /users        |                           |
| Edit user by userId   | PUT    | /user/:userId | { Name, Location, Title } |
| Delete user by userId | DELETE | /user/:userId |                           |
