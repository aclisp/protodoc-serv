<template>
  <v-container>

    <div>选取一个 Protobuf 文件，把它转换为文档。</div>

    <v-row justify="center" align="center">
      <v-col cols="6">
        <v-file-input
          show-size
          label="Protobuf 协议文件"
          @change="selectFile"
        ></v-file-input>
      </v-col>

      <v-spacer></v-spacer>
    </v-row>

    <v-row>
      <v-col cols="6">
        <code v-if="fileContent">{{fileContent}}</code>
      </v-col>

      <v-col cols="6">
        <div v-if="docContent" class="doc" v-html="docContent"></div>
      </v-col>
    </v-row>

  </v-container>
</template>

<script>
import http from '../http-common'

export default {
  name: 'ProtoFileConverter',

  data: () => ({
    fileName: '',
    fileContent: '',
    docContent: ''
  }),

  methods: {
    selectFile (file) {
      if (file == null) {
        return
      }
      this.fileName = file.name
      const reader = new FileReader()
      reader.onload = (event) => {
        const contents = event.target.result
        this.fileContent = contents

        this.convert()
      }
      reader.readAsText(file)
    },
    convert () {
      const contents = this.fileContent
      if (contents) {
        http.post('/protodoc/ProtoDoc/Convert', {
          proto: contents,
          filename: this.fileName
        })
          .then(res => {
            this.docContent = res.data.html
          })
          .catch(err => {
            console.error('Error: ' + err)
          })
      }
    }
  }
}
</script>

<style scoped>
code {
  display: block;
  white-space: pre-wrap
}
</style>

<style>
/* global styles for docContent */
</style>
