<template>
  <v-app-bar app fixed>
    <v-container class="d-flex" width="100%">
      <v-btn class="px-0 pr-2">
        <NuxtLink
          class=""
          to="https://qsaifudin.site/"
          target="_blank"
          style="color: white; text-decoration: none"
        >
          <v-toolbar-title>
            <h3 style="font-size: 30px">👉<span class="animate-text">Saifudin</span></h3>
          </v-toolbar-title>
        </NuxtLink>
      </v-btn>
      <v-spacer></v-spacer>
      <div class="">
        <v-btn class="pt-0 mt-0" href="https://www.linkedin.com/in/qsaifudin/" target="_blank"
          >Linkedin</v-btn
        >
        <v-btn class="pt-0 mt-0" href="https://qsaifudin.site/" target="_blank">Website</v-btn>
        <v-btn class="pt-0 mt-0" href="https://github.com/qsaifudin" target="_blank">Github</v-btn>
      </div>
    
    </v-container>
  </v-app-bar>
</template>

<script>
export default {
  data() {
    return {
 
    };
  },
  mounted() {

  },
  methods: {
 
  },
};
</script>

<style scoped>
.graph {
  width: 800px;
  height: 600px;
  border: 1px solid #000;
}

.animate-text {
  font-family: "Open Sans";
  font-weight: 800;
  font-size: 30px;
  background: linear-gradient(to right, currentColor 0, #a2ebd58f 10%, currentColor 20%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  animation: borderAnimate 10s infinite alternate;
}

@keyframes borderAnimate {
  0%,
  100% {
    background-position: -100px;
  }

  50% {
    background-position: 510px;
  }
}
</style>
