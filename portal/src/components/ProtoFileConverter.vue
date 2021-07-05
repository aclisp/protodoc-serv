<template>
  <v-container>

    <div>选取一个 Protobuf 文件，把它转换为文档。</div>

    <v-row no-gutters justify="center" align="center">
      <v-col cols="8">
        <v-file-input
          show-size
          label="Protobuf 协议文件"
          @change="selectFile"
        ></v-file-input>
      </v-col>

      <v-col cols="4" class="pl-2">
        <v-btn color="primary" dark @click="convert">
          转&nbsp;换
        </v-btn>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="6">
        <code>{{fileContent}}</code>
      </v-col>

      <v-col cols="6">
        <div v-html="docContent"></div>
      </v-col>
    </v-row>

  </v-container>
</template>

<script>
import http from '../http-common'

export default {
  name: 'ProtoFileConverter',

  data: () => ({
    fileContent: '',
    docContent: ''
  }),

  methods: {
    selectFile (file) {
      const reader = new FileReader()
      reader.onload = (event) => {
        const contents = event.target.result
        this.fileContent = contents
      }
      reader.readAsText(file)
    },
    convert () {
      const contents = this.fileContent
      if (contents) {
        http.post('/protodoc/ProtoDoc/Convert', { proto: contents })
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
