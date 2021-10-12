<template>
  <div class="container">
    <h1>Player vs computer</h1>
    <div class="game-container">
      <h3>Rock Paper Scissors game, please chose one option</h3>
      <div class="flex">
        <Dropdown class="list"
          :options="[{ id: 1, name: 'rock'}, { id: 2, name: 'paper'}, { id: 3, name: 'scissors'}]"
          :disabled="false"
          :maxItem="10"
          v-on:selected="selectOption"
          placeholder="Please select an option">
        </Dropdown>
      </div>
      <div class="flex">
        <button v-on:click="send">Let's Go</button>
      </div>
    </div>
  </div>
</template>

<script>
import Dropdown from "vue-simple-search-dropdown";
import axios from "axios"
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data: function () {
    return {
      selectedOption: 0
    }
  },
  components: {
    Dropdown
  },
  methods: {
    send: async function () {
      let data = await this.getEvents(this.selectedOption)
      console.log(data)
    },
    selectOption: function (event) {
      this.selectedOption = event.name
      
      },
    getEvents: async function(selectOption){
      let res = await axios.get(`http://127.0.0.1:3000/play?option=${selectOption}`);
      return res.data;
    },
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
div.flex{
  display: flex;
  padding: 20px 10%
}
div.game-container {
  border: 2px solid #2c3e50;
  border-radius: 5px;
  margin: 100px;
}
button{
  margin-left: auto;
  margin-right: auto;
}
</style>
