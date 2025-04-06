# Simple Web Server

- Start a HTTP Server
- Define and handle routes (URLS)
- Serve static html files
- optionally serve with json data as well - mock api

### Possible Features
- 1️⃣ Serve Static HTML Pages

    Create a couple of basic .html files (e.g., index.html, about.html).

    The server should return these pages when users access the corresponding routes like / or /about.
- 2️⃣ Serve Static Assets (Optional)

    Serve static files like CSS, images, or JavaScript from a /static directory.

    Let the browser load assets from the server like a real website.
- Route Handling

    Implement multiple URL paths (e.g., /, /about, /contact).

    Each route returns a different HTML page or message.
- 4️⃣ JSON Response Endpoint (Bonus)

    Add a basic /api/status or /api/info route that returns a JSON response.

    Great intro to building a RESTful API.

🚀 Optional Enhancements

    Add logging for incoming requests.

    Handle 404 errors with a custom HTML page.

    Use environment variables for port configuration.

    Add a template engine later (e.g., Go’s html/template) for dynamic content.

Context Usefulness:

    You're using ctx.Value(keyServerAddr) in the getRoot handler, but you're not setting this context value anywhere in your current code. Consider either setting it up properly or removing it if it's not yet needed.