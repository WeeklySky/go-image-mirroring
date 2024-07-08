


# Image Mirroring

This project is a simple Go module that mirrors an image horizontally or vertically. The program reads an image file, processes it, and outputs the result as another image file.

## Getting Started

These instructions will help you set up the project on your local machine for development and testing purposes.

### Prerequisites

-   [Go](https://golang.org/doc/install) (version 1.22+)

### Installing

1.  **Clone the repository:**
    
    ```sh
    git clone https://github.com/guilhermemcardoso/go-image-mirroring
    cd go-image-mirroring
    ``` 
    
2. **Build the project:**
        
    ```sh
    go build -o gomirroring
    ``` 
    

## Usage

To use the Image Mirroring application, run the following command:

`./gomirroring -input <yourimage.png> -output <youroutput.png> -effect <horizontal | vertical>` 

This command will generate an image file with the same name and extension as the output inserted param.

### Example

Suppose you have an image file named `image.png`. To mirror it vertically, use:

```sh
./gomirroring -input image.png -output output.png -effect vertical
``` 

The program will create a file named `output.png` containing the mirrored image.

## Contributing

If you have suggestions for improving this project, feel free to submit a pull request or open an issue.
