# The Description Of Ascii-art-web Project

**Ascii-art-web** is a web app that lets users generate ASCII text with different font banners via a **Graphical User Interface**. It processes input, renders ASCII output, and displays results dynamically.

## Authors Of Project

- @Ybouchraa  
- @Tamine  

## Usage: How to Run

### 1. Start the Server
- Open a terminal and navigate to the project directory.  
- Run the server using the appropriate command (e.g., `go run main.go`).

### 2. Access the Web Interface
- Open a web browser.  
- Enter `http://127.0.0.1:8080` in the address bar.

### 3. Generate ASCII Art
- Enter your text in the **INPUT** field.  
- Select a **BANNER** style from the dropdown menu.  
- Click the **Run** button to generate ASCII art.

### 4. Final Step
- The generated ASCII art will be displayed below.  
- Copy or use the output as needed.

## Algorithm for Ascii-art-web Project

### 1. Server Initialization
- Serve static files (CSS, JS, etc.) from the `static/` directory.
- Define an HTTP handler for the root route `/`.
- Start the server on port **8080**.

### 2. Handle Incoming Requests (`HandleRequest`)
#### Check URL Path:
- If the path is not `/`, return a **404 Not Found** error.

#### Load the HTML Template (`index.html`)
- If the template file is missing, return a **404 Not Found** error.

#### Handle GET Requests:
- Render the HTML page with an empty `Result`.

#### Handle POST Requests:
- Extract user input (`inp1`) and selected banner (`Files`).
- Validate the banner using `A.InitMap(banner)`.
  - If invalid, return **500 Internal Server Error**.
- Process the input text using `A.Storing(input)`.
- Render the template with the ASCII output.
- Return a **400 Bad Request** for unsupported methods.

### 3. Start the Server (`main`)
- Register HTTP handlers.
- Print a message indicating that the server is running.
- Start listening on port **8080** and handle errors gracefully.

