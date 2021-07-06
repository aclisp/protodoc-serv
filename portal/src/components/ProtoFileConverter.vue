<template>
  <v-container fluid>

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
        <div v-if="docContent" class="markdown" v-html="docContent"></div>
        <v-alert v-if="errMessage" type="error">{{errMessage}}</v-alert>
      </v-col>
    </v-row>

  </v-container>
</template>

<script>
import http from '../http-common'

function formatError (err) {
  if (err.response) {
    const e = err.response.data
    return e.id + ' ' + e.detail
  } else {
    return err.message
  }
}

export default {
  name: 'ProtoFileConverter',

  data: () => ({
    fileName: '',
    fileContent: '',
    docContent: '',
    errMessage: ''
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
            this.errMessage = ''
          })
          .catch(err => {
            this.docContent = ''
            this.errMessage = formatError(err)
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

<style src="../assets/markdown.css">
/* global styles for docContent */
</style>
