<template>
  <!-- <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/about">About</router-link>
  </nav> -->
  <router-view/>
</template>

<style>
@font-face {
  font-family: 'comic sans ms';
  src: url('~@/assets/COMIC.TTF')format('truetype');
}
body{
  overflow: hidden;
  font-family: 'comic sans ms';
}
#app {
  /* font-family: Avenir, Helvetica, Arial, sans-serif; */
  font-family: 'comic sans ms';
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #CDFBE4;
  height: 100vh;
  width: 100vw;
  background-color: #0F612F;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
*{
  margin: 0;
  padding: 0;
  box-sizing: content-box;
}
</style>
