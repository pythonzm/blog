<template>
  <div :class="className" class="link--mallki">
    <router-link
      :style="{ fontSize: text.length * 2 + 13 + 'px' }"
      :to="{
        name: 'CTQArticle',
        query: { tag: text }
      }"
      >{{ text }}<span :data-letters="text"/> <span :data-letters="text"
    /></router-link>
  </div>
</template>

<script>
export default {
  props: {
    className: {
      type: String,
      default: ''
    },
    text: {
      type: String,
      default: ''
    },
  }
}
</script>

<style>
/* Mallki */

.link--mallki {
  font-weight: 800;
  color: #555;
  font-family: "Dosis", sans-serif;
  -webkit-transition: color 0.5s 0.25s;
  transition: color 0.5s 0.25s;
  overflow: hidden;
  position: relative;
  display: inline-block;
  line-height: 1;
  outline: none;
  text-decoration: none;
}

.link--mallki:hover {
  -webkit-transition: none;
  transition: none;
  color: transparent;
}

.link--mallki::before {
  content: "";
  width: 100%;
  height: 6px;
  margin: -3px 0 0 0;
  background: #c8c9ca;
  position: absolute;
  left: 0;
  top: 50%;
  -webkit-transform: translate3d(-100%, 0, 0);
  transform: translate3d(-100%, 0, 0);
  -webkit-transition: -webkit-transform 0.4s;
  transition: transform 0.4s;
  -webkit-transition-timing-function: cubic-bezier(0.7, 0, 0.3, 1);
  transition-timing-function: cubic-bezier(0.7, 0, 0.3, 1);
}

.link--mallki:hover::before {
  -webkit-transform: translate3d(100%, 0, 0);
  transform: translate3d(100%, 0, 0);
}

.link--mallki span {
  position: absolute;
  height: 50%;
  width: 100%;
  left: 0;
  top: 0;
  overflow: hidden;
}

.link--mallki span::before {
  content: attr(data-letters);
  color: red;
  position: absolute;
  left: 0;
  width: 100%;
  color: #555;
  -webkit-transition: -webkit-transform 0.5s;
  transition: transform 0.5s;
}

.link--mallki span:nth-child(2) {
  top: 50%;
}

.link--mallki span:first-child::before {
  top: 0;
  -webkit-transform: translate3d(0, 100%, 0);
  transform: translate3d(0, 100%, 0);
}

.link--mallki span:nth-child(2)::before {
  bottom: 0;
  -webkit-transform: translate3d(0, -100%, 0);
  transform: translate3d(0, -100%, 0);
}

.link--mallki:hover span::before {
  -webkit-transition-delay: 0.3s;
  transition-delay: 0.3s;
  -webkit-transform: translate3d(0, 0, 0);
  transform: translate3d(0, 0, 0);
  -webkit-transition-timing-function: cubic-bezier(0.2, 1, 0.3, 1);
  transition-timing-function: cubic-bezier(0.2, 1, 0.3, 1);
}
</style>
