# Basic-Web-Server

Basic-Web-Server is a simple Go-based web server that serves a space-themed website called "Galactic Frontiers." This project demonstrates how to create a basic web server using Go, serve static files, handle form submissions, and generate dynamic content.

## Features

1. Static file serving
2. Custom route handling
3. Form processing
4. Dynamic HTML generation
5. Space-themed design

## Project Structure

- `src/`: Contains Go source files
  - `main.go`: Entry point of the application
  - `handlers/`: Package containing request handlers
    - `hello.go`: Handles the /hello endpoint
    - `form.go`: Processes form submissions
- `static/`: Contains static files
  - `index.html`: Homepage of Galactic Frontiers
  - `form.html`: Registration form for Galactic Citizens
  - `output.html`: Template for displaying registration results
  - `styles.css`: CSS styles for the website

## How It Works

1. The server starts by running the `main.go` file, which sets up routes and starts listening on port 8080.
2. Static files (HTML, CSS) are served from the `static/` directory.
3. The homepage (`index.html`) presents a space-themed landing page for Galactic Frontiers.
4. Users can navigate to the registration form (`form.html`) to register as a Galactic Citizen.
5. When the form is submitted, the `HandleForm` function in `form.go` processes the data and generates a dynamic response using the `output.html` template.
6. The server also includes a simple `/hello` endpoint that returns a greeting message.

## Running the Server

1. Ensure you have Go installed on your system.
2. Navigate to the project root directory.
3. Run the following command:
   ```
   go run src/main.go
   ```
4. Open a web browser and visit `http://localhost:8080` to view the website.

## Customization

You can easily customize the website by modifying the HTML, CSS, and Go files. Add new routes, update the design, or extend the functionality to suit your needs.

## License

This project is open-source and available under the MIT License.

