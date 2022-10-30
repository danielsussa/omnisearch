<template>
  <div>
    <div class="input-container">
      <input class="search">
    </div>

  </div>
</template>

<script>
export default {
  name: "InputItem"
}
</script>

<style scoped>
.input-container{
  height: 50px;
  padding: 20px;
}
.search {
  width: 100%;
  height: 40px;
  background-color: #2b3942;
  border-color: transparent;
  color: antiquewhite;
  font-size: 24px;
}
</style>