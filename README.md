# ascii-art-stylize

This project is a web application that generates ASCII art from user input. The application allows users to select different fonts to create their ASCII art.

## Project Layout
### ascii/

-    **asciiFunc.go:** Contains handler functions for the routes.
-    **filecheck.go:** Checks for the existence of files.
-    **loadbanner.go:** Loads banner files.
-    **printbanner.go:** Prints the ASCII art based on the selected banner.

### bannerfiles/

-    **shadow.txt:** Contains the shadow font.
-    **standard.txt:** Contains the standard font.
-    **thinkertoy.txt:** Contains the thinkertoy font.

### routes/

-    **routes.go:** Registers the routes for the application.

### static/css/

-    **style.css:** Contains the styles for the application.

### templates/

-    **index.html:** The main HTML template for the application.

### main.go

* Initializes the server, sets up routes, and starts listening for incoming requests.


## Getting Started

### Prerequisites

Make sure you have Go installed on your machine. You can download and install it from the official [Go website](https://golang.org/dl/).

### Installation

1. Clone the repository:
    ```
    git clone https://github.com/ombima56/ascii-art-stylize.git
    cd ascii-art-stylize
    ```

2. Install dependencies:
    ```
    go mod tidy
    ```


## Usage

- To start the server, run:
 ```
 go run main.go
 ```
- Open your browser and go to `http://localhost:8080` to access the website.
- Enter text in the textarea.
- Select a font from the dropdown.
- Click the "Generate" button to display the ASCII art.

## Contribution

Feel free to fork this repository and make changes. Pull requests are welcome!
## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.