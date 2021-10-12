<template>
  <div class="container">
    <h1>Player vs computer</h1>
    <div class="game-container">
      <h3>Rock Paper Scissors game, please chose one option</h3>
      <div class="flex">
        <Dropdown
          class="list"
          :options="[
            { id: 1, name: 'rock' },
            { id: 2, name: 'paper' },
            { id: 3, name: 'scissors' },
          ]"
          :disabled="false"
          :maxItem="10"
          v-on:selected="selectOption"
          placeholder="Please select an option"
        >
        </Dropdown>
      </div>
      <div class="flex">
        <button v-on:click="send">Let's Go</button>
      </div>
    </div>
    <div class="response-container" v-if="fetchedData">
      <table border="1" width="100%">
        <tr>
          <th>Player One move</th>
          <th>Player Two move</th>
          <th>Winner</th>
        </tr>
        <tr>
          <td v-text="information.p1Move"></td>
          <td v-text="information.p2Move"></td>
          <td v-text="information.winner || 'Draw'"></td>
        </tr>
      </table>

      <div class="flex">
        <button v-on:click="clear">Clear data</button>
      </div>
    </div>
  </div>
</template>

<script>
import Dropdown from "vue-simple-search-dropdown";
import axios from "axios";
export default {
  name: "HelloWorld",
  props: {
    msg: String,
  },
  data: function() {
    return {
      selectedOption: "",
      fetchedData: false,
      information: {},
    };
  },
  components: {
    Dropdown,
  },
  methods: {
    send: async function() {
      this.information = await this.getEvents(this.selectedOption);
      console.log(this.information);
      this.fetchedData = true;
    },
    clear: function() {
      this.information = {};
      this.fetchedData = false;
    },
    selectOption: function(event) {
      this.selectedOption = event.name;
    },
    getEvents: async function(selectOption) {
      let res = await axios.get(
        `http://127.0.0.1:3000/play?option=${selectOption}`
      );
      return res.data;
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
div.flex {
  display: flex;
  padding: 20px 10%;
}
div.game-container {
  border: 2px solid #2c3e50;
  border-radius: 5px;
  margin-left: 20%;
  margin-right: 20%;
}
div.response-container {
  margin-left: 20%;
  margin-right: 20%;
  margin-top: 100px;
}

table {
  margin-left: auto;
  margin-right: auto;
}
tr {
  margin: 0 15px;
}
button {
  margin-left: auto;
  margin-right: auto;
}
</style>
