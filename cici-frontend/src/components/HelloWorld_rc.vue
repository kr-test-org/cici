<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>{{ description }} </h2>
    <ul>
      <li>
        前端构建时间: hyperVersion
      </li>
    </ul>
    <ul>
      <li>
        后端构建时间: {{ buildstamp.data }}
      </li>
    </ul>
    <h3>推荐链接</h3>
    <ul>
      <li>
        <a
          href="https://www.koderover.com/"
          target="_blank"
        >
          koderover官网
        </a>
      </li>
      <li>
        <a
          href="https://10x.koderover.com/"
          target="_blank"
        >
          10x 演示环境
        </a>
      </li>
      <br>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: '欢迎来到cici项目',
      description: '前端采用容器化部署，后端采用传统的物理机部署',
      buildstamp: ""
    }
  },
  mounted () {
    axios
      .get('/api/buildstamp')
      .then(response => (this.buildstamp= response))
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
