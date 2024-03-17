## Golang Vue Image Processing App

This application combines the power of Golang for backend processing and Nuxt.js (Vue.js) for frontend interaction to facilitate various image processing tasks. Whether you need to convert image formats, resize images, or compress them while maintaining quality, this app has got you covered.

GitHub Repo : https://github.com/qsaifudin/golang-vue-image-processing.git

Video Demo : 


## Features

- **Convert**: Convert image files from PNG to JPEG format.
- **Resize**: Resize images according to specified dimensions (width and height).
- **Compress**: Compress images to reduce file size while maintaining reasonable quality.

## How to Use

### Running Frontend and Backend

1. Ensure you have Golang and Nuxt.js installed on your system.
2. Clone the repository to your local machine.
3. Navigate to the backend and frontend directories and install dependencies:

## Using the Frontend
1. Upload an image file.
2. Choose the desired image processing option (convert, resize, or compress).
3. View the processed image.
4. Download the processed image if required.

## Using the APIs

### Convert PNG to JPEG:

- **POST** `http://localhost:5000/convert/png-to-jpeg`

  Send a multipart form-data request with the `image` file.

### Resize Image:

- **POST** `http://localhost:5000/resize`

  Send a multipart form-data request with the `image` file, `width`, and `height` fields.

### Compress Image:

- **POST** `http://localhost:5000/compress`

  Send a multipart form-data request with the `image` file and `quality` field.

### View Processed Image:

- **GET** `http://localhost:5000/temp_img/processed-image.jpeg`

  Retrieve the processed image after performing image processing operations.


Detail setup instructions and system requirements are already available in the `README.me` file inside the backend and frontend folders.


## Author

- **Saifudin**
  - Email: qsaifudin.official@gmail.com
  - LinkedIn: [https://www.linkedin.com/in/qsaifudin/](https://www.linkedin.com/in/qsaifudin/)
  - Personal Web: [https://qsaifudin.site/](https://qsaifudin.site/)

