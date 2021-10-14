<template>
  <div class="container">
    <h1 v-text="playText"></h1>
    <div class="flex">
      <button v-on:click="changeMode" v-text="buttonText">Let's Go</button>
    </div>
    <div class="game-container" v-if="play">
      <h3>Rock Paper Scissors game, please chose one option</h3>
      <div class="flex">
        <Dropdown
          class="list"
          :options="[
            { id: 1, name: 'Rock' },
            { id: 2, name: 'Paper' },
            { id: 3, name: 'Scissors' },
            { id: 4, name: 'Spock' },
            { id: 5, name: 'Lizard' },
          ]"
          :disabled="false"
          :maxItem="10"
          v-on:selected="selectOption"
          placeholder="Please select an option"
        >
        </Dropdown>
      </div>
      <div class="flex">
        <button v-on:click="send" :disabled="selectedOption === ''">
          Let's Go
        </button>
      </div>
    </div>
    <div class="game-container" v-if="!play">
      <h3>Rock Paper Scissors game, hit the button to simulate a game</h3>
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
  name: "Computer",
  props: {
    msg: String,
  },
  data: function () {
    return {
      selectedOption: "",
      fetchedData: false,
      information: {},
      play: true,
    };
  },
  computed: {
    playText() {
      if (this.play) {
        return "Player vs Computer";
      } else {
        return "Computer vs Computer";
      }
    },
    buttonText() {
      if (this.play) {
        return "Watch a game";
      } else {
        return "Play";
      }
    },
  },
  components: {
    Dropdown,
  },
  methods: {
    send: async function () {
      this.information = await this.getEvents(this.selectedOption);
      this.fetchedData = true;
    },
    clear: function () {
      this.information = {};
      this.fetchedData = false;
    },
    changeMode: function () {
      this.play = !this.play;
      this.information = {};
      this.fetchedData = false;
    },
    selectOption: function (event) {
      if (event.name) {
        this.selectedOption = event.name;
      }
    },
    getEvents: async function (selectOption) {
      let res = {};
      if (this.play) {
        res = await axios.get(
          `http://127.0.0.1:3000/play?option=${selectOption}`
        );
      } else {
        res = await axios.get(`http://127.0.0.1:3000/watch`);
      }
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
  min-width: 100px;
}
</style>
