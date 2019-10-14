<template>
  <div>
    <nuxt />
  </div>
</template>

<style>
html {
  font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
    Roboto, 'Helvetica Neue', Arial, Menlo, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;
}

*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;
}
.page {
  height: 100vh;
  overflow: scroll;
}

.box {
  display: -webkit-box;
  display: -ms-flexbox;
  display: flex;

  /* 左右中央寄せ */
  -webkit-box-pack: center;
  -ms-flex-pack: center;
  justify-content: center;

  /* 上下中央寄せ */
  -webkit-box-align: center;
  -ms-flex-align: center;
  align-items: center;

  padding: 30px;
  box-sizing: border-box;

  width: 100vw;
  height: auto;
  min-height: 100vh;

  background-size: cover;
  background-position: center;
}
</style>
<script>
export default {
  transition: 'bounce'
}
</script>
