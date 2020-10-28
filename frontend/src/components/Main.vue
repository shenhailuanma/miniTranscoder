<template>
  <div class="main"></div>
</template>

<script>
export default {
  name: "Main",
  data: function() {
    return {
      msg: "Welcome"
    };
  },
  methods: {
    test() {
      console.log("test method");
    }
  },
  mounted: function() {
    console.log("mounted");
  },
  activated: function() {
    console.log("activated");
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
