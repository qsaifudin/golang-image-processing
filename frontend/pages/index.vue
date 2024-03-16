<template>
  <v-card class="mx-auto py-2 pb-8 text-center" elevation="8" rounded="lg">
    <v-card-title class="text-h5 pb-6">Image Processing App Using Golang and Vuejs</v-card-title>
    <v-row class="px-3">
      <v-col>
        <v-row>
          <v-col class="text-center mx-auto">
            <h3 class="mb-4">Choose an image to process</h3>

            <v-img
              class="mx-auto mb-3 "
              :src="previewImageUrl"
              max-width="500"
              max-height="400"
              contain
            ></v-img>
            <v-file-input
              variant="outlined"
              v-model="selectedFile"
              label="Choose an image"
              accept="image/*"
              @change="handleFileChange"
              show-size
              hide-details="true"
            ></v-file-input>
          </v-col>
        </v-row>
        <v-row no-gutters>
          <v-col>
            <div class="my-4">Image Processing Options</div>
            <v-row>
              <v-col>
                <v-btn
                  color="primary"
                  :loading="loading"
                  @click="convertToJPEG"
                  :disabled="!selectedFile"
                  >Convert to JPEG</v-btn
                >
              </v-col>

              <v-col>
                
                <v-btn block
                  class="mb-3"
                  :loading="loading"
                  color="primary"
                  @click="resizeImage"
                  :disabled="!selectedFile"
                  >Resize Image</v-btn
                >

                <v-text-field
                  variant="outlined"
                  v-model="width"
                  label="Width"
                  type="number"
                ></v-text-field>
                <v-text-field
                  variant="outlined"
                  v-model="height"
                  label="Height"
                  type="number"
                ></v-text-field>
              </v-col>

              <v-col>
                <v-btn
                  class="mb-3"
                  :loading="loading"
                  color="primary"
                  @click="compressImage"
                  :disabled="!selectedFile"
                  >Compress Image</v-btn
                >
                <v-text-field
                  variant="outlined"
                  v-model="quality"
                  label="Quality"
                  type="number"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-col>
        </v-row>
      </v-col>
      <v-divider vertical></v-divider>
      <v-col>
        <h3 class="mb-4">Result</h3>
        <v-row>
          <v-col class="text-center">
            <v-img
              class="text-center mx-auto"
              :src="`${processedImageUrl}?t=${Date.now()}`"
              v-if="processedImageUrl"
              max-width="500"
              max-height="400"
              @load="loading = false"
              @error="loading = false"
            ></v-img>
            <v-btn class="mt-2" v-if="processedImageUrl" color="primary" @click="downloadImage"
              >Download Processed Image</v-btn
            >
          </v-col>
        </v-row>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      selectedFile: null,
      previewImageUrl: null,
      processedImageUrl: null,
      width: null,
      height: null,
      quality: null,
      loading: false,
    };
  },
  methods: {
    handleFileChange(files) {
      const file = files.target.files[0];

      if (file) {
        this.selectedFile = file; // Update selectedFile directly
        this.previewImageUrl = URL.createObjectURL(file);
      } else {
        this.previewImageUrl = null;
      }
    },
    async convertToJPEG() {
      const formData = new FormData();
      formData.append("image", this.selectedFile);
      const config = useRuntimeConfig();
      this.processedImageUrl = null;
      this.loading = true;

      try {
        const response = await $fetch(`${config.public.BASE_URL}/convert/png-to-jpeg`, {
          method: "POST",

          body: formData,
        });
        if (!response.success) {
          throw new Error("Failed to convert to JPEG");
        }

        this.processedImageUrl = response.url; // Assuming your backend returns the URL of the processed image
      } catch (error) {
        console.error("Error converting to JPEG:", error);
      } finally {
        // Set loading back to false after the request completes
        this.loading = false;
      }
    },

    async resizeImage() {
      if (!this.selectedFile) {
        console.error("No file selected");
        return;
      }

      const formData = new FormData();
      formData.append("image", this.selectedFile);
      formData.append("width", this.width); // Using width from data
      formData.append("height", this.height); // Using height from data
      const config = useRuntimeConfig();
      this.processedImageUrl = null;
      this.loading = true;
      try {
        const response = await $fetch(`${config.public.BASE_URL}/resize`, {
          method: "POST",
          body: formData,
        });
        if (!response.success) {
          throw new Error("Failed to resize image");
        }

        this.processedImageUrl = response.url;
        console.log("Resized image URL:", this.processedImageUrl);
      } catch (error) {
        console.error("Error resizing image:", error);
      } finally {
        // Set loading back to false after the request completes
        this.loading = false;
      }
    },

    async compressImage() {
      if (!this.selectedFile) {
        console.error("No file selected");
        return;
      }

      const formData = new FormData();
      formData.append("image", this.selectedFile);
      formData.append("quality", this.quality); // Using quality from data
      const config = useRuntimeConfig();
      this.processedImageUrl = null;
      this.loading = true;
      try {
        const response = await $fetch(`${config.public.BASE_URL}/compress`, {
          method: "POST",
          body: formData,
        });
        if (!response.success) {
          throw new Error("Failed to compress image");
        }

        this.processedImageUrl = response.url;
        console.log("Compressed image URL:", this.processedImageUrl);
      } catch (error) {
        console.error("Error compressing image:", error);
      } finally {
        // Set loading back to false after the request completes
        this.loading = false;
      }
    },

    async downloadImage() {
      if (!this.processedImageUrl) {
        console.error("No processed image available to download");
        return;
      }

      try {
        // Make a GET request to the API endpoint that serves the image
        const response = await fetch("http://localhost:5000/temp_img/processed-image.jpeg", {
          method: "GET",
        });

        // Check if the request was successful
        if (!response.ok) {
          throw new Error("Failed to download image");
        }

        // Convert the response to a blob
        const blob = await response.blob();

        // Create a temporary URL for the blob
        const url = window.URL.createObjectURL(blob);

        // Create a temporary link element
        const link = document.createElement("a");
        link.href = url;
        link.setAttribute("download", "processed-image.jpeg");

        // Programmatically click on the link to trigger the download
        link.click();

        // Clean up
        window.URL.revokeObjectURL(url);
      } catch (error) {
        console.error("Error downloading image:", error);
      }
    },
  },
};
</script>
