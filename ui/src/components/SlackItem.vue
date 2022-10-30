<template>
  <div class="item-container">
    <div class="item-container-2">
      <div class="flex-container">
        <h1>{{rawObject.username}}</h1>
        <p>Hello here</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "SlackItem",
  props: {
    userName: String,
    rawObject: Object,
  },
}
</script>

<style scoped>

  .item-container{
    /*padding: 20px;*/
    padding-left: 20px;
    padding-right: 20px;
    color: wheat;
    /*background-color: red;*/
  }

  p {
    color: #a2aab4;
  }

  .item-container-2{
    width: 100%;
    border-bottom: #415260 solid 1px;
  }

  .flex-container{
    /*display: flex;*/
  }
</style>