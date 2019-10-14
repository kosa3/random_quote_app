<template>
  <div class="page">
    <div class="box quote">
      <transition name="fade">
        <p v-if="showText" class="text">
          {{ this.quote.text }}
        </p>
      </transition>
      <transition name="fade">
        <p v-if="showAuthor" class="author">
          {{ this.quote.author }}<span> 【{{ this.quote.tag }}】</span>
        </p>
      </transition>
      <transition name="fade">
        <div v-if="showRestart">
          <nuxt-link to="/" class="btn-flat-logo">RESTART</nuxt-link>
        </div>
      </transition>
    </div>
  </div>
</template>
<style scoped>
.box.quote {
  background-color: #ff7665;
  display: flex;
  flex-direction: column;
  align-items: center;
}
p {
  width: 100%;
  color: white;
  text-align: center;
  font-family: Menlo, sans-serif;
}
p.text {
  font-size: 28px;
  font-family: '游明朝', YuMincho, 'ヒラギノ明朝 ProN W3',
    'Hiragino Mincho ProN', 'HG明朝E', 'ＭＳ Ｐ明朝', 'ＭＳ 明朝', serif;
  margin-bottom: 30px;
}
div.author {
  font-style: italic;
  text-align: right;
}
.btn-flat-logo {
  display: inline-block;
  font-size: 18px;
  margin-top: 20px;
  padding: 0.25em 0.5em;
  text-decoration: none;
  color: #fff;
  background: darkgray;
  transition: 0.4s;
}

.btn-flat-logo:hover {
  background: #1ec7bb;
}
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
.fadein-enter-active,
.fadein-leave-active {
  transition: opacity 0.5s;
}
.fadein-enter,
.fadein-leave-to {
  opacity: 0;
}
</style>
<script>
import axios from 'axios'
export default {
  data() {
    return {
      showText: false,
      showAuthor: false,
      showRestart: false,
      data: this.data,
      quote: ''
    }
  },
  async asyncData({ params }) {
    const { data } = await axios.get(
      `${process.env.baseUrl}/data/${params.quote}.json`
    )

    return { data }
  },
  created() {
    const index = Math.floor(Math.random() * this.data.length)
    this.quote = this.data[index]
  },
  mounted() {
    setInterval(() => {
      this.showText = true
    }, 1000)
    setInterval(() => {
      this.showAuthor = true
    }, 3000)
    setInterval(() => {
      this.showRestart = true
    }, 10000)
  }
}
</script>
