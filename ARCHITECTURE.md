## Security

1. **Encrypting local storage**: Once the user is registered in database, the server returns `username` and encrypted `id` to the client-side, which is then stored in the localStorage. The `id` is This just seemed a better way to implement authentication without involving the complexity of _tokens_.

2. **Server-side Validation**: The `id` is always validated by the server before any online-content being provided. 

3. **Using `HTTPS`**: self-explanatory.
