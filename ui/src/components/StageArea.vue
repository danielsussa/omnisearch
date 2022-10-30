<template>
  <div v-hotkey="keymap">
    <span></span>
      <InputItem></InputItem>
      <div v-for="(menuItem,idx) in slackJson.messages.matches">
        <SlackItem :user-name="menuItem.username" :raw-object="menuItem" :v-bind="'item_'+idx"></SlackItem>
      </div>
  </div>
</template>
'
<script>

import SlackItem from "@/components/SlackItem";
import InputItem from "@/components/InputItem";
import slackJson from './slack.json'
export default {
  data() {
    return {
      slackJson: slackJson,
    };
  },

  name: "StageArea",
  components: {InputItem, SlackItem},
  async mounted() {
  },
  computed: {
    keymap() {
      return {
        "tab": this.toggleModal,
        "shift+o": this.toggleModal
      };
    }
  },
  methods: {
    toggleModal() {
      console.log('modal here')
      console.log(this.$refs.inputName)

    }
  }
}
</script>

<style scoped>

</style>