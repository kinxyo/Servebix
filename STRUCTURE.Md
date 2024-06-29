# First write everything done, complete the server code, then think of diving it into a structure.

```mylang
Go structure project:

    - Architecture
    |
    |
    *--------> Transport Layer (APIs (router, middleware, handler))
    |
    |
    *--------> Service Layer (Business Logic)
    |
    |
    *--------> Storage Layer (Data Access Layer)


    - STEPS
    |
    |
    *--------> Requests comes in
    |
    |
    *--------> Handler Fetches it (Transport Layer)
    |
    |
    *--------> HTTP handler is going to call the service (Service Layer (e.g.- `svc.getUser()`)
    |
    |
    *--------> Service talks to the Storage Layer (Storage Layer (e.g.- `storage.findUserByID()`)
```
