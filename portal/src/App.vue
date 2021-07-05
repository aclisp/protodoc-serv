<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
      absolute
    >
      <div class="d-flex align-center">
        <v-img
          alt="Logo"
          class="shrink mr-2"
          contain
          src="./assets/logo.png"
          transition="scale-transition"
          width="40"
        />

        <span
          class="shrink mt-1 headline font-weight-bold hidden-sm-and-down"
          contain
          min-width="100"
          width="100"
        >ProtoDoc</span>
      </div>

      <v-spacer></v-spacer>

      <v-btn
        href="https://github.com/aclisp/protodoc-serv"
        target="_blank"
        text
      >
        <span class="mr-2">Source</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <ProtoFileConverter/>
    </v-main>
  </v-app>
</template>

<script>
import ProtoFileConverter from './components/ProtoFileConverter'

export default {
  name: 'App',

  components: {
    ProtoFileConverter
  },

  data: () => ({
    //
  })
}
</script>
