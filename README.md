# HTTP Server Example

This is a simple Go program that demonstrates the setup of an HTTP server with two request handlers: `formHandler` and `helloHandler`.

## Requirements

- Go programming language (version 1.13 or higher)

## Getting Started

1. Clone the repository or download the code files.

2. Navigate to the project directory.

3. Build the executable by running the following command:

   ```bash
   go build
   ```

4. Run the compiled program:

   ```bash
   ./http-server-example
   ```

   You should see the message "Starting server at port 8080" indicating that the server has started successfully.

5. Open a web browser and visit `http://localhost:8080`. You should see a file server in action, serving the files from the "./static" directory.

## Handlers

### Form Handler

The `formHandler` function handles a POST request sent to the "/form" URL. It expects form data with "name" and "address" fields. Upon receiving a valid request, it parses the form data and writes the values of "name" and "address" to the response.

### Hello Handler

The `helloHandler` function handles a GET request sent to the "/hello" URL. It simply responds with the message "hello!".

## Usage

- **File Serving**: The server is configured to serve static files from the "./static" directory. Place your static files (HTML, CSS, JavaScript, images, etc.) in this directory, and they will be accessible through the root URL ("/").

- **Form Submission**: To submit a form, send a POST request to the "/form" URL with the form data. Make sure to include the "name" and "address" fields in the request payload.

- **Hello Message**: To receive a "hello!" response, send a GET request to the "/hello" URL.

## Customization

- **Port**: By default, the server runs on port 8080. If you want to use a different port, modify the `http.ListenAndServe` line in the `main` function.

- **File Serving Directory**: If you want to serve files from a different directory, update the `http.Dir` parameter in the `http.FileServer` line in the `main` function.

## Contributing

Contributions are welcome! If you find any issues or want to add new features, feel free to submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
