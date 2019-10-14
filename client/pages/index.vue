<template>
  <div class="page">
    <div class="box quote">
      <p class="quote-mark">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink"
          width="100pt"
          height="100pt"
          viewBox="0 0 100 100"
          version="1.1"
        >
          <g id="surface1">
            <path
              style=" stroke:none;fill-rule:nonzero;fill:rgb(100%,100%,100%);fill-opacity:1;"
              d="M 29.015625 12.425781 L 14.09375 12.425781 C 6.308594 12.425781 0 18.734375 0 26.519531 L 0 41.441406 C 0 49.226562 6.308594 55.539062 14.09375 55.539062 L 15.570312 55.539062 C 16.808594 55.539062 17.96875 56.148438 18.671875 57.171875 C 19.367188 58.191406 19.523438 59.496094 19.074219 60.648438 L 10.65625 82.460938 C 10.210938 83.617188 10.363281 84.921875 11.0625 85.941406 C 11.765625 86.964844 12.925781 87.574219 14.164062 87.574219 L 24.070312 87.574219 C 25.46875 87.574219 26.753906 86.796875 27.402344 85.558594 L 39.363281 62.703125 C 41.828125 57.996094 43.113281 52.765625 43.113281 47.453125 L 43.113281 26.519531 C 43.113281 18.734375 36.800781 12.425781 29.015625 12.425781 Z M 29.015625 12.425781 "
            />
            <path
              style=" stroke:none;fill-rule:nonzero;fill:rgb(100%,100%,100%);fill-opacity:1;"
              d="M 85.90625 12.425781 L 70.984375 12.425781 C 63.199219 12.425781 56.890625 18.734375 56.890625 26.519531 L 56.890625 41.441406 C 56.890625 49.226562 63.199219 55.539062 70.980469 55.539062 L 72.457031 55.539062 C 73.695312 55.539062 74.855469 56.148438 75.554688 57.171875 C 76.257812 58.191406 76.40625 59.496094 75.964844 60.648438 L 67.542969 82.460938 C 67.09375 83.617188 67.25 84.921875 67.953125 85.941406 C 68.648438 86.964844 69.808594 87.574219 71.050781 87.574219 L 80.957031 87.574219 C 82.355469 87.574219 83.640625 86.796875 84.289062 85.558594 L 96.25 62.703125 C 98.714844 57.996094 100 52.765625 100 47.453125 L 100 26.519531 C 100 18.734375 93.691406 12.425781 85.90625 12.425781 Z M 85.90625 12.425781 "
            />
          </g>
        </svg>
      </p>
      <transition name="fade">
        <div v-if="show">
          <p class="title">Random Quote</p>
        </div>
      </transition>
      <transition name="fadein">
        <div v-if="show">
          <p><nuxt-link to="/quote" class="btn-flat-logo">START</nuxt-link></p>
        </div>
      </transition>
    </div>
  </div>
</template>
<style scoped>
@-webkit-keyframes bounce {
  0%,
  25%,
  50%,
  75%,
  100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-10px);
  }
  60% {
    transform: translateY(-12px);
  }
}
@-moz-keyframes bounce {
  0%,
  25%,
  50%,
  75%,
  100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-10px);
  }
  60% {
    transform: translateY(-12px);
  }
}
@keyframes bounce {
  0%,
  25%,
  50%,
  75%,
  100% {
    transform: translateY(0);
  }
  40% {
    transform: translateY(-10px);
  }
  60% {
    transform: translateY(-12px);
  }
}

@-webkit-keyframes fadein {
  from {
    transform: translateY(50px);
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
@-moz-keyframes fadein {
  from {
    transform: translateY(50px);
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
@keyframes fadein {
  from {
    transform: translateY(50px);
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
#surface1 {
  -webkit-animation: fadeIn 1s ease-in forwards;
  animation: bounce 2s;
}
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
p.quote-mark {
  font-weight: bold;
}
p.title {
  font-size: 80px;
  font-weight: bold;
  margin-bottom: 30px;
}
.fade-enter-active,
.fade-leave-active {
  -webkit-animation: fadeIn 1s ease-in forwards;
  animation: fadein 1s;
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
.btn-flat-logo {
  display: inline-block;
  font-size: 25px;
  font-weight: bold;
  margin-top: 20px;
  padding: 0.25em 0.5em;
  text-decoration: none;
  color: #fff;
  background: #00bcd4;
  transition: 0.4s;
}

.btn-flat-logo:hover {
  background: #1ec7bb;
}
</style>
<script>
export default {
  transition: {
    name: 'test',
    mode: 'out-in'
  },
  data() {
    return {
      show: false
    }
  },
  mounted() {
    setInterval(() => {
      this.show = true
    }, 2000)
  }
}
</script>
